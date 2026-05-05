package embedding

import "testing"

func TestEmployeeLabel(t *testing.T) {
	employee := Employee{
		Contact: Contact{Name: "Grace"},
		Role:    "Engineer",
	}

	got := employee.Label()
	if got != "Grace - Engineer" {
		t.Fatalf("Label() = %q, want %q", got, "Grace - Engineer")
	}
}
