package log

import (
   "context"
   "fmt"
   "log/slog"
   "net/http"
)

func Set_Handler(lev Level) {
   h := Handler(lev)
   slog.SetDefault(slog.New(h))
}

func Set_Transport(lev Level) {
   http.DefaultClient.Transport = Transport(lev)
}

type Handler Level

func (h Handler) Enabled(_ context.Context, lev Level) bool {
   return lev >= Level(h)
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

type Level = slog.Level

type Transport Level

func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   slog.Log(context.Background(), Level(t), req.Method, "URL", req.URL)
   return http.DefaultTransport.RoundTrip(req)
}
