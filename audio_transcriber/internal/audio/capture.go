package audio

import (
	"time"
)

// AudioFormat represents the format of audio data
type AudioFormat struct {
	SampleRate int
	BitDepth   int
	Channels   int
}

// AudioChunk represents a chunk of audio data with metadata
type AudioChunk struct {
	Data      []byte
	Timestamp time.Time
	Duration  time.Duration
	Format    AudioFormat
}

// AudioCapture defines the interface for capturing system audio
type AudioCapture interface {
	// Start begins audio capture
	Start() error
	
	// Stop ends audio capture
	Stop() error
	
	// GetAudioChunk returns the next available audio chunk
	GetAudioChunk() (*AudioChunk, error)
	
	// IsCapturing returns true if currently capturing audio
	IsCapturing() bool
	
	// SetFormat configures the audio capture format
	SetFormat(format AudioFormat) error
}