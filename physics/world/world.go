package physics

import "github.com/therohans/mesh/algebra"

type WorldOptions struct {
	Gravity    algebra.Vector
	AllowSleep bool
}

// World the physics world
type World struct {
}
