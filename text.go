package text

import (
   "bytes"
   "io"
   "log"
   "log/slog"
   "net/http"
   "strconv"
   "strings"
   "text/template"
   "time"
)

func (Transport) Set(on bool) {
   if on {
      http.DefaultTransport = Transport{}
   } else {
      http.DefaultTransport = DefaultTransport
   }
   log.SetFlags(log.Ltime)
}

func (p *ProgressMeter) Set(parts int) {
   p.date = time.Now()
   p.modified = time.Now()
   p.parts.length = int64(parts)
}

var DefaultTransport = http.DefaultTransport

type Transport struct{}

func (Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if req.Method == "" {
      req.Method = "GET"
   }
   slog.Info(req.Method, "URL", req.URL)
   return DefaultTransport.RoundTrip(req)
}

var DefaultName =
   "{{if .Show}}" +
      "{{.Show}} - {{.Season}} {{.Episode}} - {{.Title}}" +
   "{{else}}" +
      "{{.Title}} - {{.Year}}" +
   "{{end}}"

func Clean(s string) string {
   mapping := func(r rune) rune {
      if strings.ContainsRune(`"*/:<>?\|`, r) {
         return '-'
      }
      return r
   }
   return strings.Map(mapping, s)
}

func Name(n Namer) (string, error) {
   text, err := new(template.Template).Parse(DefaultName)
   if err != nil {
      return "", err
   }
   var b strings.Builder
   err = text.Execute(&b, n)
   if err != nil {
      return "", err
   }
   return b.String(), nil
}

func CutBefore(s, sep []byte) ([]byte, []byte, bool) {
   if i := bytes.Index(s, sep); i >= 0 {
      return s[:i], s[i:], true
   }
   return s, nil, false
}

type Cardinal float64

type Namer interface {
   Show() string
   Season() int
   Episode() int
   Title() string
   Year() int
}

type ProgressMeter struct {
   first int
   last int64
   length int64
   parts struct {
      last int64
      length int64
   }
   modified time.Time
   date time.Time
}

func (p ProgressMeter) percent() Percent {
   return Percent(p.first) / Percent(p.length)
}

func (p ProgressMeter) rate() Rate {
   return Rate(p.first) / Rate(time.Since(p.date).Seconds())
}

func (p ProgressMeter) size() Size {
   return Size(p.first)
}

type Rate float64

type Percent float64

type Size float64

func (p *ProgressMeter) Reader(res *http.Response) io.Reader {
   p.parts.last += 1
   p.last += res.ContentLength
   p.length = p.last * p.parts.length / p.parts.last
   return io.TeeReader(res.Body, p)
}

type unit_measure struct {
   factor float64
   name string
}

func (p Percent) String() string {
   unit := unit_measure{100, " %"}
   return label(float64(p), unit)
}

func label(value float64, unit unit_measure) string {
   var prec int
   if unit.factor != 1 {
      prec = 2
      value *= unit.factor
   }
   return strconv.FormatFloat(value, 'f', prec, 64) + unit.name
}

func scale(value float64, units []unit_measure) string {
   var unit unit_measure
   for _, unit = range units {
      if unit.factor * value < 1000 {
         break
      }
   }
   return label(value, unit)
}

func (c Cardinal) String() string {
   units := []unit_measure{
      {1, ""},
      {1e-3, " thousand"},
      {1e-6, " million"},
      {1e-9, " billion"},
   }
   return scale(float64(c), units)
}

func (s Size) String() string {
   units := []unit_measure{
      {1, " byte"},
      {1e-3, " kilobyte"},
      {1e-6, " megabyte"},
      {1e-9, " gigabyte"},
   }
   return scale(float64(s), units)
}

func (r Rate) String() string {
   units := []unit_measure{
      {1, " byte/s"},
      {1e-3, " kilobyte/s"},
      {1e-6, " megabyte/s"},
      {1e-9, " gigabyte/s"},
   }
   return scale(float64(r), units)
}

func (p *ProgressMeter) Write(data []byte) (int, error) {
   p.first += len(data)
   if time.Since(p.modified) >= time.Second {
      slog.Info(p.percent().String(), "size", p.size(), "rate", p.rate())
      p.modified = time.Now()
   }
   return len(data), nil
}
