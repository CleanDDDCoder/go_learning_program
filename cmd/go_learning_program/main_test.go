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

func TestParseCapstoneYAML(t *testing.T) {
	data := []byte(strings.Join([]string{
		`title: Production HTTP Service`,
		`difficulty: expert`,
		`run_modes:`,
		`  - test`,
		`  - race`,
		`optional: true`,
		`estimated_scope: portfolio-project`,
	}, "\n"))

	got := parseCapstoneYAML(data)
	if got.Title != "Production HTTP Service" {
		t.Fatalf("Title = %q, want Production HTTP Service", got.Title)
	}
	if got.Difficulty != "expert" {
		t.Fatalf("Difficulty = %q, want expert", got.Difficulty)
	}
	if got.EstimatedScope != "portfolio-project" {
		t.Fatalf("EstimatedScope = %q, want portfolio-project", got.EstimatedScope)
	}
	if !got.Optional {
		t.Fatal("Optional = false, want true")
	}
	if len(got.RunModes) != 2 || got.RunModes[0] != "test" || got.RunModes[1] != "race" {
		t.Fatalf("RunModes = %#v, want [test race]", got.RunModes)
	}
}

func TestLoadCapstonesReadsYAMLMetadata(t *testing.T) {
	root := t.TempDir()
	capstoneDir := root + string(os.PathSeparator) + "capstones" + string(os.PathSeparator) + "03-production-http-service"
	if err := os.MkdirAll(capstoneDir, 0o755); err != nil {
		t.Fatal(err)
	}
	metadata := []byte(strings.Join([]string{
		`title: Production HTTP Service`,
		`difficulty: expert`,
		`run_modes:`,
		`  - test`,
		`optional: false`,
		`estimated_scope: multi-session`,
	}, "\n"))
	if err := os.WriteFile(capstoneDir+string(os.PathSeparator)+"capstone.yaml", metadata, 0o600); err != nil {
		t.Fatal(err)
	}

	got, err := loadCapstones(root)
	if err != nil {
		t.Fatalf("loadCapstones() error = %v", err)
	}
	if len(got) != 1 {
		t.Fatalf("len(loadCapstones()) = %d, want 1", len(got))
	}
	if got[0].Path != "03-production-http-service" {
		t.Fatalf("Path = %q, want 03-production-http-service", got[0].Path)
	}
	if got[0].Title != "Production HTTP Service" {
		t.Fatalf("Title = %q, want Production HTTP Service", got[0].Title)
	}
	if got[0].Difficulty != "expert" {
		t.Fatalf("Difficulty = %q, want expert", got[0].Difficulty)
	}
	if len(got[0].RunModes) != 1 || got[0].RunModes[0] != "test" {
		t.Fatalf("RunModes = %#v, want [test]", got[0].RunModes)
	}
	if got[0].Optional {
		t.Fatal("Optional = true, want false")
	}
}
