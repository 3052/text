package log

import (
   "log/slog"
   "testing"
)

func Test_Log(t *testing.T) {
   Set_Logger(slog.LevelDebug)
   slog.Debug("hello world")
}
