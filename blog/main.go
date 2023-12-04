package main

import (
   "154.pages.dev/log"
   "flag"
   "log/slog"
)

func main() {
   var l slog.Level
   flag.TextVar(&l, "level", l, "level")
   flag.Parse()
   log.SetHandler(l)
   log.SetTransport(0)
   get()
   slog.Debug("hello world")
}
