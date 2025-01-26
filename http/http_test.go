package http

import (
   "io"
   "net/http"
   "net/url"
   "testing"
)

func TestBytes(t *testing.T) {
   Transport{
      DisableCompression: true,
      Proxy: http.ProxyFromEnvironment,
   }.Set()
   req := http.Request{URL: &url.URL{
      Scheme: "http",
      Host: "httpbingo.org",
      Path: "/drip",
      RawQuery: "delay=0&duration=9",
   }}
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {   
      t.Fatal(err)
   }
   defer resp.Body.Close()
   var progress ProgressBytes
   progress.Set(resp)
   _, err = io.ReadAll(&progress)
   if err != nil {   
      t.Fatal(err)
   }
}

func TestParts(t *testing.T) {
   http.DefaultClient.Transport = nil
   var progress ProgressParts
   progress.Set(9)
   for {
      resp, err := http.Get("http://httpbingo.org/drip?delay=0&duration=1")
      if err != nil {   
         t.Fatal(err)
      }
      err = resp.Write(io.Discard)
      if err != nil {   
         t.Fatal(err)
      }
      if !progress.Next() {
         break
      }
   }
}
