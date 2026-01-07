package srp

// this package demonstrates the Single Responsibility Principle (SRP)

import (
	"fmt"
)

type Developer struct {
	Name string
}

func (d *Developer) WriteCode() {
	fmt.Printf("%s is writing code.\n", d.Name)
}

type Manager struct {
	Name string
}

func (m *Manager) ManageProject() {
	fmt.Printf("%s is managing the project.\n", m.Name)
}

// Example usage
func Example() {
	dev := Developer{Name: "Alice"}
	dev.WriteCode()

	mgr := Manager{Name: "Bob"}
	mgr.ManageProject()
}
