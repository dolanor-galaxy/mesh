package core

import "github.com/therohans/mesh/algebra"

// Transform position in space
type Transform struct {
	Position algebra.Vector
	Rotation algebra.Vector
	Scale    algebra.Vector
}
