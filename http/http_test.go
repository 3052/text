package http

import (
   "io"
   "log"
   "net/http"
   "net/url"
   "strings"
   "testing"
)

func TestProgressMeter(t *testing.T) {
   resp, err := http.Get("https://dl.google.com/go/go1.23.5.windows-amd64.zip")
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   var meter ProgressMeter
   meter.Set(1)
   _, err = io.ReadAll(meter.Reader(resp))
   if err != nil {
      t.Fatal(err)
   }
}

func TestTransport(t *testing.T) {
   var out strings.Builder
   log.SetOutput(&out)
   req := http.Request{
      URL: &url.URL{Scheme:"http", Host: "example.com"},
   }
   resp, err := http.DefaultClient.Do(&req)
   if err != nil {
      t.Fatal(err)
   }
   err = resp.Write(io.Discard)
   if err != nil {
      t.Fatal(err)
   }
   if !strings.HasSuffix(out.String(), " INFO GET URL=http://example.com\n") {
      t.Fatal(&out)
   }
}
