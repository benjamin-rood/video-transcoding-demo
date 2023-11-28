package main

import (
	"testing"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	mock_vspb "github.com/benjamin-rood/video-transcoding-demo/mocks"
	vspb "github.com/benjamin-rood/video-transcoding-demo/proto"
	"github.com/spf13/afero"
	"go.uber.org/mock/gomock"
)

func TestExtractSegmentNumber(t *testing.T) {
	testCases := []struct {
		filename string
		expected uint32
	}{
		{"big.buck.bunny.demo.1080p.60fps.000105.ts", 105},
		{"another.file.000003.m4v", 3},
		{"file with whitespace.000000.mkv", 0},
		{"12345", 12345},
	}

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			result := extractSegmentNumber(tc.filename)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestExtractSegmentNumberPanic(t *testing.T) {
	testCases := []struct {
		filename    string
		shouldPanic bool
	}{
		{"file.with.no.segment.number.ts", true},
		{"", true},
		{"file.with.negative.segment.number.-001198.ts", true},
	}

	for _, tc := range testCases {
		assert.Panics(t, func() { extractSegmentNumber(tc.filename) }, "The code did not panic")
	}
}

func TestStreamVideoSegmentToServer(t *testing.T) {
	oldFs := fs
	oldChunkSize := chunkSize
	fs = afero.NewMemMapFs()
	chunkSize = 5
	// Make sure to restore the original values of appfs and chunkSize after the test
	defer func() {
		fs = oldFs
		chunkSize = oldChunkSize
	}()

	p1 := "something.000103.ts"
	p2 := "something.000104.ts"
	afero.WriteFile(fs, p1, []byte("content"), 0644)
	afero.WriteFile(fs, p2, []byte("content"), 0644)

	testCases := []struct {
		filename      string
		segmentNumber uint32
	}{
		{p1, uint32(103)},
		{p2, uint32(104)},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mock_vspb.NewMockStreamingVideoIngestorClient(ctrl)
	mockStream := mock_vspb.NewMockStreamingVideoIngestor_UploadVideoClient(ctrl)

	for _, tc := range testCases {
		t.Run(tc.filename, func(t *testing.T) {
			mockClient.EXPECT().UploadVideo(gomock.Any(), gomock.Any()).Return(mockStream, nil)

			// Create the expected messages
			expectedFirstMessage := &vspb.VideoChunk{
				VideoId:       "videoID",
				SegmentNumber: tc.segmentNumber,
				SegmentStart:  true,
				SegmentEnd:    false,
				Data:          []byte("conte"),
			}

			expectedSecondMessage := &vspb.VideoChunk{
				VideoId:       "videoID",
				SegmentNumber: tc.segmentNumber,
				SegmentStart:  false,
				SegmentEnd:    true,
				Data:          []byte("nt"),
			}

			// Use gomock.Eq to match the expected message values
			mockStream.EXPECT().Send(gomock.Eq(expectedFirstMessage)).Return(nil).Times(1)
			mockStream.EXPECT().Send(gomock.Eq(expectedSecondMessage)).Return(nil).Times(1)
			mockStream.EXPECT().CloseAndRecv().Return(&empty.Empty{}, nil)

			err := streamVideoSegmentToServer("videoID", tc.filename, mockClient)
			require.NoError(t, err)
		})
	}
}
