package text

import (
   "fmt"
   "testing"
)

func (t test_bravo) Show() string {
   return t.t.show
}

func (t test_bravo) Season() int64 {
   return t.t.season
}

func (t test_bravo) Episode() int64 {
   return t.t.episode
}

func (t test_bravo) Title() string {
   return t.t.title
}

func (t test_bravo) Year() int64 {
   return t.t.year
}

type test_bravo struct {
   t test_alfa
}

type test_alfa struct {
   show string
   season int64
   episode int64
   title string
   year int64
}

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

func TestName(t *testing.T) {
   for _, test := range name_tests {
      bravo := test_bravo{test}
      fmt.Printf("%q\n", Name(bravo))
   }
}
