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
	"time"
)

type lesson struct {
	Path              string   `yaml:"path"`
	Title             string   `yaml:"title"`
	Difficulty        string   `yaml:"difficulty"`
	LessonType        string   `yaml:"lesson_type"`
	RunModes          []string `yaml:"run_modes"`
	Optional          bool     `yaml:"optional"`
	Prerequisites     []string `yaml:"prerequisites"`
	RecommendedPrereq []string `yaml:"recommended_prerequisites"`
	MinGoVersion      string   `yaml:"min_go_version"`
}

type track struct {
	ID         string
	Difficulty string
	Lessons    []lesson
	Capstones  []lesson
}

type capstone struct {
	Path           string   `yaml:"path"`
	Title          string   `yaml:"title"`
	Difficulty     string   `yaml:"difficulty"`
	EstimatedScope string   `yaml:"estimated_scope"`
	RunModes       []string `yaml:"run_modes"`
	Optional       bool     `yaml:"optional"`
}

type testEvent struct {
	Action  string `json:"Action"`
	Package string `json:"Package"`
	Test    string `json:"Test"`
	Output  string `json:"Output"`
	Time    string `json:"Time"`
}

type testSummary struct {
	Passed        bool
	CompileError  bool
	RaceError     bool
	BuildError    bool
	FirstFailing  string
	FailureCount  int
	DiagnosticOut string
}

