package text

import (
   "fmt"
   "io"
   "net/http"
   "reflect"
   "testing"
   "text/template"
)

func TestSize(t *testing.T) {
   size := reflect.TypeOf(&struct{}{}).Size()
   for _, test := range size_tests {
      if reflect.TypeOf(test).Size() > size {
         fmt.Printf("*%T\n", test)
      } else {
         fmt.Printf("%T\n", test)
      }
   }
}

var size_tests = []any{
   Cardinal(0),
   Percent(0),
   ProgressMeter{},
   Rate(0),
   Size(0),
   Transport{},
   unit_measure{},
}

func TestName(t *testing.T) {
   _, err := new(template.Template).Parse(DefaultName)
   if err != nil {
      t.Fatal(err)
   }
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
