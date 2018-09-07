package physics

import (
	"github.com/therohans/mesh/algebra"
)

type ConvexPolyhedron struct {
	Shape
	points      []algebra.Vector
	faces       [][]uint32
	uniqueAxes  []algebra.Vector
	faceNormals []algebra.Vector
}

func NewConvexPolyhedron(points []algebra.Vector, faces [][]uint32, uniqueAxes []algebra.Vector, faceNormals []algebra.Vector) *ConvexPolyhedron {
	cp := ConvexPolyhedron{
		Shape: Shape{
			Type: ShapeConvexPoly,
		},
		points:      points,
		faces:       faces,
		uniqueAxes:  uniqueAxes,
		faceNormals: faceNormals,
	}
	return &cp
}
