package os

import (
   "log"
   "os"
   _ "41.neocities.org/log"
)

func WriteFile(name string, data []byte) error {
   log.Println("WriteFile", name)
   return os.WriteFile(name, data, os.ModePerm)
}
