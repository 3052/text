package main

import (
   "io"
   "net/http"
   _ "41.neocities.org/x/http"
)

func main() {
   resp, err := http.Get("http://httpbingo.org/get")
   if err != nil {
      panic(err)
   }
   resp.Write(io.Discard)
}