type progressState struct {
	SkippedOptional map[string]bool `json:"skipped_optional"`
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

	lessons, tracks, err := loadAllCurriculum(root)
	if err != nil {
		fmt.Fprintf(stderr, "load curriculum: %v\n", err)
		return 1
	}

	capstones, err := loadCapstones(root)
	if err != nil {
		fmt.Fprintf(stderr, "load capstones: %v\n", err)
		return 1
	}

	progress, err := loadProgress(root)
	if err != nil && !os.IsNotExist(err) {
		fmt.Fprintf(stderr, "load progress: %v\n", err)
		return 1
	}

	switch args[0] {
	case "list":
		return runList(lessons, stdout)
	case "hint":
		return runHint(root, args[1:], stdout, stderr)
	case "test":
		return runTest(root, args[1:], lessons, stdout, stderr)
	case "tracks":
		return runTracks(tracks, capstones, stdout)
	case "next":
		return runNext(root, lessons, capstones, progress, stdout, stderr)
	case "progress":
		return runProgress(root, lessons, capstones, progress, stdout, stderr)
	case "diagnose":
		return runDiagnose(root, args[1:], lessons, stdout, stderr)
	case "race":
		return runRace(root, args[1:], lessons, stdout, stderr)
	case "bench":
		return runBench(root, args[1:], lessons, stdout, stderr)
	case "fuzz":
		return runFuzz(root, args[1:], lessons, stdout, stderr)
	case "profile":
		return runProfile(root, args[1:], lessons, stdout, stderr)
	case "review":
		return runReview(root, args[1:], capstones, stdout, stderr)
	case "skip":
		return runSkip(root, args[1:], lessons, capstones, progress, stdout, stderr)
	case "unskip":
		return runUnskip(root, args[1:], lessons, capstones, progress, stdout, stderr)
	case "reset":
		return runReset(root, args[1:], lessons, stdout, stderr)
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
	fmt.Fprintln(w, "  golearn tracks")
	fmt.Fprintln(w, "  golearn next")
	fmt.Fprintln(w, "  golearn progress")
	fmt.Fprintln(w, "  golearn diagnose <lesson>")
	fmt.Fprintln(w, "  golearn race <lesson>")
	fmt.Fprintln(w, "  golearn bench <lesson>")
	fmt.Fprintln(w, "  golearn fuzz <lesson>")
	fmt.Fprintln(w, "  golearn profile <lesson>")
	fmt.Fprintln(w, "  golearn review <capstone>")
	fmt.Fprintln(w, "  golearn skip <item>")
	fmt.Fprintln(w, "  golearn unskip <item>")
	fmt.Fprintln(w, "  golearn reset <lesson> [--yes] [--list-backups]")
}

// ===================== TRACKS =====================

func runTracks(tracks []track, capstones []lesson, stdout io.Writer) int {
	difficultyGroups := map[string][]string{
		"beginner":     {},
		"intermediate": {},
		"advanced":     {},
		"expert":       {},
	}

	for _, t := range tracks {
		d := strings.ToLower(t.Difficulty)
		if d == "" {
			d = "beginner"
		}
		if _, ok := difficultyGroups[d]; !ok {
			d = "advanced"
		}
		difficultyGroups[d] = append(difficultyGroups[d], t.ID)
	}

	fmt.Fprintln(stdout, "Tracks by Difficulty:")
	fmt.Fprintln(stdout, "")

	for _, diff := range []string{"beginner", "intermediate", "advanced", "expert"} {
		if len(difficultyGroups[diff]) > 0 {
			fmt.Fprintf(stdout, "%s\n", strings.ToUpper(diff[:1])+diff[1:])
			for _, id := range difficultyGroups[diff] {
				fmt.Fprintf(stdout, "  %s\n", id)
			}
			fmt.Fprintln(stdout, "")
		}
	}

	if len(capstones) > 0 {
		fmt.Fprintln(stdout, "Capstones")
		fmt.Fprintln(stdout, "  (run 'golearn review <capstone>' for details)")
	}

	return 0
}

// ===================== NEXT =====================

func runNext(root string, lessons []lesson, capstones []lesson, progress *progressState, stdout io.Writer, stderr io.Writer) int {
	next, blocked, err := findNextItem(root, lessons, capstones, progress)
	if err != nil {
		fmt.Fprintf(stderr, "find next: %v\n", err)
		return 1
	}

	if next.Path == "" {
		fmt.Fprintln(stdout, "No further lessons required.")
		return 0
	}

	if blocked != "" {
		fmt.Fprintf(stdout, "Next: %s (blocked by: %s)\n", next.Path, blocked)
	} else {
		fmt.Fprintf(stdout, "Next: %s\n", next.Path)
	}
	return 0
}

func findNextItem(root string, lessons []lesson, capstones []lesson, progress *progressState) (lesson, string, error) {
	// Simple linear scan for MVP+ behavior
	// In full implementation, this would use prerequisite DAG
	for _, l := range lessons {
		if l.Optional {
			continue
		}
		if progress != nil && progress.SkippedOptional[l.Path] {
			continue
		}
		// Check if already completed by running tests
		summary, err := runGoTestJSONForPath(root, normalizeLessonPath(l.Path))
		if err != nil {
			continue // Not completed, could be next
		}
		if !summary.Passed {
			return l, "", nil
		}
	}

	// Check capstones
	for _, c := range capstones {
		if c.Optional {
			continue
		}
		if progress != nil && progress.SkippedOptional[c.Path] {
			continue
		}
		summary, err := runGoTestJSONForPath(root, filepath.Join("capstones", c.Path))
		if err != nil {
			continue
		}
		if !summary.Passed {
			return c, "", nil
		}
	}

	return lesson{Path: "", Title: "All complete"}, "", nil
}

// ===================== PROGRESS =====================

func runProgress(root string, lessons []lesson, capstones []lesson, progress *progressState, stdout io.Writer, stderr io.Writer) int {
	counts := map[string]int{"beginner": 0, "intermediate": 0, "advanced": 0, "expert": 0, "capstones": 0}
	completed := map[string]int{"beginner": 0, "intermediate": 0, "advanced": 0, "expert": 0, "capstones": 0}
	skippedOptional := 0

	for _, l := range lessons {
		d := strings.ToLower(l.Difficulty)
		if d == "" {
			d = "beginner"
		}
		if _, ok := counts[d]; !ok {
			d = "advanced"
		}
		counts[d]++

		if progress != nil && progress.SkippedOptional[l.Path] {
			skippedOptional++
			continue
		}

		summary, _ := runGoTestJSONForPath(root, normalizeLessonPath(l.Path))
		if summary.Passed {
			completed[d]++
		}
	}

	for _, c := range capstones {
		counts["capstones"]++
		if progress != nil && progress.SkippedOptional[c.Path] {
			skippedOptional++
			continue
		}
		summary, _ := runGoTestJSONForPath(root, filepath.Join("capstones", c.Path))
		if summary.Passed {
			completed["capstones"]++
		}
	}

	fmt.Fprintln(stdout, "Progress")
	for _, diff := range []string{"beginner", "intermediate", "advanced", "expert"} {
		if counts[diff] > 0 {
			fmt.Fprintf(stdout, "%s: %d/%d complete\n", diff, completed[diff], counts[diff])
		}
	}
	if counts["capstones"] > 0 {
		fmt.Fprintf(stdout, "Capstones: %d/%d required complete\n", completed["capstones"], counts["capstones"])
	}
	if skippedOptional > 0 {
		fmt.Fprintf(stdout, "Skipped optional: %d\n", skippedOptional)
	}

	return 0
}

// ===================== DIAGNOSE =====================

func runDiagnose(root string, args []string, lessons []lesson, stdout io.Writer, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "diagnose requires exactly one lesson path")
		return 2
	}

	lessonPath := normalizeLessonPath(args[0])

	// Run tests and capture diagnostic info
	cmd := exec.Command("go", "test", "-v", "-json", "./"+lessonPath)
	cmd.Dir = root
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	cmd.Run()

	summary := parseGoTestJSON(output.Bytes())

	fmt.Fprintf(stdout, "Diagnosis for %s:\n", lessonPath)

	if summary.CompileError {
		fmt.Fprintln(stdout, "  Status: COMPILE ERROR")
		fmt.Fprintln(stdout, "  Output:")
		fmt.Fprintln(stdout, summary.DiagnosticOut)
		return 1
	}

	if summary.Passed {
		fmt.Fprintln(stdout, "  Status: PASS")
		return 0
	}

	fmt.Fprintln(stdout, "  Status: FAIL")
	fmt.Fprintf(stdout, "  Failed tests: %d\n", summary.FailureCount)
	if summary.FirstFailing != "" {
		fmt.Fprintf(stdout, "  First failure: %s\n", summary.FirstFailing)
	}
	if summary.DiagnosticOut != "" {
		fmt.Fprintln(stdout, "  Output:")
		fmt.Fprintln(stdout, summary.DiagnosticOut)
	}

	return 1
}

