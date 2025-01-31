package http

import (
   "io"
   "net/http"
   "net/url"
   "testing"
)

func TestBytes(t *testing.T) {
   tr := Transport{DisableCompression: true}
   tr.ProxyFromEnvironment()
   tr.DefaultClient()
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
