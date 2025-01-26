package os

import "testing"

func Test(t *testing.T) {
   err := WriteFile("hello.txt", []byte("hello"))
   if err != nil {
      t.Fatal(err)
   }
}
