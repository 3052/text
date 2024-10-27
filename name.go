package text

import (
   "strings"
   "text/template"
)

func Clean(s string) string {
   mapping := func(r rune) rune {
      if strings.ContainsRune(`"*/:<>?\|`, r) {
         return '-'
      }
      return r
   }
   return strings.Map(mapping, s)
}

func Name(n Namer) (string, error) {
   text, err := new(template.Template).Parse(DefaultName)
   if err != nil {
      return "", err
   }
   var b strings.Builder
   err = text.Execute(&b, n)
   if err != nil {
      return "", err
   }
   return b.String(), nil
}

type Namer interface {
   Show() string
   Season() int
   Episode() int
   Title() string
   Year() int
}

var DefaultName =
   "{{if .Show}}" +
      "{{if .Season}}" +
         "{{if .Title}}" +
            "{{.Show}} - {{.Season}} {{.Episode}} - {{.Title}}" +
         "{{else}}" +
            "{{.Show}} - {{.Season}} {{.Episode}}" +
         "{{end}}" +
      "{{else}}" +
         "{{.Show}} - {{.Title}}" +
      "{{end}}" +
   "{{else}}" +
      "{{if .Year}}" +
         "{{.Title}} - {{.Year}}" +
      "{{else}}" +
         "{{.Title}}" +
      "{{end}}" +
   "{{end}}"
