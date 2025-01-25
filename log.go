package log

import (
   "log"
   "strconv"
)

func init() {
   log.SetFlags(log.Ltime)
}

func (p Percent) String() string {
   unit := unit_measure{100, " %"}
   return label(float64(p), unit)
}

func label(value float64, unit unit_measure) string {
   var prec int
   if unit.factor != 1 {
      prec = 2
      value *= unit.factor
   }
   return strconv.FormatFloat(value, 'f', prec, 64) + unit.name
}

type Cardinal float64

func (c Cardinal) String() string {
   units := []unit_measure{
      {1, ""},
      {1e-3, " thousand"},
      {1e-6, " million"},
      {1e-9, " billion"},
   }
   return scale(float64(c), units)
}

type Rate float64

func (r Rate) String() string {
   units := []unit_measure{
      {1, " byte/s"},
      {1e-3, " kilobyte/s"},
      {1e-6, " megabyte/s"},
      {1e-9, " gigabyte/s"},
   }
   return scale(float64(r), units)
}

type Percent float64

type Size float64

func (s Size) String() string {
   units := []unit_measure{
      {1, " byte"},
      {1e-3, " kilobyte"},
      {1e-6, " megabyte"},
      {1e-9, " gigabyte"},
   }
   return scale(float64(s), units)
}

type unit_measure struct {
   factor float64
   name string
}

func scale(value float64, units []unit_measure) string {
   var unit unit_measure
   for _, unit = range units {
      if unit.factor * value < 1000 {
         break
      }
   }
   return label(value, unit)
}
