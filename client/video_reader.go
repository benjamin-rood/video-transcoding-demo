package main

import (
	"context"
	"io"
	"log"
	"os"

	vspb "github.com/benjamin-rood/video-transcoding-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// play around with different chunk sizes
	kb        = 1024
	mb        = kb * kb
	chunkSize = 256 * kb // Upload chunks of 256KB
)

func uploadVideo(client vspb.StreamingVideoIngestorClient, filepath string) error {
	// TODO: Open the video file to read from
	// Open the file to be uploaded.
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	// Create a stream for uploading the file.
	stream, err := client.UploadVideo(context.Background())
	if err != nil {
		log.Fatalf("failed to open stream: %v", err)
	}

	for {
		// Read the file in chunks and send them to the server.
		buf := make([]byte, chunkSize)
		n, err := file.Read(buf)
		if err != nil {
			log.Println(err)
		}
		if err != nil && err != io.EOF {
			log.Fatalf("failed to read file: %s", err)
		}
		if n == 0 {
			break
		}
		chunk := buf[:n]
		if err := stream.Send(&vspb.VideoChunk{
			Data: chunk,
		}); err != nil {
			log.Fatalf("failed to send chunk: %s", err)
		}

		// Simulate connection issues by randomly sleeping between bursts.
		// sleepTime := rand.Intn(175) + 25 // Sleep for 25-200ms.
		// time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	}

	if _, err = stream.CloseAndRecv(); err != nil {
		return err
	}

	log.Println("Upload finished successfully")
	return nil
}

func main() {
	videoFilepath := ""
	// Set up a connection to the server (using insecure because this is not real)
	conn, err := grpc.Dial(":59999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := vspb.NewStreamingVideoIngestorClient(conn)
	if err = uploadVideo(client, videoFilepath); err != nil {
		// Handle error based on gRPC status
		// For example, you can use status.FromError(err) to get the status code and message
	}

}
