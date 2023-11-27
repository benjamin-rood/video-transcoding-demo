package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
