package text

import (
   "testing"
   "text/template"
)

func TestName(t *testing.T) {
   _, err := new(template.Template).Parse(DefaultName)
   if err != nil {
      t.Fatal(err)
   }
}
