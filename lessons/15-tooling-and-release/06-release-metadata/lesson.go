// Package meta demonstrates release metadata embedding
package meta

import (
	"fmt"
	"time"
)

// BuildInfo holds version and build information
type BuildInfo struct {
	Version   string
	Commit    string
	BuildTime time.Time
}

// DefaultBuildInfo returns default build information
func DefaultBuildInfo() BuildInfo {
	return BuildInfo{
		Version:   "v0.0.0",
		Commit:    "dev",
		BuildTime: time.Now(),
	}
}

// GetVersion returns the version string
func GetVersion() string {
	return DefaultBuildInfo().Version
}

// GetCommit returns the commit hash
func GetCommit() string {
	return DefaultBuildInfo().Commit
}

// String returns a string representation of BuildInfo
func (b BuildInfo) String() string {
	return fmt.Sprintf("Version: %s, Commit: %s, Built: %s", b.Version, b.Commit, b.BuildTime.Format(time.RFC3339))
}

func init() {
	fmt.Println("meta package initialized")
}