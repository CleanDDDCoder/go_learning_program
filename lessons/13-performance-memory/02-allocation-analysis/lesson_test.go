package allocation

import (
	"testing"
)

func TestFilterEmpty(t *testing.T) {
	tests := []struct {
		name     string
		items    []string
		wantRes  []string
	}{
		{"empty slice", []string{}, []string{}},
		{"all empty", []string{"", "", ""}, []string{}},
		{"no empties", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"mixed", []string{"a", "", "b", "", "c"}, []string{"a", "b", "c"}},
		{"single empty", []string{""}, []string{}},
		{"single non-empty", []string{"x"}, []string{"x"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterEmpty(tt.items)
			if len(got) != len(tt.wantRes) {
				t.Errorf("FilterEmpty(%v) len = %d, want %d", tt.items, len(got), len(tt.wantRes))
				return
			}
			for i := range got {
				if got[i] != tt.wantRes[i] {
					t.Errorf("FilterEmpty(%v)[%d] = %q, want %q", tt.items, i, got[i], tt.wantRes[i])
				}
			}
		})
	}
}

func TestConcat(t *testing.T) {
	tests := []struct {
		name    string
		items   []string
		sep     string
		wantRes string
	}{
		{"empty slice", []string{}, ",", ""},
		{"single item", []string{"a"}, ",", "a"},
		{"two items", []string{"a", "b"}, ",", "a,b"},
		{"three items", []string{"a", "b", "c"}, "-", "a-b-c"},
		{"empty sep", []string{"a", "b", "c"}, "", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Concat(tt.items, tt.sep)
			if got != tt.wantRes {
				t.Errorf("Concat(%v, %q) = %q, want %q", tt.items, tt.sep, got, tt.wantRes)
			}
		})
	}
}