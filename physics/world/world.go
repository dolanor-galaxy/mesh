package world

import (
	"github.com/therohans/mesh/algebra"
	"github.com/therohans/mesh/physics/objects"
)

// WorldOptions
type WorldOptions struct {
	Gravity    algebra.Vector
	AllowSleep bool
	Broadphase interface{}
	Solver     interface{}
}

// World the physics world
type World struct {
	Options           WorldOptions
	dt                float64
	time              float64
	contacts          []interface{}
	frictionEquations []string
	stepNumber        uint32
	Bodies            []objects.Body
}

func NewWorld(options WorldOptions) *World {
	world := World{
		Options: options,
	}
	return &world
}
