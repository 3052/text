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
   th := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
      Level: v.Level,
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

func (p ProgressMeter) percent() encoding.Percent {
   return encoding.Percent(p.first) / encoding.Percent(p.length)
}

func (p ProgressMeter) rate() encoding.Rate {
   return encoding.Rate(p.first) / encoding.Rate(time.Since(p.date).Seconds())
}

func (p ProgressMeter) size() encoding.Size {
   return encoding.Size(p.first)
}

// Level
//  - godocs.io/log/slog#Level.MarshalText
//  - godocs.io/log/slog#Level.UnmarshalText
type Level struct {
   Level slog.Level
}

func TransportDebug() {
   http.DefaultClient.Transport = Level{slog.LevelDebug}
}

func TransportInfo() {
   http.DefaultClient.Transport = Level{slog.LevelInfo}
}

func (v Level) RoundTrip(r *http.Request) (*http.Response, error) {
   slog.Log(context.Background(), v.Level, r.Method, "URL", r.URL)
   return http.DefaultTransport.RoundTrip(r)
}

func (p *ProgressMeter) Write(data []byte) (int, error) {
   p.first += len(data)
   if time.Since(p.modified) >= time.Second {
      slog.Info(p.percent().String(), "size", p.size(), "rate", p.rate())
      p.modified = time.Now()
   }
   return len(data), nil
}
