package text

import "fmt"

type Namer interface {
   Show() string
   Season() int
   Episode() int
   Title() string
   Year() int
}

func Name(n Namer) string {
   var data []byte
   if n.Show() != "" {
      data = fmt.Append(data, n.Show(), " - ")
      if n.Season() >= 1 {
         data = fmt.Append(data, n.Season(), " ", n.Episode())
         if n.Title() != "" {
            data = fmt.Append(data, " - ", n.Title())
         }
      } else {
         if n.Episode() >= 1 {
            data = fmt.Append(data, n.Episode(), " - ", n.Title())
         } else {
            data = append(data, n.Title()...)
         }
      }
   } else {
      data = append(data, n.Title()...)
      if n.Year() >= 1 {
         data = fmt.Append(data, " - ", n.Year())
      }
   }
   return string(data)
}
