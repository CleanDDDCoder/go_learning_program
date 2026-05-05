# Contributing

## Repository Structure

Lessons are stored in `lessons/<track>/<lesson>/`. Reference solutions are stored in `solutions/<track>/<lesson>/`.

## Adding a Lesson

1. Add a lesson directory under the correct track.
2. Add `lesson.go`, `lesson_test.go`, `lesson.yaml`, and `hints.yaml`.
3. Add `solutions/<track>/<lesson>/lesson.go`.
4. Add the lesson path to `curriculum.yaml`.
5. Run `gofmt`, `go vet ./...`, and `go test ./...`.

## Required Metadata

`lesson.yaml` owns lesson-local metadata only:

```yaml
package: functions
title: Functions
unsolved_state: test-fail
concepts:
  - functions
difficulty: beginner
```

Do not add `id` or `track` to `lesson.yaml`; the filesystem path and `curriculum.yaml` already provide that identity.

## Authoring Rules

- Keep each lesson focused on one concept.
- Put clear task comments near the code learners should edit.
- Use named table-driven cases and `t.Run` when a lesson has multiple scenarios.
- Test success and failure paths where the lesson behavior has both.
- Do not add extra Go files to solution directories.
- Every solution must be a single `lesson.go` file.
- Every solution file must start with `//go:build ignore`, followed by a blank line and the package declaration.
- Hints should guide learners incrementally without revealing the full solution in the first hint.

## Common Mistakes

- Forgetting to list a lesson in `curriculum.yaml`.
- Adding `id` or `track` fields to `lesson.yaml`.
- Adding extra `.go` files under `solutions/`.
- Forgetting to run `gofmt`.
