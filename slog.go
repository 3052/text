package slog

import (
   "context"
   "fmt"
   "log/slog"
   "net/http"
   "strconv"
)

func (Handler) Handle(_ context.Context, r slog.Record) error {
   fmt.Print(r.Message)
   r.Attrs(func(a slog.Attr) bool {
      fmt.Print(" ", a.Key, ":", a.Value)
      return true
   })
   fmt.Println()
   return nil
}

func (h Handler) Enabled(_ context.Context, l slog.Level) bool {
   return l >= h.Level
}

func (Handler) WithAttrs([]slog.Attr) slog.Handler {
   return nil
}

func (Handler) WithGroup(string) slog.Handler {
   return nil
}

type Handler struct {
   Level slog.Level
}
type Cardinal float64

type Rate float64

type Size float64

type Percent float64

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

func scale(value float64, units []unit_measure) string {
   var unit unit_measure
   for _, unit = range units {
      if unit.factor * value < 1000 {
         break
      }
   }
   return label(value, unit)
}

type unit_measure struct {
   factor float64
   name string
}

func (p Percent) String() string {
   unit := unit_measure{factor: 100}
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

func (t Transport) RoundTrip(r *http.Request) (*http.Response, error) {
   slog.Log(context.Background(), t.Level, r.Method, "URL", r.URL)
   return http.DefaultTransport.RoundTrip(r)
}

type Transport struct {
   Level slog.Level
}

func SetHandler(l slog.Level) {
   h := Handler{l}
   slog.SetDefault(slog.New(h))
}

func SetTransport(l slog.Level) {
   http.DefaultClient.Transport = Transport{l}
}
