package core

import (
	"math"

	"github.com/robrohan/mesh/algebra"
)

// NewComponentCamera create a new default camera
func NewComponentCamera() ComponentCamera {
	return ComponentCamera{
		Component: &Component{
			Parent: &Entity{},
		},
		View:       &algebra.Matrix{},
		Projection: &algebra.Matrix{},
	}
}

// ComponentCamera component to view the world
type ComponentCamera struct {
	*Component
	View       *algebra.Matrix
	Projection *algebra.Matrix
	PixelRatio float32
}

// GetView get the current view matrix
func (c *ComponentCamera) GetView() *algebra.Matrix {
	return c.View
}

// GetProjection get the current projection matrix
func (c *ComponentCamera) GetProjection() *algebra.Matrix {
	return c.Projection
}

// UpdateViewMatrix update the view model based on the parents transform
func (c *ComponentCamera) UpdateViewMatrix() {
	transform := c.GetParent().Transform

	translation := algebra.Matrix{}
	translation.InitTranslation(&transform.Position)

	rotation := algebra.Matrix{}
	rotation.InitFUR(
		&transform.Forward,
		&transform.Up,
		&transform.Right)

	scale := algebra.Matrix{}
	scale.InitScale(&transform.Scale)

	tr := algebra.Matrix{}
	trs := algebra.Matrix{}

	translation.Mul(rotation, &tr)
	tr.Mul(scale, &trs)

	trs.Inverse(c.View)
}

// UpdatePerspective Update the perspective
func (c *ComponentCamera) UpdatePerspective(width, height int32, o algebra.PerspectiveOptions) {

	c.Projection.InitPerspective(algebra.PerspectiveOptions{
		Fov:         algebra.DegToRad(o.Fov),
		AspectRatio: (math.Floor(float64(width)*o.PixelRatio) / math.Floor(float64(height)*o.PixelRatio)),
		Near:        o.Near,
		Far:         o.Far,
	})
}
