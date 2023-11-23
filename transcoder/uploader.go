package main

import (
	"io"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	vspb "github.com/benjamin-rood/video-transcoding-demo/proto"
)

/**
 * Embed `UnsafeStreamingVideoIngestorServer` instead of `UnimplementedStreamingVideoIngestorServer`
 * to ensure we get compilation errors unless service is defined.
 * Good practice, even if it's a single rpc VideoInjestor.
 */
type VideoInjestor struct {
	vspb.UnsafeStreamingVideoIngestorServer
}

// Check interface conformity
var _ vspb.StreamingVideoIngestorServer = &VideoInjestor{}

func (vi *VideoInjestor) UploadVideo(stream vspb.StreamingVideoIngestor_UploadVideoServer) error {
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			// stream completed
			return nil
		}
		// TODO: content type inspection?
		if err != nil {
			return status.Errorf(codes.Internal, "Error while receiving data: %v", err)
		}

		// TODO: send chunk for processing
	}
}
