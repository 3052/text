package log

import (
   "context"
   "fmt"
   "log/slog"
   "net/http"
)

func Set_Logger(level slog.Level) {
   h := Handler{Level: level}
   h.Handler = slog.Default().Handler()
   *slog.Default() = *slog.New(h)
}

func Set_Transport(level slog.Level) {
   http.DefaultClient.Transport = Transport{level}
}

type Handler struct {
   Level slog.Level
   slog.Handler
}

func (h Handler) Enabled(_ context.Context, level slog.Level) bool {
   return level >= h.Level
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

type Transport struct {
   Level slog.Level
}

func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   slog.Log(
      context.Background(), t.Level, "*", "method", req.Method, "URL", req.URL,
   )
   return http.DefaultTransport.RoundTrip(req)
}