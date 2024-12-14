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

type test_bravo func() test_alfa

func (t test_bravo) Show() string {
   return t().show
}

func (t test_bravo) Season() int {
   return t().season
}

func (t test_bravo) Episode() int {
   return t().episode
}

func (t test_bravo) Title() string {
   return t().title
}

func (t test_bravo) Year() int {
   return t().year
}

func TestName(t *testing.T) {
   for _, test := range name_tests {
      var bravo test_bravo = func() test_alfa {
         return test
      }
      fmt.Printf("%q\n", Name(bravo))
   }
}
