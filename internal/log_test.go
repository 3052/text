package log

import (
   "io"
   "log"
   "net/http"
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
   log.SetFlags(log.Ltime)
   _, err = io.ReadAll(meter.Reader(resp))
   if err != nil {
      t.Fatal(err)
   }
}
