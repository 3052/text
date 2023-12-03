package main

import "log/slog"

type hello struct{}

func (hello) String() string {
   return "hello world"
}

func main() {
   slog.Info("message", "key", hello{})
}
