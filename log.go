package log

import (
   "context"
   "fmt"
   "log/slog"
   "net/http"
)

func Set_Logger(level Level) {
   h := Handler{level: level}
   h.Handler = slog.Default().Handler()
   *slog.Default() = *slog.New(h)
}

func Set_Transport(level Level) {
   http.DefaultClient.Transport = Handler{level: level}
}

type Handler struct {
   level Level
   slog.Handler
}

func (h Handler) Enabled(_ context.Context, level slog.Level) bool {
   return level >= h.level
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

func (h Handler) RoundTrip(req *http.Request) (*http.Response, error) {
   slog.Log(
      context.Background(), h.level, "*", "method", req.Method, "URL", req.URL,
   )
   return http.DefaultTransport.RoundTrip(req)
}

type Level = slog.Level
