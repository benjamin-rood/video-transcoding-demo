package main

import (
	"context"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	vspb "github.com/benjamin-rood/video-transcoding-demo/proto"
)

type VideoInjestor struct {
	vspb.UnsafeStreamingVideoIngestorServer
	ctx            context.Context
	transcodeQueue chan<- segment
}

func NewVideoInjestor(ctx context.Context, transcodeQueue chan<- segment) *VideoInjestor {
	return &VideoInjestor{
		ctx:            ctx,
		transcodeQueue: transcodeQueue,
	}
}

func (vi *VideoInjestor) UploadVideo(stream vspb.StreamingVideoIngestor_UploadVideoServer) error {
	defer stream.SendAndClose(&empty.Empty{})
	currentSegment := segment{}
	for {
		select {
		case <-vi.ctx.Done():
			return nil
		default:
			chunk, err := stream.Recv()
			if err == io.EOF {
				// stream completed
				return nil
			}
			if err != nil {
				return status.Errorf(codes.Internal, "Error while receiving data: %v", err)
			}
			if len(currentSegment) == 0 && !chunk.SegmentStart {
				return status.Errorf(codes.Internal, "expected new segment to begin with `segment_start` flag enabled")
			}
			currentSegment = append(currentSegment, chunk.Data...)
			if chunk.SegmentEnd {
				vi.transcodeQueue <- currentSegment
				// clear segment buffer
				currentSegment = segment{}
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	go handleSigint(cancel)
	ln, err := net.Listen("tcp", ":59999")
	if err != nil {
		log.Fatalf("could not initialise tcp listener: %s", err)
	}
	defer ln.Close()
	errChan := make(chan error)
	// make a buffered queue, we don't want to block receiving uploaded chunks
	inputQueue := make(chan segment, 10)
	outputQueue := make(chan segment, 10)
	go transcodeVideoSegment(ctx, inputQueue, outputQueue, errChan)
	go func() {
		transcodedCounter := 1
		for range outputQueue {
			// TODO: we need to send the transcoded segments somewhere
			log.Printf("\rsuccessfully transcoded segment %07d", transcodedCounter)
			transcodedCounter++
		}
	}()

	server := grpc.NewServer()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case err := <-errChan:
				log.Println(err)
				cancel()
			case <-ctx.Done():
				close(inputQueue)
				close(outputQueue)
				close(errChan)
				server.Stop()
				return
			}
		}
	}()

	vspb.RegisterStreamingVideoIngestorServer(server, NewVideoInjestor(ctx, inputQueue))
	if err := server.Serve(ln); err != nil {
		cancel()
		log.Fatal(err)
	}
	wg.Wait()
}

func handleSigint(cancel context.CancelFunc) {
	// setup signal handling to cancel everything if we receive a signal
	sigChan := make(chan os.Signal, 1)
	// os.Interrupt is a synonym for syscall.SIGINT and guaranteed on
	// all OS's, SIGTERM is used by kubernetes and gives you 30 seconds
	// to shutdown (timeout can be configured through the
	// terminationGracePeriodSeconds, field in your pod specification.
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	// wait on a signal
	<-sigChan
	// log the signal
	log.Println("received signal interruption")
	// cancel all running
	cancel()
}
