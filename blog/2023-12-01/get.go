package main

import (
   "154.pages.dev/http/blog"
   "log/slog"
   "net/http"
)

func get() {
   slog.Info("hello world")
   http.Get("http://example.com")
   blog.SetTransport(slog.LevelDebug)
   defer blog.SetTransport(slog.LevelInfo)
   slog.Info("hello world")
   http.Get("http://example.com")
}
