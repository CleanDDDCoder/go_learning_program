# Go Learning Program

Go Learning Program is a small, CLI-assisted curriculum for practicing Go with focused lessons. Lessons live under `lessons/<track>/<lesson>/`, and reference implementations live under the matching `solutions/<track>/<lesson>/lesson.go` path.

## Quick Start

```bash
go run ./cmd/go_learning_program list
go run ./cmd/go_learning_program hint 01-basics/01-hello-world
go run ./cmd/go_learning_program test 01-basics/01-hello-world
```

The root module uses Go 1.22. Normal repository checks are:

```bash
go test ./...
go vet ./...
gofmt -w .
```

## Lesson Layout

Each lesson directory contains:

- `lesson.go`: learner-facing code.
- `lesson_test.go`: pedagogical tests for the lesson behavior.
- `lesson.yaml`: lesson metadata.
- `hints.yaml`: ordered hints.

Each solution directory contains exactly one Go file:

- `lesson.go`: a single-file reference solution starting with `//go:build ignore`.

Package names are lowercase identifiers. Folder names use numbered kebab-case so curriculum ordering stays stable.
