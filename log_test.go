package log

import (
   "io"
   "net/http"
   "testing"
)

const address = "https://go.dev/dl/go1.21.5.windows-amd64.zip"

func TestSlog(t *testing.T) {
   Handler(Level{})
   TransportInfo()
   var meter ProgressMeter
   meter.Set(1)
   res, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   io.Copy(io.Discard, meter.Reader(res))
}
