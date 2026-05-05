package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	name := flag.String("name", "World", "Name to greet")
	greet := flag.Bool("greet", false, "Enable greeting")

	// Parse command-line arguments
	flag.Parse()

	// If greet is enabled, print the greeting
	if *greet {
		fmt.Printf("Hello, %s!\n", *name)
	}
}
