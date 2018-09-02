package physics

import "github.com/therohans/mesh/algebra"

// The available shape types.
const (
	ShapeSphere      uint = 1
	ShapePlane       uint = 2
	ShapeBox         uint = 4
	ShapeCompound    uint = 8
	ShapeConvexPoly  uint = 16
	ShapeHeightfield uint = 32
	ShapeParticle    uint = 64
	ShapeCylinder    uint = 128
	ShapeTrimesh     uint = 256
)

// ShapeOptions properites for a shape
type ShapeOptions struct {
	CollisionFilterGroup int
	CollisionFilterMask  int
	// CollisionResponse Whether to produce contact forces when in contact with other bodies. Note that contacts will be generated, but they will be disabled.
	CollisionResponse bool
	// Material             physics.Material
}

// Shape Base class for shapes
type Shape struct {
	Options ShapeOptions
	// Type the shape type (see Shape<X>)
	// Must be set to an int > 0 by subclasses.
	Type uint
	// The local bounding sphere radius of this shape.
	BoundingSphereRadius float32
}

// UpdateBoundingSphereRadius UpdateBoundingSphereRadius Computes the bounding sphere radius.
// The result is stored in the property: boundingSphereRadius
func (s *Shape) UpdateBoundingSphereRadius() {
	panic("computeBoundingSphereRadius() not implemented for shape type")
}

// Volume Get the volume of this shape
func (s *Shape) Volume() float32 {
	panic("volume() not implemented for shape type")
}

// CalculateLocalInertia Calculates the inertia in the local frame for this shape.
func (s *Shape) CalculateLocalInertia(mass float32, target algebra.Vector) {
	panic("calculateLocalInertia() not implemented for shape type ")
}
