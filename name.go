package text

import "strconv"

func Name(n Namer) string {
   var data []byte
   if n.Show() != "" {
      data = append(data, n.Show()...)
      data = append(data, " - "...)
      if n.Season() >= 1 {
         data = strconv.AppendInt(data, n.Season(), 10)
         data = append(data, ' ')
         data = strconv.AppendInt(data, n.Episode(), 10)
         if n.Title() != "" {
            data = append(data, " - "...)
            data = append(data, n.Title()...)
         }
      } else {
         if n.Episode() >= 1 {
            data = strconv.AppendInt(data, n.Episode(), 10)
            data = append(data, " - "...)
            data = append(data, n.Title()...)
         } else {
            data = append(data, n.Title()...)
         }
      }
   } else {
      data = append(data, n.Title()...)
      if n.Year() >= 1 {
         data = append(data, " - "...)
         data = strconv.AppendInt(data, n.Year(), 10)
      }
   }
   return string(data)
}

type Namer interface {
   Show() string
   Season() int64
   Episode() int64
   Title() string
   Year() int64
}
