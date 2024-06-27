package text

import (
   "io"
   "net/http"
   "testing"
)

func TestClient(t *testing.T) {
   var tr Transport
   tr.New()
   tr.Set(true)
   resp, err := http.Get("http://go.dev")
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   io.Copy(io.Discard, resp.Body)
}

func TestTransport(t *testing.T) {
   var tr Transport
   tr.New()
   tr.Set(true)
   req, err := http.NewRequest("", "http://go.dev", nil)
   if err != nil {
      t.Fatal(err)
   }
   resp, err := http.DefaultTransport.RoundTrip(req)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   io.Copy(io.Discard, resp.Body)
}

const address = "https://go.dev/dl/go1.21.5.windows-amd64.zip"

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
   var tr Transport
   tr.New()
   tr.Set(true)
   resp, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer resp.Body.Close()
   io.Copy(io.Discard, meter.Reader(resp))
}
