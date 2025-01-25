package os

import (
   "log"
   "os"
)

func WriteFile(name string, data []byte) error {
   log.Println("WriteFile", name)
   return os.WriteFile(name, data, os.ModePerm)
}
