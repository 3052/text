package log

import (
   "context"
   "fmt"
   "log/slog"
   "net/http"
)

func Set_Handler(lev slog.Level) {
   h := Handler(lev)
   slog.SetDefault(slog.New(h))
}

func Set_Transport(lev slog.Level) {
   http.DefaultClient.Transport = Transport(lev)
}

type Handler slog.Level

func (h Handler) Enabled(_ context.Context, lev slog.Level) bool {
   return lev >= slog.Level(h)
}

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

type Transport slog.Level

func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   slog.Log(context.Background(), slog.Level(t), req.Method, "URL", req.URL)
   return http.DefaultTransport.RoundTrip(req)
}
