package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Person represents a person with ID and Name
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