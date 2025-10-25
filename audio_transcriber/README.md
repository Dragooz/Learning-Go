# Chinese Audio Transcriber

## Project Goal

Build a real-time audio transcription tool that continuously listens to Windows system audio and transcribes Mandarin Chinese speech to text for transcription work.

## Requirements

### Core Functionality

-   **Real-time audio capture**: Continuously listen to Windows system audio (not microphone)
-   **Live transcription**: Process audio chunks and transcribe to Mandarin Chinese text in real-time
-   **Copy-paste friendly**: Display transcribed text in a format easy to copy for transcription work
-   **Persistent operation**: Keep running until manually stopped

### Technical Specifications

-   **Language**: Go
-   **Target OS**: Windows
-   **Audio Source**: System audio output (what's playing on speakers/headphones)
-   **Language Focus**: Mandarin Chinese only
-   **Output Format**: Simplified Chinese characters
-   **Processing**: Real-time/streaming (not batch)

### User Experience

-   **Interface**: CLI tool or simple GUI - whatever works best
-   **Operation**: Start program → automatically begins listening and transcribing
-   **Output**: Continuous text display with timestamps
-   **Interaction**: Easy text selection and copying

## Architecture Plan

### Components Needed

1. **Audio Capture Module**

    - Windows audio loopback capture (WASAPI)
    - Audio format handling (PCM, sample rate conversion)
    - Chunking for real-time processing

2. **Speech Recognition Module**

    - Free tier API integration (start with OpenAI Whisper API free tier)
    - Modular design for easy API switching
    - Chinese language optimization

3. **Text Processing Module**

    - Format transcribed text with timestamps
    - Handle text buffering and display
    - Copy-to-clipboard functionality

4. **User Interface**
    - Real-time text display
    - Status indicators (listening, processing, error states)
    - Basic controls (start/stop, clear text)

### Technology Stack

-   **Audio Capture**: Windows WASAPI (via CGO or system calls)
-   **HTTP Client**: Standard Go net/http for API calls
-   **Speech Recognition**: OpenAI Whisper API (free tier)
-   **Fallback Options**: Google Speech-to-Text, Azure Speech Services
-   **UI**: Terminal-based initially, with option for simple GUI later

## Implementation Rules

### Code Organization

-   Modular design with clear separation of concerns
-   Interface-based design for easy API provider switching
-   Configuration file for API keys and settings
-   Proper error handling and logging

### Performance Requirements

-   **Latency**: < 3 seconds from audio to text display
-   **Memory**: Efficient audio buffer management
-   **CPU**: Minimal impact on system performance
-   **Network**: Batch API calls when possible to reduce requests

### Error Handling

-   Network connectivity issues
-   API rate limiting and quota management
-   Audio device disconnection/changes
-   Invalid audio format handling

## Testing Strategy

### Unit Tests

-   Audio capture functionality
-   API client modules
-   Text processing and formatting
-   Configuration management

### Integration Tests

-   End-to-end audio capture to text output
-   API provider switching
-   Error recovery scenarios

### Manual Testing

-   Real-world audio sources (videos, music, speech)
-   Different audio qualities and volumes
-   Extended runtime stability
-   Copy-paste workflow validation

## Development Phases

### Phase 1: Core Audio Capture

-   Implement Windows audio loopback
-   Basic audio format handling
-   Test with simple audio output

### Phase 2: Speech Recognition Integration

-   OpenAI Whisper API integration
-   Chinese language configuration
-   Basic text output

### Phase 3: Real-time Processing

-   Streaming audio chunks
-   Continuous transcription
-   Text display with timestamps

### Phase 4: User Experience

-   Improved text formatting
-   Copy-paste optimization
-   Error handling and status display

### Phase 5: Enhancement

-   Configuration options
-   Multiple API provider support
-   Performance optimization

## Success Criteria

-   Successfully captures Windows system audio
-   Transcribes Mandarin speech with reasonable accuracy
-   Displays text in real-time (< 3 second delay)
-   Runs continuously without crashes
-   Easy to copy transcribed text for work use
-   Uses free APIs within reasonable limits

## Future Considerations

-   Paid API integration for better accuracy/limits
-   Support for other Chinese dialects
-   Audio file transcription mode
-   Text export functionality
-   Simple GUI interface
