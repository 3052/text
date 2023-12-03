package slog

import (
   "io"
   "log/slog"
   "net/http"
   "time"
)

type Progress struct {
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

func Progress_Length(length int64) *Progress {
   var p Progress
   p.length = length
   p.modified = time.Now()
   p.date = time.Now()
   return &p
}

func Progress_Parts(length int) *Progress {
   var p Progress
   p.modified = time.Now()
   p.date = time.Now()
   p.parts.length = int64(length)
   return &p
}

func (p *Progress) Write(b []byte) (int, error) {
   p.first += len(b)
   if time.Since(p.modified) >= time.Second {
      slog.Info("*", "percent", p.percent(), "size", p.size(), "rate", p.rate())
      p.modified = time.Now()
   }
   return len(b), nil
}

func (p *Progress) Reader(res *http.Response) io.Reader {
   if p.parts.length >= 1 {
      p.parts.last += 1
      p.last += res.ContentLength
      p.length = p.last * p.parts.length / p.parts.last
   }
   return io.TeeReader(res.Body, p)
}

func (p Progress) percent() Percent {
   return Percent(p.first) / Percent(p.length)
}

func (p Progress) rate() Rate {
   return Rate(p.first) / Rate(time.Since(p.date).Seconds())
}

func (p Progress) size() Size {
   return Size(p.first)
}
