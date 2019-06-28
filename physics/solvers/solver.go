package solvers

import (
	"github.com/therohans/mesh/physics/equations"
	"github.com/therohans/mesh/physics/world"
)

type Solver interface {
	Solve(dt float64, world world.World)
	AddEquation(e equations.Equation)
	RemoveEquation(e equations.Equation)
	RemoveAll()
}
