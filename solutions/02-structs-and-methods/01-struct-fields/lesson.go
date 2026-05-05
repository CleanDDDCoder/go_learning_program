//go:build ignore

package struct_fields

import "fmt"

// Person groups a name and age.
type Person struct {
	Name string
	Age  int
}

// Summary returns a short sentence describing the person.
func Summary(person Person) string {
	return fmt.Sprintf("%s is %d", person.Name, person.Age)
}
