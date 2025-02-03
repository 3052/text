package http

import (
   "io"
   "net/http"
   "testing"
)

func TestParts(t *testing.T) {
   http.DefaultClient.Transport = nil
   var parts [9]struct{}
   var progress ProgressParts
   progress.Set(len(parts))
   for range parts {
      func() {
         resp, err := http.Get("http://httpbingo.org/drip?delay=0&duration=1")
         if err != nil {   
            t.Fatal(err)
         }
         defer resp.Body.Close()
         _, err = io.Copy(io.Discard, resp.Body)
         if err != nil {   
            t.Fatal(err)
         }
      }()
      progress.Next()
   }
}
