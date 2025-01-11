package text

import (
   "fmt"
   "log"
   "net/http"
   "strings"
   "testing"
)

func TestTransport(t *testing.T) {
   var out strings.Builder
   log.SetOutput(&out)
   Transport{}.Set()
   _, err := http.Head("http://example.com")
   if err != nil {
      t.Fatal(err)
   }
   if !strings.HasSuffix(out.String(), " INFO HEAD URL=http://example.com\n") {
      t.Fatal(&out)
   }
}

func TestSize(t *testing.T) {
   for _, test := range size_tests {
      if fmt.Sprint(test.in) != test.out {
         t.Fatal(test)
      }
   }
}

func TestRate(t *testing.T) {
   for _, test := range rate_tests {
      if fmt.Sprint(test.in) != test.out {
         t.Fatal(test)
      }
   }
}

var rate_tests = []struct{
   in Rate
   out string
}{
   {123.45, "123 byte/s"},
   {123.45*1000, "123.45 kilobyte/s"},
   {123.45*1000*1000, "123.45 megabyte/s"},
   {123.45*1000*1000*1000, "123.45 gigabyte/s"},
}

var size_tests = []struct{
   in Size
   out string
}{
   {123.45, "123 byte"},
   {123.45*1000, "123.45 kilobyte"},
   {123.45*1000*1000, "123.45 megabyte"},
   {123.45*1000*1000*1000, "123.45 gigabyte"},
}

func TestCardinal(t *testing.T) {
   for _, test := range cardinal_tests {
      if fmt.Sprint(test.in) != test.out {
         t.Fatal(test)
      }
   }
}

var cardinal_tests = []struct{
   in Cardinal
   out string
}{
   {123.45, "123"},
   {123.45*1000, "123.45 thousand"},
   {123.45*1000*1000, "123.45 million"},
   {123.45*1000*1000*1000, "123.45 billion"},
}

var percent_tests = []struct{
   in Percent
   out string
}{
   {0.0123, "1.23 %"},
   {0.1234, "12.34 %"},
}

func TestPercent(t *testing.T) {
   for _, test := range percent_tests {
      if fmt.Sprint(test.in) != test.out {
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
