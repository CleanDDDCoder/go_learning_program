// Package crosscompile demonstrates cross-compilation techniques
package crosscompile

import (
	"fmt"
	"runtime"
)

// TargetInfo returns information about the compilation target
func TargetInfo() string {
	return fmt.Sprintf("OS: %s, Arch: %s", runtime.GOOS, runtime.GOARCH)
}

// IsWindows returns true if running on Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux returns true if running on Linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsDarwin returns true if running on macOS
func IsDarwin() bool {
	return runtime.GOOS == "darwin"
}

func init() {
	fmt.Println("crosscompile package initialized")
}