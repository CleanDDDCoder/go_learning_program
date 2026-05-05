package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestRunListWritesLessons(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"list"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("run list exit = %d, stderr = %s", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "01-basics/01-hello-world - Hello World") {
		t.Fatalf("list output missing first lesson: %s", stdout.String())
	}
}

func TestRunHintFiltersByLevel(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run([]string{"hint", "01-basics/02-variables", "--level", "2"}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("run hint exit = %d, stderr = %s", code, stderr.String())
	}
	got := stdout.String()
	if !strings.Contains(got, "2.") {
		t.Fatalf("hint output did not include selected level: %s", got)
	}
	if strings.Contains(got, "1.") || strings.Contains(got, "3.") {
		t.Fatalf("hint output included unselected levels: %s", got)
	}
}

func TestNormalizeLessonPath(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{name: "bare forward slash path", in: "01-basics/01-hello-world", want: "lessons/01-basics/01-hello-world"},
		{name: "already rooted", in: "lessons/01-basics/01-hello-world", want: "lessons/01-basics/01-hello-world"},
		{name: "dot prefix", in: "./01-basics/01-hello-world", want: "lessons/01-basics/01-hello-world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeLessonPath(tt.in)
			if got != tt.want {
				t.Fatalf("normalizeLessonPath(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestParseGoTestJSONFailure(t *testing.T) {
	data := []byte(strings.Join([]string{
		`{"Action":"run","Package":"example","Test":"TestAdd"}`,
		`{"Action":"fail","Package":"example","Test":"TestAdd"}`,
		`{"Action":"fail","Package":"example","Test":"TestMultiply"}`,
	}, "\n"))

	got := parseGoTestJSON(data)
	if got.FirstFailing != "TestAdd" {
		t.Fatalf("FirstFailing = %q, want TestAdd", got.FirstFailing)
	}
	if got.FailureCount != 2 {
		t.Fatalf("FailureCount = %d, want 2", got.FailureCount)
	}
}

func TestParseGoTestJSONCompileError(t *testing.T) {
	data := []byte(strings.Join([]string{
		`{"Action":"output","Package":"example","Output":"lesson.go:3: missing return\n"}`,
		`{"Action":"build-fail","Package":"example"}`,
	}, "\n"))

	got := parseGoTestJSON(data)
	if !got.CompileError {
		t.Fatal("CompileError = false, want true")
	}
	if !strings.Contains(got.DiagnosticOut, "missing return") {
		t.Fatalf("DiagnosticOut = %q, want compile diagnostic", got.DiagnosticOut)
	}
}

func TestRunTestShowsEndOfCurriculum(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping subprocess test in short mode")
	}

	root, err := findProjectRoot()
	if err != nil {
		t.Fatalf("findProjectRoot() error = %v", err)
	}
	lessons, err := loadCurriculum(root + string(os.PathSeparator) + "curriculum.yaml")
	if err != nil {
		t.Fatalf("loadCurriculum() error = %v", err)
	}
	lastLesson := strings.TrimPrefix(lessons[len(lessons)-1].Path, "lessons/")
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	code := run([]string{"test", lastLesson}, &stdout, &stderr)
	if code != 0 {
		t.Fatalf("run test exit = %d, stderr = %s", code, stderr.String())
	}
	if !strings.Contains(stdout.String(), "PASS "+lastLesson) {
		t.Fatalf("test output missing pass summary: %s", stdout.String())
	}
	if !strings.Contains(stdout.String(), "You've completed the curriculum. Nice work.") {
		t.Fatalf("test output missing final curriculum message: %s", stdout.String())
	}
}

func TestLoadCurriculumRequiresLessons(t *testing.T) {
	dir := t.TempDir()
	path := dir + string(os.PathSeparator) + "curriculum.yaml"
	if err := os.WriteFile(path, []byte("curriculum:\n  name: empty\n"), 0o600); err != nil {
		t.Fatal(err)
	}

	_, err := loadCurriculum(path)
	if err == nil {
		t.Fatal("loadCurriculum returned nil error for empty curriculum")
	}
}
