package geometry_test

import (
	"testing"

	"github.com/robrohan/mesh/internal/algebra"
	"github.com/robrohan/mesh/internal/geometry"
)

func TestVertex(t *testing.T) {
	vx := geometry.Vertex{}
	v := algebra.Vector{}

	if vx.Pos != v {
		t.Errorf("Vertex: should be not be nil")
	}

	if geometry.VertexSize != 14 {
		t.Errorf("Vertex Size: should be 14")
	}
}
