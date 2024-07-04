package text

import (
   "io"
   "net/http"
   "testing"
)

const address = "https://dl.google.com/go/go1.21.5.windows-amd64.zip"

func TestMeter(t *testing.T) {
   var meter ProgressMeter
   meter.Set(1)
   Transport{}.Set(true)
   resp, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   io.Copy(io.Discard, meter.Reader(resp))
}
