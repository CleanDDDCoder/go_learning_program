package main

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	t.Run("creates logger with specified level", func(t *testing.T) {
		logger := NewLogger(LevelInfo)
		if logger == nil {
			t.Error("expected non-nil logger")
		}
	})
}

func TestLoggerLog(t *testing.T) {
	t.Run("logs only when level >= minimum", func(t *testing.T) {
		logger := NewLogger(LevelWarn)
		// Info should be skipped
		logger.Info("test info", nil)
		// Warn should be logged
		logger.Warn("test warn", nil)
	})
}

func TestLevelName(t *testing.T) {
	tests := []struct {
		level LogLevel
		want  string
	}{
		{LevelDebug, "DEBUG"},
		{LevelInfo, "INFO"},
		{LevelWarn, "WARN"},
		{LevelError, "ERROR"},
		{LevelFatal, "FATAL"},
		{99, "UNKNOWN"},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			got := LevelName(tt.level)
			if got != tt.want {
				t.Errorf("LevelName(%d) = %q, want %q", tt.level, got, tt.want)
			}
		})
	}
}