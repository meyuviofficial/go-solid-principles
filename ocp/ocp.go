package ocp

// this package demonstrates the Open/Closed Principle (OCP)

import (
	"fmt"

	. "github.com/meyuviofficial/go_solid_principles/interfaces"
)

type Spiderman struct {
	Name string
}

func (s *Spiderman) Power() {
	fmt.Printf("%s uses agility and spider-sense!\n", s.Name)
}

type Ironman struct {
	Name string
}

func (i *Ironman) Power() {
	fmt.Printf("%s uses advanced technology and armor!\n", i.Name)
}

// Example: Simple
func SimpleExample() {
	heroes := []SuperHero{
		&Spiderman{Name: "Peter Parker"},
		&Ironman{Name: "Tony Stark"},
		// New heroes can be added here without modifying existing code
	}

	for _, hero := range heroes {
		hero.Power()
	}
}

// Advanced Example
type SuperHeroConstructor func(name string) SuperHero

var heroRegistry = make(map[string]SuperHeroConstructor)

func RegisterHero(heroType string, constructor SuperHeroConstructor) {
	heroRegistry[heroType] = constructor
}

func CreateHero(heroType, name string) SuperHero {
	if constructor, exists := heroRegistry[heroType]; exists {
		return constructor(name)
	}
	return nil
}

func AdvancedExample() {
	// Register heroes
	RegisterHero("Spiderman", func(name string) SuperHero {
		return &Spiderman{Name: name}
	})
	RegisterHero("Ironman", func(name string) SuperHero {
		return &Ironman{Name: name}
	})

	// Create heroes
	heroes := []SuperHero{
		CreateHero("Spiderman", "Peter Parker"),
		CreateHero("Ironman", "Tony Stark"),
		// New heroes can be registered and created without modifying existing code
	}
	for _, hero := range heroes {
		hero.Power()
	}
}
