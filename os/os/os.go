package main

import "41.neocities.org/x/os"

func main() {
   err := os.WriteFile("hello.txt", []byte("hello"))
   if err != nil {
      panic(err)
   }
}
