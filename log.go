package text

import (
   "log/slog"
   "net/http"
)

type Transport struct {
   Transport http.RoundTripper
}

func (t *Transport) New() {
   t.Transport = http.DefaultTransport
}

func (t Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if req.Method == "" {
      req.Method = "GET"
   }
   slog.Info(req.Method, "URL", req.URL)
   return t.Transport.RoundTrip(req)
}

func (t Transport) Set(on bool) {
   if on {
      http.DefaultTransport = t
   } else {
      http.DefaultTransport = t.Transport
   }
}
