package text

import (
   "log/slog"
   "net/http"
   "os"
)

type LogLevel struct {
   Level slog.Level
}

func (LogLevel) RoundTrip(req *http.Request) (*http.Response, error) {
   slog.Info(req.Method, "URL", req.URL)
   return http.DefaultTransport.RoundTrip(req)
}

func (v LogLevel) Set() {
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

func (LogLevel) SetTransport(value bool) {
   if value {
      http.DefaultClient.Transport = LogLevel{}
   } else {
      http.DefaultClient.Transport = nil
   }
}
