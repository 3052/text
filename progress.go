package text

import (
   "io"
   "log/slog"
   "net/http"
   "time"
)

func (p *ProgressMeter) Set(parts int) {
   p.date = time.Now()
   p.modified = time.Now()
   p.parts.length = int64(parts)
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

func (p *ProgressMeter) percent() Percent {
   return Percent(p.first) / Percent(p.length)
}

func (p *ProgressMeter) rate() Rate {
   return Rate(p.first) / Rate(time.Since(p.date).Seconds())
}

func (p ProgressMeter) size() Size {
   return Size(p.first)
}

func (p *ProgressMeter) Reader(resp *http.Response) io.Reader {
   p.parts.last += 1
   p.last += resp.ContentLength
   p.length = p.last * p.parts.length / p.parts.last
   return io.TeeReader(resp.Body, p)
}

func (p *ProgressMeter) Write(data []byte) (int, error) {
   p.first += len(data)
   if time.Since(p.modified) >= time.Second {
      slog.Info(p.percent().String(), "size", p.size(), "rate", p.rate())
      p.modified = time.Now()
   }
   return len(data), nil
}
