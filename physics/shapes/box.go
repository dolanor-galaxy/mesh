package physics

import (
	"github.com/therohans/mesh/algebra"
)

// Box A 3d box shape.
type Box struct {
	Shape
	HalfExtents algebra.Vector
}

func NewBox(halfExtents algebra.Vector) Box {
	box := Box{}
	box.Type = ShapeBox
	box.HalfExtents = halfExtents
	return box
}

func (b *Box) CalculateLocalInertia(mass float64, target *algebra.Vector) {
	b.calculateInertia(b.HalfExtents, mass, target)
}

func (b *Box) calculateInertia(halfExtents algebra.Vector, mass float64, target *algebra.Vector) {
	e := halfExtents
	target.X = 1.0 / 12.0 * mass * (2*e.Y*2*e.Y + 2*e.Z*2*e.Z)
	target.Y = 1.0 / 12.0 * mass * (2*e.X*2*e.X + 2*e.Z*2*e.Z)
	target.Z = 1.0 / 12.0 * mass * (2*e.Y*2*e.Y + 2*e.X*2*e.X)
}

// GetSideNormals Get the box 6 side normals
func (b *Box) GetSideNormals(sixTargetVectors *[6]algebra.Vector, quat algebra.Quaternion) {
	sides := sixTargetVectors
	ex := b.HalfExtents

	sides[0].Set(ex.X, 0, 0, 0)
	sides[1].Set(0, ex.Y, 0, 0)
	sides[2].Set(0, 0, ex.Z, 0)
	sides[3].Set(-ex.X, 0, 0, 0)
	sides[4].Set(0, -ex.Y, 0, 0)
	sides[5].Set(0, 0, -ex.Z, 0)

	for i := 0; i != 5; i++ {
		tmp := algebra.Vector{}
		// quat.vmult(sides[i], sides[i])
		quat.MulV(sides[i], &tmp)
		sides[i] = tmp
	}
}

// UpdateConvexPolyhedronRepresentation
func (b *Box) UpdateConvexPolyhedronRepresentation() {

}
