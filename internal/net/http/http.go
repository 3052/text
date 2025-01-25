package http

import (
   "log"
   "net/http"
)

type transport struct{}

func (transport) RoundTrip(req *http.Request) (*http.Response, error) {
   log.Print(req.URL)
   return http.DefaultTransport.RoundTrip(req)
}
