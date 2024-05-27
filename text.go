package text

import (
   "bytes"
   "log/slog"
   "net/http"
   "os"
   "strconv"
   "strings"
   "text/template"
)

func SetTransport(r http.RoundTripper) {
   http.DefaultClient.Transport = r
}

// Level
//  - godocs.io/log/slog#Level.MarshalText
//  - godocs.io/log/slog#Level.UnmarshalText
type Level struct {
   Level slog.Level
}

func (v Level) Set() {
   text := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
      Level: v.Level,
      ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
         switch a.Key {
         case slog.LevelKey, slog.TimeKey:
            return slog.Attr{}
         }
         return a
      },
   })
   slog.SetDefault(slog.New(text))
}

type Transport struct{}

func (Transport) RoundTrip(r *http.Request) (*http.Response, error) {
   slog.Info(r.Method, "URL", r.URL)
   return http.DefaultTransport.RoundTrip(r)
}

func (t Transport) Set() {
   SetTransport(t)
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

type Cardinal float64

func (c Cardinal) String() string {
   units := []unit_measure{
      {1, ""},
      {1e-3, " thousand"},
      {1e-6, " million"},
      {1e-9, " billion"},
      {1e-12, " trillion"},
   }
   return scale(float64(c), units)
}

type Percent float64

func (p Percent) String() string {
   unit := unit_measure{100, " %"}
   return label(float64(p), unit)
}

type Rate float64

func (r Rate) String() string {
   units := []unit_measure{
      {1, " byte/s"},
      {1e-3, " kilobyte/s"},
      {1e-6, " megabyte/s"},
      {1e-9, " gigabyte/s"},
      {1e-12, " terabyte/s"},
   }
   return scale(float64(r), units)
}

type Size float64

func (s Size) String() string {
   units := []unit_measure{
      {1, " byte"},
      {1e-3, " kilobyte"},
      {1e-6, " megabyte"},
      {1e-9, " gigabyte"},
      {1e-12, " terabyte"},
   }
   return scale(float64(s), units)
}

type unit_measure struct {
   factor float64
   name string
}
var Format = 
   "{{if .Show}}" +
      "{{.Show}} - {{.Season}} {{.Episode}} - {{.Title}}" +
   "{{else}}" +
      "{{.Title}} - {{.Year}}" +
   "{{end}}"

func CutBefore(s, sep []byte) ([]byte, []byte, bool) {
   if i := bytes.Index(s, sep); i >= 0 {
      return s[:i], s[i:], true
   }
   return s, nil, false
}

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
   text, err := new(template.Template).Parse(Format)
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

type Namer interface {
   Show() string
   Season() int
   Episode() int
   Title() string
   Year() int
}
