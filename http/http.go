package http

import (
   "41.neocities.org/log"
   "io"
   "net/http"
   "time"
   stdLog "log"
)

func init() {
   http.DefaultClient.Transport = Transport{}
}

func (p *ProgressMeter) Write(data []byte) (int, error) {
   p.first += len(data)
   if time.Since(p.modified) >= time.Second {
      stdLog.Printf("%v, %v, %v", p.percent(), p.size(), p.rate())
      p.modified = time.Now()
   }
   return len(data), nil
}

func (Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if req.Method == "" {
      req.Method = "GET"
   }
   stdLog.Println(req.Method, req.URL)
   return http.DefaultTransport.RoundTrip(req)
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

func (p *ProgressMeter) percent() log.Percent {
   return log.Percent(p.first) / log.Percent(p.length)
}

func (p *ProgressMeter) rate() log.Rate {
   return log.Rate(p.first) / log.Rate(time.Since(p.date).Seconds())
}

func (p *ProgressMeter) size() log.Size {
   return log.Size(p.first)
}

func (p *ProgressMeter) Set(parts int) {
   p.date = time.Now()
   p.modified = time.Now()
   p.parts.length = int64(parts)
}

func (p *ProgressMeter) Reader(resp *http.Response) io.Reader {
   p.parts.last += 1
   p.last += resp.ContentLength
   p.length = p.last * p.parts.length / p.parts.last
   return io.TeeReader(resp.Body, p)
}

type Transport struct{}
