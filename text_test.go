package text

import (
   "fmt"
   "io"
   "net/http"
   "testing"
)

func TestTransport(t *testing.T) {
   req, err := http.NewRequest("", "http://go.dev", nil)
   if err != nil {
      t.Fatal(err)
   }
   Transport{}.Set(true)
   resp, err := http.DefaultTransport.RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   io.Copy(io.Discard, resp.Body)
}

func TestClient(t *testing.T) {
   get := func() {
      resp, err := http.Get("https://go.dev")
      if err != nil {
         t.Fatal(err)
      }
      defer resp.Body.Close()
      io.Copy(io.Discard, resp.Body)
   }
   Transport{}.Set(true)
   get()
   Transport{}.Set(false)
   get()
}

const address = "https://dl.google.com/go/go1.21.5.windows-amd64.zip"

func TestMeterOne(t *testing.T) {
   var meter ProgressMeter
   meter.Set(1)
   resp, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   io.Copy(io.Discard, meter.Reader(resp))
}

func TestMeterTwo(t *testing.T) {
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

func TestPercent(t *testing.T) {
   fmt.Println(Percent(1234) / 10000)
}
