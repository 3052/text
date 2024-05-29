package text

import (
   "io"
   "net/http"
   "testing"
)

const address = "https://go.dev/dl/go1.21.5.windows-amd64.zip"

func TestOne(t *testing.T) {
   var meter ProgressMeter
   meter.Set(1)
   res, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   io.Copy(io.Discard, meter.Reader(res))
}

func TestTwo(t *testing.T) {
   var meter ProgressMeter
   meter.Set(1)
   LogLevel{}.SetTransport(true)
   res, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   io.Copy(io.Discard, meter.Reader(res))
}

func TestThree(t *testing.T) {
   var meter ProgressMeter
   meter.Set(1)
   var log LogLevel
   log.Set()
   log.SetTransport(true)
   res, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   io.Copy(io.Discard, meter.Reader(res))
}