// ===================== RACE =====================

func runRace(root string, args []string, lessons []lesson, stdout io.Writer, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "race requires exactly one lesson path")
		return 2
	}

	lessonPath := normalizeLessonPath(args[0])

	// Check if lesson supports race mode
	supported := false
	for _, l := range lessons {
		if l.Path == lessonPath {
			for _, m := range l.RunModes {
				if m == "race" {
					supported = true
				}
			}
		}
	}

	if !supported {
		fmt.Fprintf(stderr, "lesson %s does not support race mode\n", lessonPath)
		return 1
	}

	// Run with race detector
	cmd := exec.Command("go", "test", "-race", "-v", "./"+lessonPath)
	cmd.Dir = root
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	err := cmd.Run()

	if err == nil {
		fmt.Fprintf(stdout, "PASS race check: %s\n", lessonPath)
		fmt.Fprintln(stdout, "No data races were detected in the exercised code paths.")
		return 0
	}

	// Check for race report
	if bytes.Contains(output.Bytes(), []byte("DATA RACE")) {
		fmt.Fprintf(stdout, "FAIL race check: %s\n", lessonPath)
		fmt.Fprintln(stdout, "A data race was detected. Review the conflicting access report above.")
		return 1
	}

	fmt.Fprintf(stdout, "FAIL race check: %s\n", lessonPath)
	fmt.Fprintln(stdout, output.String())
	return 1
}

