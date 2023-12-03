package blog

import (
   "context"
   "fmt"
   "log/slog"
   "net/http"
)

type Handler struct {
   Level slog.Level
}

func (h Handler) Enabled(_ context.Context, l slog.Level) bool {
   return l >= h.Level
}

func (Handler) Handle(_ context.Context, r slog.Record) error {
   fmt.Print(r.Message)
   r.Attrs(func(a slog.Attr) bool {
      fmt.Print(" ", a)
      return true
   })
   fmt.Println()
   return nil
}

func (Handler) WithAttrs([]slog.Attr) slog.Handler {
   return nil
}

func (Handler) WithGroup(string) slog.Handler {
   return nil
}

type Transport struct {
   Level slog.Level
}

func (t Transport) RoundTrip(r *http.Request) (*http.Response, error) {
   slog.Log(context.Background(), t.Level, r.Method, "URL", r.URL)
   return http.DefaultTransport.RoundTrip(r)
}

func SetTransport(l slog.Level) {
   http.DefaultClient.Transport = Transport{l}
}
