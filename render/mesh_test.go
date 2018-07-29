package render_test

import (
	"reflect"
	"testing"

	"github.com/therohans/mesh/algebra"
	"github.com/therohans/mesh/geometry"
	"github.com/therohans/mesh/render"
)

func makePolygon() geometry.Polyhedron {
	return geometry.Polyhedron{
		Vertices: []geometry.Vertex{
			{Pos: algebra.Vector{X: 1, Y: 2, Z: 3, W: 0},
				Color:    algebra.Vector{X: .3, Y: .3, Z: .3},
				TexCoord: algebra.Vector{},
				Normal:   algebra.Vector{},
				Tangent:  algebra.Vector{},
			},
			{Pos: algebra.Vector{X: 4, Y: 5, Z: 6},
				Color:    algebra.Vector{X: .4, Y: .4, Z: .4},
				TexCoord: algebra.Vector{},
				Normal:   algebra.Vector{},
				Tangent:  algebra.Vector{}},
			{Pos: algebra.Vector{X: 7, Y: 8, Z: 9},
				Color:    algebra.Vector{X: .5, Y: .5, Z: .5},
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

	actual := render.VertexBuffer(p)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}

func TestIndexBuffer(t *testing.T) {
	p := makePolygon()

	expected := []uint16{0, 1, 2}

	actual := render.IndexBuffer(p)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}