// ===================== BENCH =====================

func runBench(root string, args []string, lessons []lesson, stdout io.Writer, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "bench requires exactly one lesson path")
		return 2
	}

	lessonPath := normalizeLessonPath(args[0])

	// Check if lesson supports bench mode
	supported := false
	for _, l := range lessons {
		if l.Path == lessonPath {
			for _, m := range l.RunModes {
				if m == "bench" {
					supported = true
				}
			}
		}
	}

	if !supported {
		fmt.Fprintf(stderr, "lesson %s does not support bench mode\n", lessonPath)
		return 1
	}

	// Run benchmarks with short duration
	cmd := exec.Command("go", "test", "-bench=.", "-benchtime=1s", "-json", "./"+lessonPath)
	cmd.Dir = root
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	cmd.Run()

	fmt.Fprintf(stdout, "Benchmark results for %s:\n", lessonPath)

	// Parse benchmark output
	scanner := bufio.NewScanner(bytes.NewReader(output.Bytes()))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Benchmark") && !strings.Contains(line, "PASS") {
			fmt.Fprintln(stdout, line)
		}
	}

	fmt.Fprintln(stdout, "Note: Results may vary by machine. This is a smoke check only.")
	return 0
}

// ===================== FUZZ =====================

func runFuzz(root string, args []string, lessons []lesson, stdout io.Writer, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "fuzz requires exactly one lesson path")
		return 2
	}

	lessonPath := normalizeLessonPath(args[0])

	// Check if lesson supports fuzz mode
	supported := false
	for _, l := range lessons {
		if l.Path == lessonPath {
			for _, m := range l.RunModes {
				if m == "fuzz" {
					supported = true
				}
			}
		}
	}

	if !supported {
		fmt.Fprintf(stderr, "lesson %s does not support fuzz mode\n", lessonPath)
		return 1
	}

	// Run fuzzing with bounded time
	cmd := exec.Command("go", "test", "-fuzz=.", "-fuzztime=10s", "-v", "./"+lessonPath)
	cmd.Dir = root
	var output bytes.Buffer
	cmd.Stdout = &output
	cmd.Stderr = &output
	err := cmd.Run()

	fmt.Fprintf(stdout, "Fuzz results for %s:\n", lessonPath)

	if err == nil {
		fmt.Fprintln(stdout, "Fuzzing completed without failures.")
	} else {
		fmt.Fprintln(stdout, "Fuzzing completed with failures.")
	}

	return 0
}

// ===================== PROFILE =====================

func runProfile(root string, args []string, lessons []lesson, stdout io.Writer, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "profile requires exactly one lesson path")
		return 2
	}

	lessonPath := normalizeLessonPath(args[0])

	// Check if lesson supports profile mode
	supported := false
	for _, l := range lessons {
		if l.Path == lessonPath {
			for _, m := range l.RunModes {
				if m == "profile" || m == "bench" {
					supported = true
				}
			}
		}
	}

	if !supported {
		fmt.Fprintf(stderr, "lesson %s does not support profile mode\n", lessonPath)
		return 1
	}

	// Create profile output directory
	profileDir := filepath.Join(root, ".golearn", "profiles")
	if err := os.MkdirAll(profileDir, 0755); err != nil {
		fmt.Fprintf(stderr, "create profile dir: %v\n", err)
		return 1
	}

	cpuProfile := filepath.Join(profileDir, lessonPath+".cpu.prof")
	memProfile := filepath.Join(profileDir, lessonPath+".mem.prof")

	// Run with CPU profiling
	cmd := exec.Command("go", "test", "-cpuprofile", cpuProfile, "-memprofile", memProfile, "-bench=.", "-benchtime=1s", "./"+lessonPath)
	cmd.Dir = root
	cmd.Run()

	fmt.Fprintf(stdout, "Profile artifacts written to:\n")
	fmt.Fprintf(stdout, "  CPU: %s\n", cpuProfile)
	fmt.Fprintf(stdout, "  Memory: %s\n", memProfile)
	fmt.Fprintln(stdout, "Note: Profiles are not committed to version control.")

	return 0
}

