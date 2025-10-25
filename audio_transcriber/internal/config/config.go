package config

import (
	"fmt"
	"os"
	"path/filepath"
	
	"gopkg.in/yaml.v3"
)

// AudioConfig holds audio capture configuration
type AudioConfig struct {
	ChunkSizeSeconds int `yaml:"chunk_size_seconds"`
	SampleRate       int `yaml:"sample_rate"`
	BitDepth         int `yaml:"bit_depth"`
	Channels         int `yaml:"channels"`
}

// SpeechRecognitionConfig holds speech recognition settings
type SpeechRecognitionConfig struct {
	PrimaryService  string `yaml:"primary_service"`
	FallbackService string `yaml:"fallback_service"`
	Language        string `yaml:"language"`
}

// APIKeysConfig holds API authentication keys
type APIKeysConfig struct {
	OpenAI string `yaml:"openai"`
}

// UIConfig holds user interface settings
type UIConfig struct {
	ColoredOutput      bool `yaml:"colored_output"`
	ShowTimestamps     bool `yaml:"show_timestamps"`
	ClearScreenOnStart bool `yaml:"clear_screen_on_start"`
}

// LoggingConfig holds logging configuration
type LoggingConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

// Config represents the complete application configuration
type Config struct {
	Audio             AudioConfig             `yaml:"audio"`
	SpeechRecognition SpeechRecognitionConfig `yaml:"speech_recognition"`
	APIKeys           APIKeysConfig           `yaml:"api_keys"`
	UI                UIConfig                `yaml:"ui"`
	Logging           LoggingConfig           `yaml:"logging"`
}

// ConfigManager defines the interface for configuration management
type ConfigManager interface {
	// LoadConfig loads configuration from file
	LoadConfig(configPath string) (*Config, error)
	
	// SaveConfig saves configuration to file
	SaveConfig(config *Config, configPath string) error
	
	// ValidateConfig validates the configuration
	ValidateConfig(config *Config) error
	
	// GetDefaultConfig returns default configuration
	GetDefaultConfig() *Config
}

// DefaultConfigManager implements ConfigManager
type DefaultConfigManager struct{}

// NewConfigManager creates a new configuration manager
func NewConfigManager() ConfigManager {
	return &DefaultConfigManager{}
}

// GetDefaultConfig returns the default configuration
func (cm *DefaultConfigManager) GetDefaultConfig() *Config {
	return &Config{
		Audio: AudioConfig{
			ChunkSizeSeconds: 3,
			SampleRate:       44100,
			BitDepth:         16,
			Channels:         2,
		},
		SpeechRecognition: SpeechRecognitionConfig{
			PrimaryService:  "vosk",
			FallbackService: "openai",
			Language:        "zh-CN",
		},
		APIKeys: APIKeysConfig{
			OpenAI: "",
		},
		UI: UIConfig{
			ColoredOutput:      true,
			ShowTimestamps:     true,
			ClearScreenOnStart: false,
		},
		Logging: LoggingConfig{
			Level: "info",
			File:  "transcriber.log",
		},
	}
}

// LoadConfig loads configuration from the specified file
func (cm *DefaultConfigManager) LoadConfig(configPath string) (*Config, error) {
	// If config file doesn't exist, create default config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		defaultConfig := cm.GetDefaultConfig()
		if err := cm.SaveConfig(defaultConfig, configPath); err != nil {
			return nil, fmt.Errorf("failed to create default config: %w", err)
		}
		return defaultConfig, nil
	}
	
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	
	if err := cm.ValidateConfig(&config); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}
	
	return &config, nil
}

// SaveConfig saves configuration to the specified file
func (cm *DefaultConfigManager) SaveConfig(config *Config, configPath string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}
	
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}
	
	return nil
}

// ValidateConfig validates the configuration
func (cm *DefaultConfigManager) ValidateConfig(config *Config) error {
	if config.Audio.ChunkSizeSeconds <= 0 {
		return fmt.Errorf("audio chunk size must be positive")
	}
	
	if config.Audio.SampleRate <= 0 {
		return fmt.Errorf("audio sample rate must be positive")
	}
	
	if config.Audio.BitDepth != 16 && config.Audio.BitDepth != 24 && config.Audio.BitDepth != 32 {
		return fmt.Errorf("audio bit depth must be 16, 24, or 32")
	}
	
	if config.Audio.Channels <= 0 {
		return fmt.Errorf("audio channels must be positive")
	}
	
	if config.SpeechRecognition.Language == "" {
		return fmt.Errorf("speech recognition language cannot be empty")
	}
	
	return nil
}