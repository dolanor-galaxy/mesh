package physics

import (
	"github.com/therohans/mesh/algebra"
)

// ConvexPolyhedron A set of polygons describing a convex shape.
// The shape MUST be convex for the code to work properly. No polygons may
// be coplanar (contained in the same 3D plane), instead these should be
// merged into one polygon.
// @see http://www.altdevblogaday.com/2011/05/13/contact-generation-between-3d-convex-meshes/
// @see http://bullet.googlecode.com/svn/trunk/src/BulletCollision/NarrowPhaseCollision/btPolyhedralContactClipping.cpp
// @todo Move the clipping functions to ContactGenerator?
// @todo Automatically merge coplanar polygons in constructor.
type ConvexPolyhedron struct {
	Shape
	Vertices    []algebra.Vector
	Faces       [][]uint32
	UniqueAxes  []algebra.Vector
	UniqueEdges []algebra.Vector
	FaceNormals []algebra.Vector
}

// NewConvexPolyhedron create a new convex polyhedron
func NewConvexPolyhedron(points []algebra.Vector, faces [][]uint32, uniqueAxes []algebra.Vector, faceNormals []algebra.Vector) (*ConvexPolyhedron, error) {
	cp := ConvexPolyhedron{
		Shape: Shape{
			Type: ShapeConvexPoly,
		},
		Vertices:    points,
		Faces:       faces,
		UniqueAxes:  uniqueAxes,
		UniqueEdges: []algebra.Vector{},
		FaceNormals: faceNormals,
	}
	return &cp, nil
}

var computeEdgesTmpEdge = algebra.Vector{}

// ComputeEdges Computes uniqueEdges
func (c *ConvexPolyhedron) ComputeEdges() {
	faces := c.Faces
	vertices := c.Vertices
	edges := &c.UniqueEdges

	edge := &computeEdgesTmpEdge

	for i := 0; i != len(faces); i++ {
		face := faces[i]
		numVertices := len(face)

		for j := 0; j != numVertices; j++ {
			var k = (j + 1) % numVertices
			vertices[face[j]].SubV(vertices[face[k]], edge)
			edge.Normalized(edge) //.normalize()

			found := false
			for p := 0; p != len(*edges); p++ {
				if (*edges)[p].AlmostEquals(edge) || (*edges)[p].AlmostEquals(edge) {
					found = true
					break
				}
			}

			if !found {
				// edges.push(edge.clone())
				newVec := algebra.Vector{}
				edge.Clone(&newVec)
				c.UniqueEdges = append(c.UniqueEdges, newVec)
			}
		}
	}
}
