package log

import (
   "154.pages.dev/encoding"
   "context"
   "io"
   "log/slog"
   "net/http"
   "os"
   "time"
)

func Handler(v Level) {
   th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
      Level: slog.Level(v),
      ReplaceAttr: func(_ []string, a slog.Attr) slog.Attr {
         switch a.Key {
         case slog.LevelKey, slog.TimeKey:
            return slog.Attr{}
         }
         return a
      },
   })
   slog.SetDefault(slog.New(th))
}

func TransportDebug() {
   http.DefaultClient.Transport = Level(slog.LevelDebug)
}

func TransportInfo() {
   http.DefaultClient.Transport = Level(slog.LevelInfo)
}

type Level slog.Level

func (v Level) RoundTrip(r *http.Request) (*http.Response, error) {
   slog.Log(
      context.Background(), slog.Level(v), "request",
      "method", r.Method, "URL", r.URL,
   )
   return http.DefaultTransport.RoundTrip(r)
}

type ProgressMeter struct {
   first int
   last int64
   length int64
   parts struct {
      last int64
      length int64
   }
   modified time.Time
   date time.Time
}

func (p *ProgressMeter) Reader(res *http.Response) io.Reader {
   p.parts.last += 1
   p.last += res.ContentLength
   p.length = p.last * p.parts.length / p.parts.last
   return io.TeeReader(res.Body, p)
}

func (p *ProgressMeter) Set(parts int) {
   p.date = time.Now()
   p.modified = time.Now()
   p.parts.length = int64(parts)
}

func (p *ProgressMeter) Write(data []byte) (int, error) {
   p.first += len(data)
   if time.Since(p.modified) >= time.Second {
      slog.Info(
         "progress", "%", p.percent(), "size", p.size(), "rate", p.rate(),
      )
      p.modified = time.Now()
   }
   return len(data), nil
}

func (p ProgressMeter) percent() encoding.Percent {
   return encoding.Percent(p.first) / encoding.Percent(p.length)
}

func (p ProgressMeter) rate() encoding.Rate {
   return encoding.Rate(p.first) / encoding.Rate(time.Since(p.date).Seconds())
}

func (p ProgressMeter) size() encoding.Size {
   return encoding.Size(p.first)
}
