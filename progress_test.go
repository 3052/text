package slog

import (
   "io"
   "net/http"
   "testing"
)

const address = "https://http-a-darwin.hulustream.com/183/196861183/stream_018b1aa0-d60b-618d-356f-0035e43c0fcf_1000128640662_H264_6000_1000128648587_video.mp4?authToken=1701729250_5dedf21efdbf94aa4fc7cf853d458791"

func Test_Progress(t *testing.T) {
   res, err := http.Get(address)
   if err != nil {
      t.Fatal(err)
   }
   defer res.Body.Close()
   r := Progress_Length(res.ContentLength).Reader(res)
   SetHandler(0)
   io.Copy(io.Discard, r)
}
