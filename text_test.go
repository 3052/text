package text

import (
   "fmt"
   "testing"
)

var cardinal_tests = []struct{
   in Cardinal
   out string
}{
   {123.45, "123"},
   {123.45*1000, "123.45 thousand"},
   {123.45*1000*1000, "123.45 million"},
   {123.45*1000*1000*1000, "123.45 billion"},
}

func TestCardinal(t *testing.T) {
   for _, test := range cardinal_tests {
      out := fmt.Sprint(test.in)
      if out != test.out {
         t.Fatal(test)
      }
   }
}

func TestClean(t *testing.T) {
   out := Clean(`hello "*/:<>?\| world`)
   if out != "hello --------- world" {
      t.Fatal(out)
   }
}

func TestName(t *testing.T) {
   for _, test := range name_tests {
      out := Name(&test)
      if out != test.out {
         t.Fatal(test)
      }
   }
}

func (n *name_test) Show() string {
   return n.in.show
}

func (n *name_test) Season() int {
   return n.in.season
}

func (n *name_test) Episode() int {
   return n.in.episode
}

func (n *name_test) Title() string {
   return n.in.title
}

func (n *name_test) Year() int {
   return n.in.year
}

type name_zero struct {
   episode int
   season int
   show string
   title string
   year int
}

type name_test struct {
   in name_zero
   out string
}

var name_tests = []name_test{
   {
      in: name_zero{
         title: "title",
         year: 4,
      },
      out: "title - 4",
   },
   {
      in: name_zero{
         show: "show",
         season: 2,
         episode: 3,
         title: "title",
      },
      out: "show - 2 3 - title",
   },
   {
      in: name_zero{
         show: "show",
         episode: 3,
         title: "title",
      },
      out: "show - 3 - title",
   },
   {
      in: name_zero{
         show: "show",
         title: "title",
      },
      out: "show - title",
   },
}
