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

## 13. Expert lesson metadata

Expert lessons may add `lesson_type`, `min_go_version`, `run_modes`, `optional`, `prerequisites`, and `recommended_prerequisites`. Keep `difficulty` one of `beginner`, `intermediate`, `advanced`, or `expert`.

## 14. Expert lesson types

Use `unit`, `integration`, `fuzz`, `benchmark`, `profile`, `race`, or `capstone-prep` only when the lesson actually supports that mode.

## 15. Multi-file solution overlays

MVP lessons keep one ignored `lesson.go` solution. Expert lessons that need overlays must document every overwritten learner file and keep generated artifacts out of source control.

## 16. Race lesson authoring

Race lessons must include deterministic behavioral tests. The race detector is evidence for executed paths, not a proof of total race freedom.

## 17. Fuzz lesson authoring

Commit seed and regression corpus files under `testdata/fuzz/<FuzzTarget>/`. Keep normal CI bounded and keep `.golearn/fuzz/` ephemeral.

## 18. Goroutine leak lesson authoring

Leak lessons must assert shutdown through synchronization or a pinned local dependency. Do not use race mode as the leak check.

## 19. Benchmark lesson authoring

Prefer benchmark existence, correctness, and allocation structure over absolute timing thresholds. Explain measurement tradeoffs in the task comment or review prompt.

## 20. Integration lesson authoring

Integration lessons use local resources only: `httptest`, temp files, temp directories, loopback listeners, in-memory fakes, or committed fixtures.

## 21. Capstone authoring

Capstones use `README.md`, `capstone.yaml`, `starter/`, `tests/`, `fixtures/`, `rubric.yaml`, and `maintainer_notes.md`. The README owns learner-facing requirements.

## 22. Review rubric authoring

Rubrics must cover correctness, API design, error handling, concurrency safety, test quality, performance reasoning, security and reliability, observability, maintainability, and documentation.

## 23. Generated artifact policy

Generated learner state and command outputs live under `.golearn/`, including profiles, benchmarks, fuzz output, reset backups, and `progress.json`.

## 24. Dependency policy

Prefer the standard library. Third-party dependencies must be pinned, documented, credential-free, license-acceptable, and deterministic in CI.

## 25. Expert CLI expectations

Expert commands must preserve beginner behavior and leave plain `go test` usable. Commands that generate files must write only under `.golearn/`.

## 26. Progress and reset safety

Progress state is local. Reset must show an overwrite plan, require confirmation or `--yes`, create a retained backup, and avoid `.golearn/progress.json`.

## 27. Concept tag normalization

Use lowercase kebab-case concept tags. Do not rename MVP concepts without a compatibility plan.

## 28. Optional and prerequisite semantics

Only optional lessons or capstones can be skipped. Skipped items do not satisfy prerequisites for other items.

## 29. Reset backup retention

Reset backups under `.golearn/resets/` are retained until the learner deletes them manually.

## 30. Version, dependency, and MVP concept lifecycle TODO

Future lifecycle policy must define the supported Go version window, dependency review cadence, deprecation behavior, replacement compatibility, and MVP concept migration rules.