// ===================== REVIEW =====================

func runReview(root string, args []string, capstones []lesson, stdout io.Writer, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "review requires exactly one capstone path")
		return 2
	}

	capstonePath := args[0]
	capstoneDir := filepath.Join(root, "capstones", capstonePath)

	// Find the capstone
	var found capstone
	for _, c := range capstones {
		if c.Path == capstonePath {
			found = capstone{
				Path:     c.Path,
				Title:    c.Title,
				RunModes: c.RunModes,
			}
			break
		}
	}

	if found.Path == "" {
		fmt.Fprintf(stderr, "capstone %s not found\n", capstonePath)
		return 1
	}

	fmt.Fprintf(stdout, "Capstone: %s\n", found.Title)
	fmt.Fprintf(stdout, "Path: %s\n", capstonePath)
	fmt.Fprintln(stdout, "")

	// Try to load rubric
	rubricPath := filepath.Join(capstoneDir, "rubric.yaml")
	if data, err := os.ReadFile(rubricPath); err == nil {
		fmt.Fprintln(stdout, "Review Rubric:")
		fmt.Fprintln(stdout, string(data))
		fmt.Fprintln(stdout, "")
	}

	// Run automated tests to show status
	testDir := filepath.Join(capstoneDir, "tests")
	if _, err := os.Stat(testDir); err == nil {
		testPackage := filepath.ToSlash(filepath.Join("capstones", capstonePath, "tests"))
		cmd := exec.Command("go", "test", "-v", "./"+testPackage)
		cmd.Dir = root
		cmd.Stdout = io.Discard
		cmd.Stderr = io.Discard
		err := cmd.Run()

		if err == nil {
			fmt.Fprintln(stdout, "Automated tests: PASS (submittable)")
		} else {
			fmt.Fprintln(stdout, "Automated tests: FAIL (not yet submittable)")
		}
	}

	return 0
}

// ===================== SKIP =====================

func runSkip(root string, args []string, lessons []lesson, capstones []lesson, progress *progressState, stdout io.Writer, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "skip requires exactly one item path")
		return 2
	}

	itemPath := args[0]

	// Check if item exists and is optional
	isOptional := false
	for _, l := range lessons {
		if l.Path == itemPath && l.Optional {
			isOptional = true
		}
	}
	for _, c := range capstones {
		if c.Path == itemPath && c.Optional {
			isOptional = true
		}
	}

	if !isOptional {
		fmt.Fprintf(stderr, "item %s is not optional and cannot be skipped\n", itemPath)
		return 1
	}

	// Load or create progress
	if progress == nil {
		progress = &progressState{SkippedOptional: map[string]bool{}}
	}
	if progress.SkippedOptional == nil {
		progress.SkippedOptional = map[string]bool{}
	}

	progress.SkippedOptional[itemPath] = true

	if err := saveProgress(root, progress); err != nil {
		fmt.Fprintf(stderr, "save progress: %v\n", err)
		return 1
	}

	fmt.Fprintf(stdout, "Skipped optional: %s\n", itemPath)
	return 0
}

// ===================== UNSKIP =====================

func runUnskip(root string, args []string, lessons []lesson, capstones []lesson, progress *progressState, stdout io.Writer, stderr io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintln(stderr, "unskip requires exactly one item path")
		return 2
	}

	itemPath := args[0]

	if progress == nil || progress.SkippedOptional == nil {
		fmt.Fprintf(stdout, "Item %s was not skipped\n", itemPath)
		return 0
	}

	if !progress.SkippedOptional[itemPath] {
		fmt.Fprintf(stdout, "Item %s was not skipped\n", itemPath)
		return 0
	}

	delete(progress.SkippedOptional, itemPath)

	if err := saveProgress(root, progress); err != nil {
		fmt.Fprintf(stderr, "save progress: %v\n", err)
		return 1
	}

	fmt.Fprintf(stdout, "Unskipped: %s\n", itemPath)
	return 0
}

