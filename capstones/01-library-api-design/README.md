# Library API Design

Design a small Go library API that remains stable while the implementation evolves.

## Requirements

- Define a public package surface with clear exported names and documentation.
- Keep implementation details behind internal helpers or unexported types.
- Return errors that callers can inspect with behavior checks or `errors.Is`.
- Include tests that exercise compatibility and failure behavior.
- Explain one design tradeoff between a minimal API and future extension points.

## Operational Expectations

- The package must be usable without network access or credentials.
- Tests should be deterministic and runnable with plain `go test`.

## Stretch Goals

- Add a compatibility example that demonstrates a caller upgrading without code changes.
