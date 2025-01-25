package os

import (
   "log"
   "os"
)

func write_file(name string, data []byte) error {
   log.Print(name)
   return os.WriteFile(name, data, os.ModePerm)
}
