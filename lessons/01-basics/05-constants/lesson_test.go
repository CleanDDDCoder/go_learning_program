package constants

import (
	"testing"
)

func TestDaysOfWeek(t *testing.T) {
	tests := []struct {
		name string
		day  int
		want string
	}{
		{"Sunday", 0, "Sunday"},
		{"Monday", 1, "Monday"},
		{"Tuesday", 2, "Tuesday"},
		{"Wednesday", 3, "Wednesday"},
		{"Thursday", 4, "Thursday"},
		{"Friday", 5, "Friday"},
		{"Saturday", 6, "Saturday"},
		{"invalid", 7, ""},
		{"negative", -1, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DaysOfWeek(tt.day)
			if got != tt.want {
				t.Errorf("DaysOfWeek(%d) = %q, want %q", tt.day, got, tt.want)
			}
		})
	}
}
