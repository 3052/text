package text

import (
   "log/slog"
   "net/http"
)

func (Transport) Set(on bool) {
   if on {
      http.DefaultTransport = Transport{}
   } else {
      http.DefaultTransport = DefaultTransport
   }
}

var DefaultTransport = http.DefaultTransport

type Transport struct{}

func (Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if req.Method == "" {
      req.Method = "GET"
   }
   slog.Info(req.Method, "URL", req.URL)
   return DefaultTransport.RoundTrip(req)
}
