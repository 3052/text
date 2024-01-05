package log

import (
   "io"
   "net/http"
   "log/slog"
   "testing"
)

const address = "https://go.dev/dl/go1.21.5.windows-amd64.zip"

func Test_Progress(t *testing.T) {
   Set_Transport(slog.LevelInfo)
   res, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   r := New_Progress(1).Reader(res)
   io.Copy(io.Discard, r)
}
