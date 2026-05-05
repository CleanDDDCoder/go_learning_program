package main

import (
	"time"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// LogEntry represents a structured log entry
type LogEntry struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
	Fields    map[string]interface{}
}

// Logger provides structured logging capabilities
type Logger struct {
	level LogLevel
}

// NewLogger creates a new Logger with the specified minimum level
func NewLogger(minLevel LogLevel) *Logger {
	return &Logger{level: minLevel}
}

// Log creates a log entry with the given level and message
func (l *Logger) Log(level LogLevel, msg string, fields map[string]interface{}) {
	if level < l.level {
		return
	}
	entry := LogEntry{
		Timestamp: time.Now(),
		Level:     level,
		Message:   msg,
		Fields:    fields,
	}
	// TODO: Implement actual logging output
	_ = entry
}

// Debug logs a debug message
func (l *Logger) Debug(msg string, fields map[string]interface{}) {
	l.Log(LevelDebug, msg, fields)
}

// Info logs an info message
func (l *Logger) Info(msg string, fields map[string]interface{}) {
	l.Log(LevelInfo, msg, fields)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string, fields map[string]interface{}) {
	l.Log(LevelWarn, msg, fields)
}

// Error logs an error message
func (l *Logger) Error(msg string, fields map[string]interface{}) {
	l.Log(LevelError, msg, fields)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(msg string, fields map[string]interface{}) {
	l.Log(LevelFatal, msg, fields)
}

// LevelName returns the string representation of a log level
func LevelName(level LogLevel) string {
	switch level {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

func main() {}