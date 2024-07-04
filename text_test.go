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

func TestPercent(t *testing.T) {
   fmt.Println(Percent(1234) / 10000)
}
