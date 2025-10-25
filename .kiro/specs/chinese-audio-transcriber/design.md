# Design Document

## Overview

The Chinese Audio Transcriber is a Windows desktop application built in Go that captures system audio in real-time and transcribes Mandarin Chinese speech to simplified Chinese text. The system uses a modular architecture with pluggable speech recognition services, starting with free options and allowing easy migration to paid services.

## Architecture

### High-Level Architecture

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│ Terminal Output │    │  Audio Capture   │    │ Speech Recognition│
│                 │    │     Module       │    │     Service     │
│ - Text Display  │◄───┤                  │◄───┤                 │
│ - Status Info   │    │ - WASAPI Capture │    │ - Free APIs     │
│ - Keyboard Input│    │ - Audio Buffering│    │ - OpenAI Whisper│
└─────────────────┘    └──────────────────┘    └─────────────────┘
         │                       │                       │
         ▼                       ▼                       ▼
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│  Configuration  │    │   Text Processor │    │   HTTP Client   │
│    Manager      │    │                  │    │                 │
│ - Config File   │    │ - Formatting     │    │ - API Requests  │
│ - API Keys      │    │ - Timestamps     │    │ - Error Handling│
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

## Components and Interfaces

### 1. Audio Capture Module

**Interface:**

```go
type AudioCapture interface {
    Start() error
    Stop() error
    GetAudioChunk() ([]byte, error)
    IsCapturing() bool
}
```

**Implementation Details:**

-   Uses Windows WASAPI through CGO bindings
-   Captures system audio at 16-bit PCM, 44.1kHz
-   Buffers audio in 2-3 second chunks
-   Handles device disconnection and reconnection
-   Runs in separate goroutine to avoid blocking

### 2. Speech Recognition Service

**Interface:**

```go
type SpeechRecognizer interface {
    Transcribe(audioData []byte, format AudioFormat) (string, error)
    SetLanguage(language string) error
    GetServiceName() string
}
```

**Free Service Priority:**

1. **Vosk (Offline)** - Local Chinese model if available
2. **SpeechRecognition API** - Browser-based free tier
3. **OpenAI Whisper API** - Free tier with API key

**Implementation Details:**

-   Factory pattern for service creation
-   Configurable service switching
-   Automatic fallback on service failure
-   Rate limiting and quota management

### 3. Terminal Interface

**Components:**

-   **Text Output**: Continuous display of transcribed text with timestamps
-   **Status Line**: Shows current status (🎧 Listening, ⚡ Processing, ❌ Error)
-   **Keyboard Controls**:
    -   Ctrl+C to exit
    -   Ctrl+L to clear screen
    -   Enter to pause/resume

**Features:**

-   Colored output for better readability
-   Real-time text streaming
-   Easy text selection and copying from terminal
-   Status indicators using emojis/symbols

### 4. Configuration Manager

**Configuration File (config.yaml):**

```yaml
audio:
    chunk_size_seconds: 3
    sample_rate: 44100
    bit_depth: 16

speech_recognition:
    primary_service: "vosk"
    fallback_service: "openai"
    language: "zh-CN"

api_keys:
    openai: ""

ui:
    colored_output: true
    show_timestamps: true
    clear_screen_on_start: false

logging:
    level: "info"
    file: "transcriber.log"
```

## Data Models

### Audio Data Structure

```go
type AudioChunk struct {
    Data      []byte
    Timestamp time.Time
    Duration  time.Duration
    Format    AudioFormat
}

type AudioFormat struct {
    SampleRate int
    BitDepth   int
    Channels   int
}
```

### Transcription Result

```go
type TranscriptionResult struct {
    Text      string
    Timestamp time.Time
    Confidence float64
    Service   string
}
```

### Application State

```go
type AppState struct {
    IsCapturing     bool
    CurrentService  string
    LastError       error
    TotalProcessed  int
    SessionStart    time.Time
}
```

## Error Handling

### Error Categories and Responses

1. **Audio Capture Errors**

    - Device not found: Retry every 5 seconds
    - Permission denied: Show user guidance
    - Format not supported: Attempt format conversion

2. **Network/API Errors**

    - Connection timeout: Retry with exponential backoff
    - Rate limit exceeded: Queue requests and wait
    - Invalid API key: Prompt user for configuration

3. **Processing Errors**
    - Audio format conversion failed: Log and skip chunk
    - Transcription service unavailable: Switch to fallback
    - Memory issues: Clear old transcriptions

### Error Recovery Strategy

-   Non-blocking error handling - continue operation when possible
-   User notification through status bar
-   Detailed logging for debugging
-   Graceful degradation (skip problematic audio chunks)

## Testing Strategy

### Unit Tests

-   Audio capture mock testing
-   Speech recognition service mocking
-   Configuration loading and validation
-   Text formatting and timestamp generation

### Integration Tests

-   End-to-end audio capture to text display
-   Service switching functionality
-   Error recovery scenarios
-   Configuration changes during runtime

### Manual Testing Scenarios

1. **Audio Sources**: Test with various audio (music, speech, mixed)
2. **Network Conditions**: Test with poor/no internet connectivity
3. **Extended Runtime**: 2+ hour continuous operation
4. **Service Switching**: Manual and automatic fallback testing
5. **Terminal Interaction**: Text copying, keyboard controls, colored output

## Performance Considerations

### Memory Management

-   Limit audio buffer size to prevent memory leaks
-   Clear old transcription results after threshold
-   Efficient string handling for Chinese text

### CPU Optimization

-   Audio processing in separate goroutines
-   Batch API requests when possible
-   Minimize terminal output frequency

### Network Efficiency

-   Compress audio data before API calls
-   Implement request queuing for rate limiting
-   Cache service responses when appropriate

## Security Considerations

### API Key Management

-   Store API keys in encrypted configuration
-   Never log API keys or responses
-   Validate API endpoints to prevent injection

### Audio Data Privacy

-   Process audio in memory only
-   No persistent audio storage
-   Clear audio buffers after processing

## Deployment and Distribution

### Build Requirements

-   Go 1.21+
-   CGO enabled for Windows WASAPI
-   Windows SDK for audio APIs

### Distribution Package

-   Single executable with embedded assets
-   Default configuration file
-   Installation instructions
-   Offline Chinese language model (optional)

### System Requirements

-   Windows 10/11
-   Minimum 4GB RAM
-   Internet connection for cloud APIs
-   Audio output device
