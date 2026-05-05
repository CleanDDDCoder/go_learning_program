package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

type lesson struct {
	Path  string
	Title string
}

type testEvent struct {
	Action  string
	Package string
	Test    string
	Output  string
}

type testSummary struct {
	Passed        bool
	CompileError  bool
	FirstFailing  string
	FailureCount  int
	DiagnosticOut string
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout io.Writer, stderr io.Writer) int {
	if len(args) == 0 {
		printUsage(stderr)
		return 2
	}

	root, err := findProjectRoot()
	if err != nil {
		fmt.Fprintf(stderr, "find project root: %v\n", err)
		return 1
	}
	lessons, err := loadCurriculum(filepath.Join(root, "curriculum.yaml"))
	if err != nil {
		fmt.Fprintf(stderr, "load curriculum: %v\n", err)
		return 1
	}

	switch args[0] {
	case "list":
		return runList(lessons, stdout)
	case "hint":
		return runHint(root, args[1:], stdout, stderr)
	case "test":
		return runTest(root, args[1:], lessons, stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown command %q\n", args[0])
		printUsage(stderr)
		return 2
	}
}

func printUsage(w io.Writer) {
	fmt.Fprintln(w, "usage:")
	fmt.Fprintln(w, "  golearn list")
	fmt.Fprintln(w, "  golearn test <lesson>")
	fmt.Fprintln(w, "  golearn hint <lesson> [--level N]")
}

func runList(lessons []lesson, stdout io.Writer) int {
	for _, item := range lessons {
		fmt.Fprintf(stdout, "%s - %s\n", strings.TrimPrefix(item.Path, "lessons/"), item.Title)
	}
	return 0
}

func runHint(root string, args []string, stdout io.Writer, stderr io.Writer) int {
	lessonPath, level, err := parseHintArgs(args)
	if err != nil {
		fmt.Fprintf(stderr, "%v\n", err)
		return 2
	}

	hints, err := loadHints(filepath.Join(root, normalizeLessonPath(lessonPath), "hints.yaml"))
	if err != nil {
		fmt.Fprintf(stderr, "load hints: %v\n", err)
		return 1
	}
	for _, hint := range hints {
		if level != 0 && hint.level != level {
			continue
		}
		fmt.Fprintf(stdout, "%d. %s\n", hint.level, hint.text)
	}
	return 0
}

func parseHintArgs(args []string) (string, int, error) {
	var lessonPath string
	var level int
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--level":
			if i+1 >= len(args) {
				return "", 0, errors.New("hint --level requires a number")
			}
			parsed, err := strconv.Atoi(args[i+1])
			if err != nil {
				return "", 0, fmt.Errorf("hint --level: %w", err)
			}
			level = parsed
			i++
		case strings.HasPrefix(arg, "--level="):
			parsed, err := strconv.Atoi(strings.TrimPrefix(arg, "--level="))
			if err != nil {
				return "", 0, fmt.Errorf("hint --level: %w", err)
			}
			level = parsed
		case strings.HasPrefix(arg, "-"):
			return "", 0, fmt.Errorf("unknown hint flag %q", arg)
		default:
			if lessonPath != "" {
				return "", 0, errors.New("hint requires exactly one lesson path")
			}
			lessonPath = arg
		}
	}
	if lessonPath == "" {
		return "", 0, errors.New("hint requires exactly one lesson path")
	}
	return lessonPath, level, nil
}

func runTest(root string, args []string, lessons []lesson, stdout io.Writer, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "test requires exactly one lesson path")
		return 2
	}

	lessonPath := normalizeLessonPath(args[0])
	summary, err := runGoTestJSON(root, lessonPath)
	if err != nil && summary.DiagnosticOut == "" {
		fmt.Fprintf(stderr, "run go test: %v\n", err)
		return 1
	}
	if summary.Passed {
		fmt.Fprintf(stdout, "PASS %s\n", strings.TrimPrefix(lessonPath, "lessons/"))
		if next, ok := nextLesson(lessons, lessonPath); ok {
			fmt.Fprintf(stdout, "Next lesson: %s\n", strings.TrimPrefix(next.Path, "lessons/"))
		} else {
			fmt.Fprintln(stdout, "You've completed the curriculum. Nice work.")
		}
		return 0
	}
	if summary.CompileError {
		fmt.Fprintf(stdout, "COMPILE ERROR %s\n%s", strings.TrimPrefix(lessonPath, "lessons/"), summary.DiagnosticOut)
		return 1
	}
	if summary.FirstFailing != "" {
		fmt.Fprintf(stdout, "FAIL %s: %s failed (%d failing tests)\n", strings.TrimPrefix(lessonPath, "lessons/"), summary.FirstFailing, summary.FailureCount)
		return 1
	}
	fmt.Fprintf(stdout, "FAIL %s\n%s", strings.TrimPrefix(lessonPath, "lessons/"), summary.DiagnosticOut)
	return 1
}

func findProjectRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "curriculum.yaml")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("curriculum.yaml not found")
		}
		dir = parent
	}
}

func loadCurriculum(path string) ([]lesson, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lessons []lesson
	var pending *lesson
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "- path:") {
			value := strings.TrimSpace(strings.TrimPrefix(line, "- path:"))
			lessons = append(lessons, lesson{Path: filepath.ToSlash(filepath.Clean(value))})
			pending = &lessons[len(lessons)-1]
			continue
		}
		if pending != nil && strings.HasPrefix(line, "title:") {
			pending.Title = strings.Trim(strings.TrimSpace(strings.TrimPrefix(line, "title:")), `"`)
			pending = nil
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(lessons) == 0 {
		return nil, errors.New("no lessons listed")
	}
	return lessons, nil
}

type hint struct {
	level int
	text  string
}

func loadHints(path string) ([]hint, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var hints []hint
	var pending *hint
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "- hint:") {
			text := strings.Trim(strings.TrimSpace(strings.TrimPrefix(line, "- hint:")), `"`)
			hints = append(hints, hint{level: len(hints) + 1, text: text})
			pending = &hints[len(hints)-1]
			continue
		}
		if strings.HasPrefix(line, "- ") && !strings.HasPrefix(line, "- hint:") {
			text := strings.Trim(strings.TrimSpace(strings.TrimPrefix(line, "- ")), `"`)
			hints = append(hints, hint{level: len(hints) + 1, text: text})
			pending = nil
			continue
		}
		if pending != nil && strings.HasPrefix(line, "level:") {
			level, err := strconv.Atoi(strings.TrimSpace(strings.TrimPrefix(line, "level:")))
			if err != nil {
				return nil, err
			}
			pending.level = level
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(hints) == 0 {
		return nil, errors.New("no hints found")
	}
	return hints, nil
}

func normalizeLessonPath(value string) string {
	clean := filepath.ToSlash(filepath.Clean(filepath.FromSlash(value)))
	clean = strings.TrimPrefix(clean, "./")
	if !strings.HasPrefix(clean, "lessons/") {
		clean = "lessons/" + clean
	}
	return clean
}

func nextLesson(lessons []lesson, path string) (lesson, bool) {
	for i, item := range lessons {
		if item.Path == path && i+1 < len(lessons) {
			return lessons[i+1], true
		}
	}
	return lesson{}, false
}

func runGoTestJSON(root string, lessonPath string) (testSummary, error) {
	cmd := exec.Command("go", "test", "-json", "./"+lessonPath)
	cmd.Dir = root
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	err := cmd.Run()
	summary := parseGoTestJSON(output.Bytes())
	if err == nil {
		summary.Passed = true
	}
	return summary, err
}

func parseGoTestJSON(data []byte) testSummary {
	var summary testSummary
	var diagnostics strings.Builder
	failedTests := map[string]bool{}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		var event testEvent
		if err := json.Unmarshal(scanner.Bytes(), &event); err != nil {
			continue
		}
		if event.Output != "" {
			diagnostics.WriteString(event.Output)
		}
		if event.Action == "build-fail" {
			summary.CompileError = true
		}
		if event.Action == "fail" && event.Test != "" {
			if summary.FirstFailing == "" {
				summary.FirstFailing = event.Test
			}
			if !failedTests[event.Test] {
				failedTests[event.Test] = true
				summary.FailureCount++
			}
		}
	}
	if summary.CompileError && summary.FirstFailing == "" {
		summary.DiagnosticOut = diagnostics.String()
	}
	if summary.FailureCount == 0 && !summary.CompileError {
		summary.DiagnosticOut = diagnostics.String()
	}
	return summary
}
