package test_helpers

import "testing"

func TestParsePort(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{name: "valid", input: "8080", want: 8080, wantErr: false},
		{name: "not numeric", input: "http", want: 0, wantErr: true},
		{name: "too low", input: "0", want: 0, wantErr: true},
		{name: "too high", input: "70000", want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePort(tt.input)
			assertErrState(t, err, tt.wantErr)
			if got != tt.want {
				t.Fatalf("ParsePort(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func assertErrState(t *testing.T, err error, wantErr bool) {
	t.Helper()
	if wantErr && err == nil {
		t.Fatal("error = nil, want non-nil")
	}
	if !wantErr && err != nil {
		t.Fatalf("error = %v, want nil", err)
	}
}
