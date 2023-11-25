package main

import (
	"bytes"
	"context"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type segment []byte

func transcodeVideoSegment(ctx context.Context, input <-chan segment, errChan chan<- error) (output chan segment) {
	output = make(chan segment, 10)
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(output)
				return
			case recv, ok := <-input:
				if ok {
					transcoded, err := transcodeToHLS(recv)
					if err != nil {
						errChan <- err
						continue
					}
					output <- transcoded
				}
			}
		}
	}()
	return output
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
