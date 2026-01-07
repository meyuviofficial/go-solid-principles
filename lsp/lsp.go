package lsp

// this package demonstrates the Liskov Substitution Principle (LSP)

import (
	"fmt"

	. "github.com/meyuviofficial/go_solid_principles/interfaces"
)

// bad example: Violates LSP - not all heroes can fly!
/*
type FlyingHero interface {
	Fly()
	Power()
}

type Ironman struct {
	name string
}

func (i *Ironman) Fly() {
	fmt.Printf("%s flies with jet propulsion!\n", i.name)
}

func (i *Ironman) Power() {
	fmt.Printf("%s uses advanced technology!\n", i.name)
}

type Hulk struct {
	name string
}

func (h *Hulk) Fly() {
	// LSP Violation: Hulk can't fly but forced to implement this
	panic("Hulk can't fly!")
	// Or worse: empty implementation that silently does nothing
}

func (h *Hulk) Power() {
	fmt.Printf("%s smashes with incredible strength!\n", h.name)
}

// Client code expects ANY FlyingHero to fly safely
func MakeHeroFly(hero FlyingHero) {
	hero.Fly() // Works for Ironman, crashes for Hulk!
}
*/

// good example: Respects LSP - substitutable without surprises
type Thor struct {
	name string
}

func (t *Thor) Power() {
	fmt.Printf("%s wields the power of thunder!\n", t.name)
}

type CaptainAmerica struct {
	name string
}

func (c *CaptainAmerica) Power() {
	fmt.Printf("%s is worthy and can use Mjolnir!\n", c.name)
}

// Client code works with ANY SuperHero without special handling
func Call(sh SuperHero) {
	sh.Power() // Works consistently for all heroes
}

// Example usage
func Example() {
	thor := &Thor{name: "Thor"}
	captain := &CaptainAmerica{name: "Steve Rogers"}

	Call(thor)    // Valid: Thor uses his power
	Call(captain) // Valid: Captain America uses his power
	
	// Both heroes are perfectly substitutable - no crashes, no surprises!
}
