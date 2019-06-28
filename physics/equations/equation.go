package equations

import (
	"github.com/therohans/mesh/physics/maths"
	"github.com/therohans/mesh/physics/objects"
)

type Equator interface {
}

// Equation base class
type Equation struct {
	ID       uint64
	MinForce float64 // -1e6
	MaxForce float64 // -1e6
	Bi       objects.Body
	Bj       objects.Body
	// Spook parameter
	A float64 // 0.0
	// Spook parameter
	B float64 // 0.0
	// Spook parameter
	Eps              float64 // 0.0
	JacobianElementA maths.JacobianElement
	JacobianElementB maths.JacobianElement
	Enabled          bool
	// A number, proportional to the force added to the bodies.
	Multiplier float64 // 0.0
}
