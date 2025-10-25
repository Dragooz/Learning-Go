package app

import (
	"chinese-audio-transcriber/internal/audio"
	"chinese-audio-transcriber/internal/config"
	"chinese-audio-transcriber/internal/speech"
	"context"
	"fmt"
	"time"
)

// AppState represents the current state of the application
type AppState struct {
	IsCapturing    bool
	CurrentService string
	LastError      error
	TotalProcessed int
	SessionStart   time.Time
}

// Application represents the main application
type Application struct {
	config          *config.Config
	configManager   config.ConfigManager
	audioCapture    audio.AudioCapture
	speechRecognizer speech.SpeechRecognizer
	state           *AppState
	ctx             context.Context
	cancel          context.CancelFunc
}

// NewApplication creates a new application instance
func NewApplication() *Application {
	ctx, cancel := context.WithCancel(context.Background())
	
	return &Application{
		configManager: config.NewConfigManager(),
		state: &AppState{
			SessionStart: time.Now(),
		},
		ctx:    ctx,
		cancel: cancel,
	}
}

// Initialize initializes the application with configuration
func (app *Application) Initialize(configPath string) error {
	// Load configuration
	cfg, err := app.configManager.LoadConfig(configPath)
	if err != nil {
		return fmt.Errorf("failed to load configuration: %w", err)
	}
	app.config = cfg
	
	return nil
}

// Start starts the application
func (app *Application) Start() error {
	if app.config == nil {
		return fmt.Errorf("application not initialized")
	}
	
	app.state.IsCapturing = true
	app.state.SessionStart = time.Now()
	
	// TODO: Initialize audio capture and speech recognition
	// This will be implemented in subsequent tasks
	
	return nil
}

// Stop stops the application
func (app *Application) Stop() error {
	app.cancel()
	app.state.IsCapturing = false
	
	if app.audioCapture != nil {
		return app.audioCapture.Stop()
	}
	
	return nil
}

// GetState returns the current application state
func (app *Application) GetState() *AppState {
	return app.state
}

// GetConfig returns the current configuration
func (app *Application) GetConfig() *config.Config {
	return app.config
}