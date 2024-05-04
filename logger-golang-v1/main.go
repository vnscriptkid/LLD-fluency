package main

import "github.com/vnscriptkid/LLD-fluency/logger-golang-v1/logger"

// Example usage with structured logging
func main() {
	// logger := NewLogger(INFO, StdoutWriter{})
	logger := logger.GetInstance()
	logger.Info("User logged in", map[string]interface{}{"user_id": 1234, "ip": "192.168.1.1"})
}
