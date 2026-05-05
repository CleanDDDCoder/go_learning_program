package test_helpers

import (
	"fmt"
	"strconv"
)

// ParsePort converts text to a TCP port number.
func ParsePort(text string) (int, error) {
	port, err := strconv.Atoi(text)
	if err != nil {
		return 0, fmt.Errorf("parse port: %w", err)
	}
	if port < 1 || port > 65535 {
		return 0, fmt.Errorf("port %d out of range", port)
	}
	return port, nil
}