// ===================== RESET =====================

func runReset(root string, args []string, lessons []lesson, stdout io.Writer, stderr io.Writer) int {
	var lessonPath string
	var listBackups bool
	var yesFlag bool

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "--list-backups":
			listBackups = true
		case "--yes", "-y":
			yesFlag = true
		case "--":
			// Skip
		default:
			if !strings.HasPrefix(arg, "-") {
				lessonPath = arg
			}
		}
	}

	if lessonPath == "" {
		fmt.Fprintln(stderr, "reset requires a lesson path")
		return 2
	}

	normalizedPath := normalizeLessonPath(lessonPath)
	lessonDir := filepath.Join(root, "lessons", normalizedPath)
	solutionDir := filepath.Join(root, "solutions", normalizedPath)

	// List backups if requested
	if listBackups {
		backupsDir := filepath.Join(root, ".golearn", "resets", normalizedPath)
		if _, err := os.Stat(backupsDir); os.IsNotExist(err) {
			fmt.Fprintf(stdout, "No backups found for %s\n", lessonPath)
			return 0
		}

		fmt.Fprintf(stdout, "Backups for %s:\n", lessonPath)
		entries, _ := os.ReadDir(backupsDir)
		for _, e := range entries {
			fmt.Fprintf(stdout, "  %s\n", e.Name())
		}
		fmt.Fprintln(stdout, "To restore, copy files from the backup directory manually.")
		return 0
	}

	// Check if lesson directory exists
	if _, err := os.Stat(lessonDir); os.IsNotExist(err) {
		fmt.Fprintf(stderr, "lesson not found: %s\n", lessonPath)
		return 1
	}

	// Check if solution exists
	if _, err := os.Stat(solutionDir); os.IsNotExist(err) {
		fmt.Fprintf(stderr, "no solution found for: %s\n", lessonPath)
		return 1
	}

	// Print what will be overwritten
	fmt.Fprintf(stdout, "Files to be overwritten:\n")
	entries, _ := os.ReadDir(lessonDir)
	for _, e := range entries {
		if !e.IsDir() {
			fmt.Fprintf(stdout, "  %s\n", e.Name())
		}
	}
	fmt.Fprintln(stdout, "")

	// Confirm unless --yes
	if !yesFlag {
		fmt.Fprint(stdout, "Proceed with reset? (y/N) ")
		var response string
		fmt.Scanln(&response)
		if response != "y" && response != "Y" {
			fmt.Fprintln(stdout, "Aborted.")
			return 0
		}
	}

	// Create backup
	backupsDir := filepath.Join(root, ".golearn", "resets", normalizedPath)
	timestamp := time.Now().UTC().Format(time.RFC3339Nano)
	backupDir := filepath.Join(backupsDir, timestamp)
	if err := os.MkdirAll(backupDir, 0755); err != nil {
		fmt.Fprintf(stderr, "create backup dir: %v\n", err)
		return 1
	}

	// Copy current files to backup
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		src := filepath.Join(lessonDir, e.Name())
		dst := filepath.Join(backupDir, e.Name())
		data, err := os.ReadFile(src)
		if err != nil {
			continue
		}
		os.WriteFile(dst, data, 0644)
	}

	// Copy solution to lesson
	solutionEntries, _ := os.ReadDir(solutionDir)
	for _, e := range solutionEntries {
		if e.IsDir() {
			continue
		}
		src := filepath.Join(solutionDir, e.Name())
		dst := filepath.Join(lessonDir, e.Name())
		data, err := os.ReadFile(src)
		if err != nil {
			continue
		}
		os.WriteFile(dst, data, 0644)
	}

	fmt.Fprintf(stdout, "Reset complete. Backup saved to: .golearn/resets/%s/%s\n", normalizedPath, timestamp)
	return 0
}

