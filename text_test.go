package text

import (
   "fmt"
   "io"
   "net/http"
   "testing"
)

var name_tests = []test_alfa{
   {
      title: "title",
      year: 2024,
   },
   {
      show: "show",
      title: "title",
   },
   {
      show: "show",
      episode: 3,
      title: "title",
   },
   {
      show: "show",
      season: 2,
      episode: 3,
      title: "title",
   },
}

type test_alfa struct {
   show string
   season int
   episode int
   title string
   year int
}

type test_bravo struct {
   t test_alfa
}

func (t test_bravo) Show() string {
   return t.t.show
}

func (t test_bravo) Season() int {
   return t.t.season
}

func (t test_bravo) Episode() int {
   return t.t.episode
}

func (t test_bravo) Title() string {
   return t.t.title
}

func (t test_bravo) Year() int {
   return t.t.year
}

func TestName(t *testing.T) {
   for _, test := range name_tests {
      bravo := test_bravo{test}
      fmt.Printf("%q\n", Name(bravo))
   }
}
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
