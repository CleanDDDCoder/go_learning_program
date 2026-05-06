package staticanalysis

import (
	"testing"
)

func TestNewAnalyzer(t *testing.T) {
	a := NewAnalyzer("test")
	if a == nil {
		t.Error("NewAnalyzer returned nil")
	}
	if a == nil || a.Name != "test" {
		t.Errorf("Expected name 'test', got '%s'", a.Name)
	}
}

func TestAnalyzerAnalyze(t *testing.T) {
	a := NewAnalyzer("demo")
	if a == nil {
		t.Fatal("NewAnalyzer returned nil")
	}
	result := a.Analyze()
	if result == "" {
		t.Error("Analyze returned empty string")
	}
}

func TestCheckCodeQuality(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected bool
	}{
		{"empty code", "", false},
		{"non-empty code", "package main", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CheckCodeQuality(tt.code)
			if result != tt.expected {
				t.Errorf("CheckCodeQuality(%q) = %v; want %v", tt.code, result, tt.expected)
			}
		})
	}
}

func TestDetectIssues(t *testing.T) {
	tests := []struct {
		name        string
		code        string
		expectEmpty bool
	}{
		{"empty code", "", false},
		{"normal code", "package main", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			issues := DetectIssues(tt.code)
			if tt.expectEmpty && len(issues) > 0 {
				t.Errorf("DetectIssues(%q) = %v; want empty", tt.code, issues)
			}
			if !tt.expectEmpty && len(issues) == 0 {
				t.Errorf("DetectIssues(%q) = empty; want issues", tt.code)
			}
		})
	}
}