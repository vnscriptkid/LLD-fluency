package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

// LogLevel type
type LogLevel int

// Enum for log levels
const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

// LogMessage now includes Data for structured logging
type LogMessage struct {
	Level     LogLevel
	Timestamp time.Time
	Message   string
	Data      map[string]interface{} // For structured data
}

// LogWriter interface for different output destinations
type LogWriter interface {
	Write(message LogMessage) error
}

// Logger struct now needs to handle structured logging
type Logger struct {
	level   LogLevel
	writers []LogWriter
	mu      sync.Mutex // Ensures thread safety
}

var (
	once     sync.Once
	instance *Logger // Singleton instance
)

// NewLogger remains the same.

// log method is extended to accept structured data
func (l *Logger) log(level LogLevel, msg string, data map[string]interface{}) {
	// In different environments, we may want to log only messages of a certain level or higher.
	// Example:
	// Stage: DEBUG, INFO, WARN, ERROR, FATAL
	// Prod: INFO, WARN, ERROR, FATAL
	if level < l.level {
		return
	}
	message := LogMessage{
		Level:     level,
		Timestamp: time.Now(),
		Message:   msg,
		Data:      data,
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, w := range l.writers {
		if err := w.Write(message); err != nil {
			log.Printf("Error writing log message: %v\n", err)
		}
	}
}

// Logging methods are modified to support structured logging
func (l *Logger) Debug(msg string, data map[string]interface{}) { l.log(DEBUG, msg, data) }
func (l *Logger) Info(msg string, data map[string]interface{})  { l.log(INFO, msg, data) }
func (l *Logger) Warn(msg string, data map[string]interface{})  { l.log(WARN, msg, data) }
func (l *Logger) Error(msg string, data map[string]interface{}) { l.log(ERROR, msg, data) }
func (l *Logger) Fatal(msg string, data map[string]interface{}) { l.log(FATAL, msg, data) }

// StdoutWriter Write method to support structured logging
type StdoutWriter struct{}

func (s StdoutWriter) Write(m LogMessage) error {
	// Convert structured data to JSON string
	dataString, err := json.Marshal(m.Data)
	if err != nil {
		return err
	}
	_, err = fmt.Printf("%s [%v] %s %s\n", m.Timestamp.Format(time.RFC3339), m.Level, m.Message, string(dataString))
	return err
}

// GetInstance returns the singleton instance of the Logger.
func GetInstance() *Logger {
	// even if multiple goroutines call Do simultaneously.
	// The first call to Do executes the function, while subsequent calls do nothing.
	// Creating a singleton instance of Logger just needs to be done once.
	once.Do(func() {
		instance = &Logger{
			level: DEBUG,
			writers: []LogWriter{
				StdoutWriter{},
			},
		} // Initialize with default values or configuration
	})
	return instance
}
