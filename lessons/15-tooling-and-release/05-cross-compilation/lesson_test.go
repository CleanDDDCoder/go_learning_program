package crosscompile

import (
	"testing"
)

func TestTargetInfo(t *testing.T) {
	info := TargetInfo()
	if info == "" {
		t.Error("TargetInfo returned empty string")
	}
}

func TestIsWindows(t *testing.T) {
	result := IsWindows()
	if result && runtime.GOOS != "windows" {
		t.Error("IsWindows returned true on non-Windows platform")
	}
}

func TestIsLinux(t *testing.T) {
	result := IsLinux()
	if result && runtime.GOOS != "linux" {
		t.Error("IsLinux returned true on non-Linux platform")
	}
}

func TestIsDarwin(t *testing.T) {
	result := IsDarwin()
	if result && runtime.GOOS != "darwin" {
		t.Error("IsDarwin returned true on non-Darwin platform")
	}
}