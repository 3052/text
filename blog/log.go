package main

import (
   "154.pages.dev/log"
   "flag"
   "log/slog"
   "net/http"
)

func main() {
   var lev slog.Level
   flag.TextVar(&lev, "level", lev, "level")
   flag.Parse()
   log.Set_Handler(lev)
   log.Set_Transport(slog.LevelInfo)
   slog.Info("hello world")
   http.Get("http://example.com")
   func() {
      log.Set_Transport(slog.LevelDebug)
      defer log.Set_Transport(slog.LevelInfo)
      slog.Info("hello world")
      http.Get("http://example.com")
   }()
   slog.Debug("hello world")
}
