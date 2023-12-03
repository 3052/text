package main

import (
   "154.pages.dev/http/blog"
   "flag"
   "log/slog"
)

func main() {
   var h blog.Handler
   flag.TextVar(&h.Level, "level", h.Level, "level")
   flag.Parse()
   slog.SetDefault(slog.New(h))
   blog.SetTransport(slog.LevelInfo)
   get()
   slog.Debug("hello world")
}
