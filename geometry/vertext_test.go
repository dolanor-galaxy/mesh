package geometry_test

import (
	"testing"

	"github.com/therohans/mesh/algebra"

	"github.com/therohans/mesh/geometry"
)

func TestVertex(t *testing.T) {
	vx := geometry.Vertex{}
	v := algebra.Vector{}

	if vx.Pos != v {
		t.Errorf("Vertex: should be not be nil")
	}
}
