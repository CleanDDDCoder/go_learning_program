package meta

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	version := GetVersion()
	if version == "" {
		t.Error("GetVersion returned empty string")
	}
}

func TestGetCommit(t *testing.T) {
	commit := GetCommit()
	if commit == "" {
		t.Error("GetCommit returned empty string")
	}
}

func TestDefaultBuildInfo(t *testing.T) {
	info := DefaultBuildInfo()
	if info.Version == "" {
		t.Error("DefaultBuildInfo returned empty version")
	}
	if info.Commit == "" {
		t.Error("DefaultBuildInfo returned empty commit")
	}
}

func TestBuildInfoString(t *testing.T) {
	info := BuildInfo{
		Version:   "v1.0.0",
		Commit:   "abc123",
		BuildTime: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	str := info.String()
	if str == "" {
		t.Error("String returned empty string")
	}
}