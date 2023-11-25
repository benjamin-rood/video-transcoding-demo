package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

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

func streamVideoToServer(inputVideoDir string, client vspb.StreamingVideoIngestorClient) error {
	files, err := os.ReadDir(inputVideoDir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		if err := streamVideoSegmentToServer(filepath.Join(inputVideoDir, file.Name()), client); err != nil {
			return fmt.Errorf("Failed to stream file %s: %w", file.Name(), err)
		}
	}
	return nil
}

func streamVideoSegmentToServer(inputVideoPath string, client vspb.StreamingVideoIngestorClient) error {
	file, err := os.Open(inputVideoPath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Create a stream for uploading the file.
	stream, err := client.UploadVideo(context.Background())
	if err != nil {
		log.Fatalf("failed to open stream: %v", err)
	}

	buf := make([]byte, chunkSize)
	isFirstChunk := true
	for {
		// Read the file in chunks and send them to the server.
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return fmt.Errorf("failed to read file: %w", err)
		}
		if n == 0 {
			// no bytes read, done
			break
		}
		chunk := buf[:n]
		if err := stream.Send(&vspb.VideoChunk{
			Data:         chunk,
			SegmentStart: isFirstChunk,
			SegmentEnd:   err == io.EOF,
		}); err != nil {
			log.Fatalf("failed to send chunk: %s", err)
		}
		isFirstChunk = false

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
	var videoDirPath string
	flag.StringVar(&videoDirPath, "dir", "", "Path to the directory containing video segments")
	flag.Parse()

	// Ensure the directory path is provided
	if videoDirPath == "" {
		log.Fatal("Directory path is required")
	}
	// Set up a connection to the server (using insecure because this is not real)
	conn, err := grpc.Dial(":59999", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := vspb.NewStreamingVideoIngestorClient(conn)
	if err = streamVideoToServer(videoDirPath, client); err != nil {
		// Handle error based on gRPC status
		// For example, you can use status.FromError(err) to get the status code and message
	}

}
