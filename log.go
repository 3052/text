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

func (h Handler) Enabled(_ context.Context, lev slog.Level) bool {
   return lev >= h.Level
}

type Handler struct {
   Level slog.Level
}

func Set_Handler(h Handler) {
   slog.SetDefault(slog.New(h))
}

type Transport struct {
   Level slog.Level
}

func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   slog.Log(
      context.Background(), t.Level, "*", "method", req.Method, "URL", req.URL,
   )
   return http.DefaultTransport.RoundTrip(req)
}

func Set_Transport(lev slog.Level) {
   http.DefaultClient.Transport = Transport{lev}
}
