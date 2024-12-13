package text

import (
   "fmt"
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

type test_bravo struct {
   t test_alfa
}

func TestName(t *testing.T) {
   for _, test := range name_tests {
      bravo := test_bravo{test}
      fmt.Printf("%q\n", Name(bravo))
   }
}
