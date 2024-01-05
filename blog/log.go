package main

import (
   "154.pages.dev/log"
   "flag"
   "log/slog"
)

func main() {
   var level log.Level
   flag.TextVar(&level, "v", level, "level")
   flag.Parse()
   log.Set_Logger(level)
   slog.Debug("hello world")
}
