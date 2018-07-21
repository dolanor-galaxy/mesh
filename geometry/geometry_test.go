package geometry_test

import (
	"reflect"
	"testing"

	"github.com/therohans/mesh/algebra"

	"github.com/therohans/mesh/geometry"
)

func makePolygon() geometry.Polyhedron {
	return geometry.Polyhedron{
		Vertices: []geometry.Vertex{
			{Pos: algebra.Vector{1, 2, 3, 0},
				Color:    algebra.Vector{.3, .3, .3, 0},
				TexCoord: algebra.Vector{},
				Normal:   algebra.Vector{},
				Tangent:  algebra.Vector{},
			},
			{Pos: algebra.Vector{4, 5, 6, 0},
				Color:    algebra.Vector{.4, .4, .4, 0},
				TexCoord: algebra.Vector{},
				Normal:   algebra.Vector{},
				Tangent:  algebra.Vector{}},
			{Pos: algebra.Vector{7, 8, 9, 0},
				Color:    algebra.Vector{.5, .5, .5, 0},
				TexCoord: algebra.Vector{},
				Normal:   algebra.Vector{},
				Tangent:  algebra.Vector{}},
		},
		Indices: []uint16{0, 1, 2},
	}
}

func TestVertexBuffer(t *testing.T) {
	p := makePolygon()

	expected := []float32{
		1, 2, 3, 0.3, 0.3, 0.3, 0, 0, 0,
		0, 0, 0, 0, 0, 4, 5, 6, 0.4, 0.4,
		0.4, 0, 0, 0, 0, 0, 0, 0, 0, 7, 8,
		9, 0.5, 0.5, 0.5, 0, 0, 0, 0, 0, 0,
		0, 0}

	actual := geometry.VertexBuffer(p)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}

func TestIndexBuffer(t *testing.T) {
	p := makePolygon()

	expected := []uint16{0, 1, 2}

	actual := geometry.IndexBuffer(p)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}
