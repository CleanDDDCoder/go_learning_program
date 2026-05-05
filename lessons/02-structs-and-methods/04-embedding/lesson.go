package embedding

// Contact stores a display name.
type Contact struct {
	Name string
}

// Employee embeds Contact and adds a role.
type Employee struct {
	Contact
	Role string
}

// Label returns a compact employee label.
func (employee Employee) Label() string {
	return employee.Name + " - " + employee.Role
}
