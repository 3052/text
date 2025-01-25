package http

import (
   "io"
   "net/http"
   "testing"
)

func TestHttp(t *testing.T) {
   http.DefaultClient.Transport = transport{}
   resp, err := http.Get("http://httpbingo.org/get")
   if err != nil {
      t.Fatal(err)
   }
   resp.Write(io.Discard)
}
