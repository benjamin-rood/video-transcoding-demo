package main

import (
	"bytes"
	"context"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type segment []byte

func transcodeVideoSegment(ctx context.Context, input <-chan segment, output chan<- segment, errChan chan<- error) {
	for {
		select {
		case <-ctx.Done():
			return
		case recv, ok := <-input:
			if ok {
				transcoded, err := transcodeToHLS(recv)
				if err != nil {
					errChan <- err
					continue
				}
				output <- transcoded
			} else {
				return
			}
		}
	}
}

func transcodeToHLS(videoSegment segment) (segment, error) {
	buf := &bytes.Buffer{}
	err := ffmpeg.Input("pipe:0").
		Output("pipe:1", ffmpeg.KwArgs{"format": "hls"}).
		WithInput(bytes.NewReader(videoSegment)).
		WithOutput(buf).
		Run()
	return buf.Bytes(), err
}
