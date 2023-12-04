package main

import (
   "154.pages.dev/log"
   "log/slog"
   "net/http"
)

func get() {
   slog.Info("hello world")
   http.Get("http://example.com")
   log.SetTransport(slog.LevelDebug)
   defer log.SetTransport(slog.LevelInfo)
   slog.Info("hello world")
   http.Get("http://example.com")
}
