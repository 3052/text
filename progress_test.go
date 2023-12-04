package log

import (
   "io"
   "net/http"
   "testing"
)

const address = "https://http-a-darwin.hulustream.com/183/196861183/stream_018b1aa0-d60b-618d-356f-0035e43c0fcf_1000128640668_H264_2500_1000128655856_video.mp4?authToken=1701739147_8791ba024cff470def777352c3ce94f8"

func Test_Progress(t *testing.T) {
   res, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   r := New_Progress(1).Reader(res)
   Set_Handler(0)
   io.Copy(io.Discard, r)
}
