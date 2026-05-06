# Go Learning Program

Go Learning Program is a small, CLI-assisted curriculum for practicing Go with focused lessons. Lessons live under `lessons/<track>/<lesson>/`, and reference implementations live under the matching `solutions/<track>/<lesson>/lesson.go` path.

The expert extension adds advanced tracks, specialized run modes, local progress state, and capstone self-review. Beginner lessons remain usable with plain `go test` and do not require expert tooling.

## Quick Start

```bash
go run ./cmd/go_learning_program list
go run ./cmd/go_learning_program hint 01-basics/01-hello-world
go run ./cmd/go_learning_program test 01-basics/01-hello-world
```

Expert commands:

```bash
go run ./cmd/go_learning_program tracks
go run ./cmd/go_learning_program next
go run ./cmd/go_learning_program progress
go run ./cmd/go_learning_program diagnose 01-basics/01-hello-world
go run ./cmd/go_learning_program race 11-advanced-concurrency/05-race-debugging
go run ./cmd/go_learning_program bench 13-performance-memory/01-benchmarking
go run ./cmd/go_learning_program fuzz 12-quality-engineering/04-fuzzing
go run ./cmd/go_learning_program profile 13-performance-memory/01-benchmarking
go run ./cmd/go_learning_program review 03-production-http-service
go run ./cmd/go_learning_program skip <optional-item>
go run ./cmd/go_learning_program unskip <optional-item>
go run ./cmd/go_learning_program reset <lesson> --yes
go run ./cmd/go_learning_program reset --list-backups <lesson>
```

The root module uses Go 1.22. Normal repository checks are:

```bash
go test ./...
go vet ./...
gofmt -w .
```

Generated CLI artifacts live under `.golearn/` and are ignored by git. Committed fixtures and fuzz corpora belong in lesson-local `testdata/` or capstone `fixtures/`.

## Lesson Layout

Each lesson directory contains:

- `lesson.go`: learner-facing code.
- `lesson_test.go`: pedagogical tests for the lesson behavior.
- `lesson.yaml`: lesson metadata.
- `hints.yaml`: ordered hints.

Each solution directory contains exactly one Go file:

- `lesson.go`: a single-file reference solution starting with `//go:build ignore`.

Package names are lowercase identifiers. Folder names use numbered kebab-case so curriculum ordering stays stable.

## Capstones

Capstones live under `capstones/<capstone-id>/` and contain learner-facing requirements, metadata, starter code, automated acceptance tests, fixtures, a rubric, and maintainer notes. Normal learner commands print the README, test status, and rubric; maintainer notes are for authors and reviewers.
