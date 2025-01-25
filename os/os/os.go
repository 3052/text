package main

import "41.neocities.org/log/os"

func main() {
   err := os.WriteFile("hello.txt", []byte("hello"))
   if err != nil {
      panic(err)
   }
}