// ===================== MVP COMMANDS (unchanged) =====================

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
	summary, err := runGoTestJSONForPath(root, lessonPath)
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

// ===================== HELPER FUNCTIONS =====================

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
	lessons, _, err := loadCurriculumFile(path)
	return lessons, err
}

func loadAllCurriculum(root string) ([]lesson, []track, error) {
	return loadCurriculumFile(filepath.Join(root, "curriculum.yaml"))
}

func loadCurriculumFile(curriculumPath string) ([]lesson, []track, error) {
	file, err := os.Open(curriculumPath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var lessons []lesson
	tracksMap := map[string]*track{}

	scanner := bufio.NewScanner(file)
	var currentTrack string
	pendingLesson := -1
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "- id:") {
			pendingLesson = -1
			currentTrack = strings.TrimSpace(strings.TrimPrefix(line, "- id:"))
			if _, ok := tracksMap[currentTrack]; !ok {
				tracksMap[currentTrack] = &track{ID: currentTrack, Difficulty: "intermediate"}
			}
			continue
		}
		if strings.HasPrefix(line, "- path:") {
			value := strings.TrimSpace(strings.TrimPrefix(line, "- path:"))
			value = filepath.ToSlash(filepath.Clean(value))
			lessons = append(lessons, lesson{Path: value})
			pendingLesson = len(lessons) - 1
			if currentTrack != "" {
				tracksMap[currentTrack].Lessons = append(tracksMap[currentTrack].Lessons, lessons[len(lessons)-1])
			}
			continue
		}
		if pendingLesson >= 0 {
			if strings.HasPrefix(line, "title:") {
				lessons[pendingLesson].Title = strings.Trim(strings.TrimSpace(strings.TrimPrefix(line, "title:")), `"`)
				if currentTrack != "" && tracksMap[currentTrack] != nil && len(tracksMap[currentTrack].Lessons) > 0 {
					tracksMap[currentTrack].Lessons[len(tracksMap[currentTrack].Lessons)-1].Title = lessons[pendingLesson].Title
				}
			}
			if strings.HasPrefix(line, "difficulty:") {
				lessons[pendingLesson].Difficulty = strings.Trim(strings.TrimSpace(strings.TrimPrefix(line, "difficulty:")), `"`)
				if currentTrack != "" && tracksMap[currentTrack] != nil {
					tracksMap[currentTrack].Difficulty = lessons[pendingLesson].Difficulty
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	if len(lessons) == 0 {
		return nil, nil, errors.New("no lessons listed")
	}

	var tracks []track
	for _, t := range tracksMap {
		tracks = append(tracks, *t)
	}

	return lessons, tracks, nil
}

func loadCapstones(root string) ([]lesson, error) {
	capstonesDir := filepath.Join(root, "capstones")
	if _, err := os.Stat(capstonesDir); os.IsNotExist(err) {
		return []lesson{}, nil
	}

	var capstones []lesson
	entries, err := os.ReadDir(capstonesDir)
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		capstonePath := filepath.Join(capstonesDir, e.Name(), "capstone.yaml")
		data, err := os.ReadFile(capstonePath)
		if err != nil {
			continue
		}

		c := parseCapstoneYAML(data)

		capstones = append(capstones, lesson{
			Path:       e.Name(),
			Title:      c.Title,
			Difficulty: c.Difficulty,
			RunModes:   c.RunModes,
			Optional:   c.Optional,
		})
	}

	return capstones, nil
}

func parseCapstoneYAML(data []byte) capstone {
	var c capstone
	var listField string
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		switch {
		case strings.HasPrefix(line, "title:"):
			c.Title = trimYAMLScalar(strings.TrimPrefix(line, "title:"))
			listField = ""
		case strings.HasPrefix(line, "difficulty:"):
			c.Difficulty = trimYAMLScalar(strings.TrimPrefix(line, "difficulty:"))
			listField = ""
		case strings.HasPrefix(line, "estimated_scope:"):
			c.EstimatedScope = trimYAMLScalar(strings.TrimPrefix(line, "estimated_scope:"))
			listField = ""
		case strings.HasPrefix(line, "optional:"):
			c.Optional = strings.EqualFold(trimYAMLScalar(strings.TrimPrefix(line, "optional:")), "true")
			listField = ""
		case strings.HasPrefix(line, "run_modes:"):
			listField = "run_modes"
		case strings.HasPrefix(line, "- "):
			if listField == "run_modes" {
				c.RunModes = append(c.RunModes, trimYAMLScalar(strings.TrimPrefix(line, "- ")))
			}
		default:
			listField = ""
		}
	}
	return c
}

func trimYAMLScalar(value string) string {
	return strings.Trim(strings.TrimSpace(value), `"`)
}

func loadProgress(root string) (*progressState, error) {
	progressPath := filepath.Join(root, ".golearn", "progress.json")
	data, err := os.ReadFile(progressPath)
	if err != nil {
		return nil, err
	}

	var progress progressState
	if err := json.Unmarshal(data, &progress); err != nil {
		return nil, err
	}

	return &progress, nil
}

func saveProgress(root string, progress *progressState) error {
	progressDir := filepath.Join(root, ".golearn")
	if err := os.MkdirAll(progressDir, 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(progress, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filepath.Join(progressDir, "progress.json"), data, 0644)
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

func runGoTestJSONForPath(root string, lessonPath string) (testSummary, error) {
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

// Extended parser for build-error JSON and legacy text output
func parseGoTestJSON(data []byte) testSummary {
	var summary testSummary
	var diagnostics strings.Builder
	failedTests := map[string]bool{}

	// First, try to parse as JSON lines
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		var event testEvent
		if err := json.Unmarshal(scanner.Bytes(), &event); err != nil {
			// Not valid JSON - might be legacy text output
			continue
		}
		if event.Output != "" {
			diagnostics.WriteString(event.Output)
		}
		if event.Action == "build-fail" || event.Action == "fail" && event.Test == "" {
			summary.CompileError = true
			summary.BuildError = true
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
		// Check for race detector output
		if strings.Contains(event.Output, "DATA RACE") {
			summary.RaceError = true
		}
	}

	// If no JSON events found, try legacy text parsing
	if summary.FailureCount == 0 && !summary.CompileError {
		legacySummary := parseLegacyTestOutput(data)
		if legacySummary.CompileError {
			summary.CompileError = true
			summary.DiagnosticOut = legacySummary.DiagnosticOut
		}
		if legacySummary.FailureCount > 0 {
			summary.FirstFailing = legacySummary.FirstFailing
			summary.FailureCount = legacySummary.FailureCount
			summary.DiagnosticOut = legacySummary.DiagnosticOut
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

func parseLegacyTestOutput(data []byte) testSummary {
	var summary testSummary
	var diagnostics strings.Builder
	failedTests := map[string]bool{}

	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		diagnostics.WriteString(line)
		diagnostics.WriteString("\n")

		// Check for compile errors
		if strings.Contains(line, "# ") || strings.Contains(line, "cannot find") {
			summary.CompileError = true
		}

		// Check for test failures
		if strings.HasPrefix(line, "---") && strings.HasSuffix(line, "FAIL") {
			name := strings.TrimSpace(strings.TrimPrefix(strings.TrimSuffix(line, "FAIL"), "---"))
			if !failedTests[name] {
				failedTests[name] = true
				summary.FailureCount++
				if summary.FirstFailing == "" {
					summary.FirstFailing = name
				}
			}
		}
	}

	summary.DiagnosticOut = diagnostics.String()
	return summary
}
