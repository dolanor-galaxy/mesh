package physics_test

import (
	"testing"

	"github.com/therohans/mesh/model"
	physics "github.com/therohans/mesh/physics/shapes"

	"github.com/therohans/mesh/algebra"
)

func makeTestObject() ([]algebra.Vector, [][]uint32) {
	// Create a test cube
	model, _ := model.CreateTestPoly()

	verts := model.GetVerticies()
	index := model.GetIndices()

	points := []algebra.Vector{}
	for i := 0; i < len(verts); i++ {
		points = append(points, verts[i].Pos)
	}

	faces := [][]uint32{}
	for i := 0; i < len(index); i += 3 {
		faces = append(faces, []uint32{
			uint32(index[i]),
			uint32(index[i+1]),
			uint32(index[i+2]),
		})
	}
	return points, faces
}

func TestNewConvexPolyhedron(t *testing.T) {
	points, faces := makeTestObject()

	poly, _ := physics.NewConvexPolyhedron(
		points, faces, []algebra.Vector{}, []algebra.Vector{},
	)

	poly.ComputeEdges()

	if len(poly.UniqueEdges) == 0 {
		t.Errorf("Unique edges didn't get made %v", poly.UniqueEdges)
	}

	expected := algebra.Vector{
		X: -0.7071067811865475, Y: 0.7071067811865475, Z: 0, W: 0}
	if poly.UniqueEdges[len(poly.UniqueEdges)-1] != expected {
		t.Errorf("Unique edges didn't get made correctly %v", poly.UniqueEdges)
	}
}

func TestComputeEdges(t *testing.T) {
	points, faces := makeTestObject()

	poly, _ := physics.NewConvexPolyhedron(
		points, faces, []algebra.Vector{}, []algebra.Vector{},
	)

	poly.ComputeEdges()

	if len(poly.UniqueEdges) == 0 {
		t.Errorf("Unique edges didn't get made %v", poly.UniqueEdges)
	}

	expected := algebra.Vector{
		X: -0.7071067811865475, Y: 0.7071067811865475, Z: 0, W: 0}
	if poly.UniqueEdges[len(poly.UniqueEdges)-1] != expected {
		t.Errorf("Unique edges didn't get made correctly %v", poly.UniqueEdges)
	}
}
