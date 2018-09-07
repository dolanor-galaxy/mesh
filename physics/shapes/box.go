package physics

import (
	"github.com/therohans/mesh/algebra"
)

// Box A 3d box shape.
type Box struct {
	Shape
	HalfExtents      algebra.Vector
	convexPolyhedron *ConvexPolyhedron
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

// UpdateConvexPolyhedron
func (b *Box) UpdateConvexPolyhedron() {
	sx := b.HalfExtents.Z
	sy := b.HalfExtents.Y
	sz := b.HalfExtents.Z
	// V := algebra.Vector{}

	vertices := []algebra.Vector{
		algebra.Vector{X: -sx, Y: -sy, Z: -sz},
		algebra.Vector{X: sx, Y: -sy, Z: -sz},
		algebra.Vector{X: sx, Y: sy, Z: -sz},
		algebra.Vector{X: -sx, Y: sy, Z: -sz},
		algebra.Vector{X: -sx, Y: -sy, Z: sz},
		algebra.Vector{X: sx, Y: -sy, Z: sz},
		algebra.Vector{X: sx, Y: sy, Z: sz},
		algebra.Vector{X: -sx, Y: sy, Z: sz},
	}

	indices := [][]uint32{
		[]uint32{3, 2, 1, 0}, // -z
		[]uint32{4, 5, 6, 7}, // +z
		[]uint32{5, 4, 0, 1}, // -y
		[]uint32{2, 3, 7, 6}, // +y
		[]uint32{0, 4, 7, 3}, // -x
		[]uint32{1, 2, 6, 5}, // +x
	}

	axes := []algebra.Vector{
		algebra.AxisZ,
		algebra.AxisY,
		algebra.AxisX,
	}

	normals := []algebra.Vector{}

	h := NewConvexPolyhedron(vertices, indices, axes, normals)
	b.convexPolyhedron = h
	// h.material = b.material
}

// Volume Caluates the volume of the box (uses half extents)
func (b *Box) Volume() float64 {
	return 8.0 * b.HalfExtents.X * b.HalfExtents.Y * b.HalfExtents.Z
}

// UpdateBoundingSphereRadius using the half extents, update the bounding sphere
func (b *Box) UpdateBoundingSphereRadius() {
	b.BoundingSphereRadius = b.HalfExtents.Norm()
}

var worldCornersTemp = [8]algebra.Vector{
	algebra.Vector{},
	algebra.Vector{},
	algebra.Vector{},
	algebra.Vector{},
	algebra.Vector{},
	algebra.Vector{},
	algebra.Vector{},
	algebra.Vector{},
}

// CalculateWorldAABB finds the min and max vectors given a world position (both are modifed by this)
func (b *Box) CalculateWorldAABB(pos *algebra.Vector, quat *algebra.Quaternion, min *algebra.Vector, max *algebra.Vector) {
	var e = b.HalfExtents
	worldCornersTemp[0].Set(e.X, e.Y, e.Z, 0)
	worldCornersTemp[1].Set(-e.X, e.Y, e.Z, 0)
	worldCornersTemp[2].Set(-e.X, -e.Y, e.Z, 0)
	worldCornersTemp[3].Set(-e.X, -e.Y, -e.Z, 0)
	worldCornersTemp[4].Set(e.X, -e.Y, -e.Z, 0)
	worldCornersTemp[5].Set(e.X, e.Y, -e.Z, 0)
	worldCornersTemp[6].Set(-e.X, e.Y, -e.Z, 0)
	worldCornersTemp[7].Set(e.X, -e.Y, e.Z, 0)

	var wc = worldCornersTemp[0]
	quat.MulV(wc, &wc)
	pos.AddV(wc, &wc)
	max.Copy(wc)
	min.Copy(wc)

	for i := 1; i < 8; i++ {
		var wc = worldCornersTemp[i]
		quat.MulV(wc, &wc)
		pos.AddV(wc, &wc)
		var x = wc.X
		var y = wc.Y
		var z = wc.Z
		if x > max.X {
			max.X = x
		}
		if y > max.Y {
			max.Y = y
		}
		if z > max.Z {
			max.Z = z
		}
		if x < min.X {
			min.X = x
		}
		if y < min.Y {
			min.Y = y
		}
		if z < min.Z {
			min.Z = z
		}
	}
}
