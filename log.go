package log

import (
   "context"
   "fmt"
   "log/slog"
   "net/http"
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

func (Handler) WithAttrs([]slog.Attr) slog.Handler {
   return nil
}

func (Handler) WithGroup(string) slog.Handler {
   return nil
}

func (h Handler) Enabled(_ context.Context, l slog.Level) bool {
   return l >= slog.Level(h)
}

type Handler slog.Level

func SetHandler(l slog.Level) {
   h := Handler(l)
   slog.SetDefault(slog.New(h))
}

type Transport slog.Level

func (t Transport) RoundTrip(r *http.Request) (*http.Response, error) {
   slog.Log(context.Background(), slog.Level(t), r.Method, "URL", r.URL)
   return http.DefaultTransport.RoundTrip(r)
}

func SetTransport(l slog.Level) {
   http.DefaultClient.Transport = Transport(l)
}
