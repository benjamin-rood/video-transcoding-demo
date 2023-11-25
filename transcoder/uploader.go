package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	vspb "github.com/benjamin-rood/video-transcoding-demo/proto"
)

type VideoInjestor struct {
	vspb.UnsafeStreamingVideoIngestorServer
	ctx            context.Context
	transcodeQueue chan<- segment
	errs           chan<- error
}

func NewVideoInjestor(ctx context.Context, transcodeQueue chan<- segment, errChan chan<- error) VideoInjestor {
	return VideoInjestor{
		ctx:            ctx,
		transcodeQueue: transcodeQueue,
		errs:           errChan,
	}
}

func (vi *VideoInjestor) UploadVideo(stream vspb.StreamingVideoIngestor_UploadVideoServer) error {
	currentSegment := segment{}
	totalBytesReceived := int64(0)
	for {
		select {
		case <-vi.ctx.Done():
			return nil
		default:
			chunk, err := stream.Recv()
			if err == io.EOF {
				// stream completed
				break
			}
			if err != nil {
				// Is it better to just return an error instead of sending it to an error handler over a channel?
				vi.errs <- status.Errorf(codes.Internal, "Error while receiving data: %v", err)
				continue
			}
			if len(currentSegment) == 0 && !chunk.SegmentStart {
				// Is it better to just return an error instead of sending it to an error handler over a channel?
				vi.errs <- status.Errorf(codes.Internal, "expected new segment to begin with `segment_start` flag enabled")
				continue
			}
			totalBytesReceived += int64(len(chunk.GetData()))
			fmt.Printf("\rreceived %v bytes", totalBytesReceived)
			currentSegment = append(currentSegment, chunk.Data...)
			if chunk.SegmentEnd {
				vi.transcodeQueue <- currentSegment
				// clear segment buffer
				currentSegment = segment{}
				// reset bytes received counter
				totalBytesReceived = 0
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ln, err := net.Listen("tcp", ":59999")
	if err != nil {
		log.Fatalf("could not initialise tcp listener: %s", err)
	}
	defer ln.Close()
	ctx, cancel := context.WithCancel(context.Background())
	// make a buffered queue, we don't want to block receiving uploaded chunks
	transcodeQueue := make(chan segment, 10)
	errChan := make(chan error)
	injestionService := NewVideoInjestor(ctx, transcodeQueue, errChan)
	server := grpc.NewServer()
	vspb.RegisterStreamingVideoIngestorServer(server, &injestionService)
	outputQueue := transcodeVideoSegment(ctx, transcodeQueue, errChan)
	wg.Add(1)
	go func() {
		for i := range <-outputQueue {
			// TODO: we need to send the transcoded segments somewhere
			log.Printf("successfully transcoded segment %v", i)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case err := <-errChan:
				log.Println(err)
				cancel()
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	if err := server.Serve(ln); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
