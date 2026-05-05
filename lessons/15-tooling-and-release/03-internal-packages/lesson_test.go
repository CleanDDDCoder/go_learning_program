package internalpackages

import (
	"testing"
)

func TestGetInternalDemo(t *testing.T) {
	result := GetInternalDemo()
	if result == "" {
		t.Error("GetInternalDemo returned empty string")
	}
}

func TestGetServiceName(t *testing.T) {
	name := GetServiceName()
	if name == "" {
		t.Error("GetServiceName returned empty string")
	}
}