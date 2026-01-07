package main

import (
	"github.com/meyuviofficial/go_solid_principles/dip"
	"github.com/meyuviofficial/go_solid_principles/isp"
	"github.com/meyuviofficial/go_solid_principles/lsp"
	"github.com/meyuviofficial/go_solid_principles/ocp" // open/closed principle
	"github.com/meyuviofficial/go_solid_principles/srp" // single responsibility principle
)

func main() {
	// explains single responsibility principle
	srp.Example()

	// explains open/closed principle
	// ocp.SimpleExample() // uncomment to see simple example
	ocp.AdvancedExample()

	// explains liskov substitution principle
	lsp.Example()

	// explains interface segregation principle
	isp.Example()

	// explain dependency inversion principle
	dip.Example()
}
