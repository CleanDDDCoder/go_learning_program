# JSON Streaming and Validation

Work with streaming JSON and validate incoming data.

## Objectives

- Use `json.Decoder` for streaming JSON
- Validate JSON structure and types
- Handle partial reads gracefully

## Exercise

Create a program that:

1. Reads JSON array from stdin using streaming (not loading entire input)
2. Validates each object has required fields: `id` (number), `name` (string)
3. Prints valid objects to stdout as JSON
4. Reports errors for invalid objects without stopping

## Hints

- Use `json.NewDecoder(os.Stdin)` to stream from stdin
- Use `json.Decoder.Token()` to read array elements
- Check for `json.UnmarshalTypeError` for type mismatches

## Solution

```go
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	dec := json.NewDecoder(os.Stdin)

	// Read the opening bracket
	token, err := dec.Token()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading token: %v\n", err)
		return
	}
	_ = token

	fmt.Println("[")

	first := true
	for dec.More() {
		var person Person
		if err := dec.Decode(&person); err != nil {
			fmt.Fprintf(os.Stderr, "Error decoding: %v\n", err)
			continue
		}
		if person.ID == 0 || person.Name == "" {
			fmt.Fprintf(os.Stderr, "Missing required fields: %+v\n", person)
			continue
		}
		if !first {
			fmt.Println(",")
		}
		jsonBytes, _ := json.Marshal(person)
		fmt.Print(string(jsonBytes))
		first = false
	}

	fmt.Println("\n]")
}
```