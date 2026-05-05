package buildtags

import (
	"runtime"
	"testing"
)

func TestGetPlatformInfo(t *testing.T) {
	info := GetPlatformInfo()
	if info == "" {
		t.Error("GetPlatformInfo returned empty string")
	}
}

func TestIsLinuxAMD64(t *testing.T) {
	result := IsLinuxAMD64()
	if result && (runtime.GOOS != "linux" || runtime.GOARCH != "amd64") {
		t.Error("IsLinuxAMD64 returned true on non-Linux AMD64 platform")
	}
}