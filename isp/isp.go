package isp

// this package demonstrates the Interface Segregation Principle (ISP)

import (
	"fmt"
)

// bad example: Fat interface
/*
type Employee interface {
	Code()
	ProductManagement()
	AgileManagement()
}

type Developer struct {
	Name string
}

func (d *Developer) Code() {
	fmt.Printf("%s is writing code.\n", d.Name)
}

func (d *Developer) ProductManagement() {
	// Not applicable for Developer
}

func (d *Developer) AgileManagement() {
	// Not applicable for Developer
}

type Manager struct {
	Name string
}

func (m *Manager) Code() {
	// Not applicable for Manager
}

func (m *Manager) ProductManagement() {
	fmt.Printf("%s is managing the product.\n", m.Name)
}

func (m *Manager) AgileManagement() {
	fmt.Printf("%s is managing agile processes.\n", m.Name)
}
*/

// good example: Segregated interfaces
type Coder interface {
	Code()
}

type ProductManager interface {
	ProductManagement()
}

type AgileManager interface {
	AgileManagement()
}

type Developer struct {
	Name string
}

func (d *Developer) Code() {
	fmt.Printf("%s is writing code.\n", d.Name)
}

type Manager struct {
	Name string
}

func (m *Manager) ProductManagement() {
	fmt.Printf("%s is managing the product.\n", m.Name)
}

func (m *Manager) AgileManagement() {
	fmt.Printf("%s is managing agile processes.\n", m.Name)
}

// Example usage
func Example() {
	dev := &Developer{Name: "Alice"}
	dev.Code()

	mgr := &Manager{Name: "Bob"}
	mgr.ProductManagement()
	mgr.AgileManagement()
}
