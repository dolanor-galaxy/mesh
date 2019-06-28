package maths

import (
	"github.com/therohans/mesh/algebra"
)

// JacobianElement An element containing 6 entries, 3 spatial and 3 rotational degrees of freedom.
type JacobianElement struct {
	Spatial    algebra.Vector
	Rotational algebra.Vector
}

// MultiplyElement multiply with other JacobianElement
func (j *JacobianElement) MultiplyElement(element JacobianElement) float64 {
	return element.Spatial.Dot(j.Spatial) + element.Rotational.Dot(j.Rotational)
}

// MultiplyVectors with two vectors
func (j *JacobianElement) MultiplyVectors(spatial algebra.Vector, rotational algebra.Vector) float64 {
	return spatial.Dot(j.Spatial) + rotational.Dot(j.Rotational)
}
