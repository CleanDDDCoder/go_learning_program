package struct_fields

import "testing"

func TestSummary(t *testing.T) {
	tests := []struct {
		name   string
		person Person
		want   string
	}{
		{name: "adult", person: Person{Name: "Ada", Age: 36}, want: "Ada is 36"},
		{name: "child", person: Person{Name: "Sam", Age: 8}, want: "Sam is 8"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Summary(tt.person)
			if got != tt.want {
				t.Fatalf("Summary(%+v) = %q, want %q", tt.person, got, tt.want)
			}
		})
	}
}
