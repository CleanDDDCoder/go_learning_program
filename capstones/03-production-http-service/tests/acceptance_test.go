package tests

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCapstoneAuthoringFilesExist(t *testing.T) {
	root := filepath.Dir(mustGetwd(t))
	required := []string{
		"README.md",
		"capstone.yaml",
		"starter/service.go",
		"fixtures/widget.json",
		"rubric.yaml",
		"maintainer_notes.md",
	}
	for _, name := range required {
		if _, err := os.Stat(filepath.Join(root, name)); err != nil {
			t.Fatalf("required capstone file %s: %v", name, err)
		}
	}
}

func mustGetwd(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	return dir
}
