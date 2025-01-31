package http

import (
   "log"
   "net/http"
)

type Transport http.Transport

func (t Transport) DefaultClient() {
   http.DefaultClient.Transport = &t
}

func (t *Transport) ProxyFromEnvironment() {
   t.Proxy = http.ProxyFromEnvironment
}

func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if req.Method == "" {
      req.Method = "GET"
   }
   log.Println(req.Method, req.URL)
   return (*http.Transport)(t).RoundTrip(req)
}
