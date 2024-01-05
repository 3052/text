package main

import (
   "154.pages.dev/log"
   "flag"
   "log/slog"
)

func main() {
   var h log.Handler
   flag.TextVar(&h.Level, "v", h.Level, "level")
   flag.Parse()
   log.Set_Logger(h.Level)
   slog.Debug("hello world")
}
