package http

import (
   "41.neocities.org/x/strconv"
   "io"
   "log"
   "net/http"
   "time"
)

func (p *ProgressMeter) Write(data []byte) (int, error) {
   p.first += len(data)
   now := time.Now()
   if now.Sub(p.modified) >= time.Second {
      log.Print(strconv.Percent(p.first) / strconv.Percent(p.length))
      p.modified = now
   }
   return len(data), nil
}

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

func (p *ProgressMeter) Reader(resp *http.Response) io.Reader {
   p.parts.last += 1
   p.last += resp.ContentLength
   p.length = p.last * p.parts.length / p.parts.last
   return io.TeeReader(resp.Body, p)
}
