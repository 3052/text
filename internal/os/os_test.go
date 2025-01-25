package os

import "testing"

func TestOs(t *testing.T) {
   err := write_file("hello.txt", []byte("hello"))
   if err != nil {
      t.Fatal(err)
   }
}
