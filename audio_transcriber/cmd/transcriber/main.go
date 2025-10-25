package main

import (
	"chinese-audio-transcriber/internal/app"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
)

const (
	defaultConfigFile = "config.yaml"
)

func main() {
	// Parse command line flags
	var configPath string
	flag.StringVar(&configPath, "config", defaultConfigFile, "Path to configuration file")
	flag.Parse()
	
	// Make config path absolute
	if !filepath.IsAbs(configPath) {
		execDir, err := os.Executable()
		if err != nil {
			log.Fatalf("Failed to get executable directory: %v", err)
		}
		configPath = filepath.Join(filepath.Dir(execDir), configPath)
	}
	
	// Create and initialize application
	application := app.NewApplication()
	if err := application.Initialize(configPath); err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}
	
	// Set up signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	// Start the application
	fmt.Println("Starting Chinese Audio Transcriber...")
	fmt.Printf("Configuration loaded from: %s\n", configPath)
	
	if err := application.Start(); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
	
	fmt.Println("Application started successfully. Press Ctrl+C to stop.")
	
	// Wait for shutdown signal
	<-sigChan
	
	fmt.Println("\nShutting down...")
	if err := application.Stop(); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
	
	fmt.Println("Application stopped.")
}