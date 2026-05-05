# Hello CLI

In this lesson, you'll create a simple command-line interface (CLI) program in Go.

## Learning Objectives

- Learn to use the `flag` package for CLI arguments
- Understand how to handle boolean and string flags
- Build a simple CLI tool that responds to command-line inputs

## Instructions

Create a file called `main.go` that:
1. Accepts a `--name` flag (string) that defaults to "World"
2. Accepts a `--greet` flag (boolean) that enables the greeting
3. When `--greet` is set, prints "Hello, {name}!" to stdout

## Example Usage

```bash
# Default behavior (no greeting)
go run main.go
# Output: (nothing)

# With --greet flag
go run main.go --greet
# Output: Hello, World!

# With custom name
go run main.go --name=Alice --greet
# Output: Hello, Alice!
```

## Hints

- Use `flag.String()` to define string flags
- Use `flag.Bool()` to define boolean flags
- Call `flag.Parse()` before using the flag values
- Access flag values via the pointers returned by `flag.String()` and `flag.Bool()`

## Solution

See `solution/main.go` for a reference implementation.