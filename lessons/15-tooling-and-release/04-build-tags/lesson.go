// Package buildtags demonstrates Go build tags
//go:build linux && amd64
// +build linux,amd64

package buildtags

import (
	"fmt"
	"runtime"
)

// GetPlatformInfo returns platform-specific information
func GetPlatformInfo() string {
	return fmt.Sprintf("Running on %s/%s", runtime.GOOS, runtime.GOARCH)
}

// IsLinuxAMD64 returns true if running on Linux AMD64
func IsLinuxAMD64() bool {
	return runtime.GOOS == "linux" && runtime.GOARCH == "amd64"
}

func init() {
	fmt.Println("buildtags package initialized")
}