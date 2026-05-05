# HTTP Clients and Timeouts

Create an HTTP client with proper timeout configuration.

## Objectives

- Understand Go's `http.Client`
- Configure timeouts for HTTP requests
- Handle request and response properly

## Exercise

Create an HTTP client that:

1. Creates a client with a 5-second timeout
2. Makes a GET request to `http://httpbin.org/delay/2`
3. Handles timeout errors gracefully

## Hints

- Use `http.Client` with `Timeout` field
- Use `context.WithTimeout` for more granular control
- Check for `context.DeadlineExceeded` errors

## Solution

```go
package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://httpbin.org/delay/2", nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Request timed out")
		} else {
			fmt.Printf("Error making request: %v\n", err)
		}
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Response status: %s\n", resp.Status)
}
```