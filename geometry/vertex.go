package geometry

import "github.com/robrohan/mesh/algebra"

const (
	// VertexSize number of elements in a vertex
	VertexSize uint8 = 14
)

// Vertex an element of some 3D geometry which has a position and some other attributes
type Vertex struct {
	Pos      algebra.Vector
	Color    algebra.Vector
	TexCoord algebra.Vector
	Normal   algebra.Vector
	Tangent  algebra.Vector
}
