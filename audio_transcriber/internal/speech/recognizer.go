package speech

import (
	"chinese-audio-transcriber/internal/audio"
)

// TranscriptionResult represents the result of speech recognition
type TranscriptionResult struct {
	Text       string
	Timestamp  string
	Confidence float64
	Service    string
}

// SpeechRecognizer defines the interface for speech recognition services
type SpeechRecognizer interface {
	// Transcribe converts audio data to text
	Transcribe(audioChunk *audio.AudioChunk) (*TranscriptionResult, error)
	
	// SetLanguage configures the recognition language
	SetLanguage(language string) error
	
	// GetServiceName returns the name of the recognition service
	GetServiceName() string
	
	// IsAvailable checks if the service is currently available
	IsAvailable() bool
}

// RecognizerFactory creates speech recognizer instances
type RecognizerFactory interface {
	// CreateRecognizer creates a new recognizer instance
	CreateRecognizer(serviceName string) (SpeechRecognizer, error)
	
	// ListAvailableServices returns available recognition services
	ListAvailableServices() []string
}