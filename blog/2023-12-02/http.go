package http

import (
   "bytes"
   "fmt"
   "net/http"
)

func Trace() func() {
   k := &http.DefaultClient.Transport
   v := *k
   *k = transport(func(req *http.Request) error {
      b := new(bytes.Buffer)
      err := req.Write(b)
      if err != nil {
         return err
      }
      fmt.Println(b)
      if req.Body != nil {
         req.Body, err = req.GetBody()
         if err != nil {
            return err
         }
      }
      return nil
   })
   return func() { *k = v }
}

func Verbose() func() {
   k := &http.DefaultClient.Transport
   v := *k
   *k = transport(func(req *http.Request) error {
      fmt.Println(req.Method, req.URL)
      return nil
   })
   return func() { *k = v }
}

type transport func(*http.Request) error

func (t transport) RoundTrip(req *http.Request) (*http.Response, error) {
   if err := t(req); err != nil {
      return nil, err
   }
   return http.DefaultTransport.RoundTrip(req)
}
