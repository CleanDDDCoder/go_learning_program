package word_count

import "testing"

func TestCountWords(t *testing.T) {
	counts := CountWords("go test go")
	if counts["go"] != 2 {
		t.Fatalf("counts[go] = %d, want 2", counts["go"])
	}
	if counts["test"] != 1 {
		t.Fatalf("counts[test] = %d, want 1", counts["test"])
	}
	if len(CountWords("")) != 0 {
		t.Fatal("CountWords(empty) should return an empty map")
	}
}
