package physics

import (
	"log"
	"math"

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

// Temp Vector
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

// Temp Vectors
var cb = algebra.Vector{}
var ab = algebra.Vector{}

// ComputeNormal Get face normal given 3 vertices
func (c *ConvexPolyhedron) ComputeNormal(va algebra.Vector, vb algebra.Vector, vc algebra.Vector, target *algebra.Vector) {
	vb.SubV(va, &ab)
	vc.SubV(vb, &cb)
	cb.Cross(ab, target)
	if !target.IsZero() {
		target.Normalized(target)
	}
}

// GetFaceNormal Compute the normal of a face from its vertices
func (c *ConvexPolyhedron) GetFaceNormal(i uint32, target *algebra.Vector) {
	var f = c.Faces[i]
	var va = c.Vertices[f[0]]
	var vb = c.Vertices[f[1]]
	var vc = c.Vertices[f[2]]
	c.ComputeNormal(va, vb, vc, target)
}

// ComputeNormals Compute the normals of the faces. Will reuse existing Vec3 objects in
// the .faceNormals array if they exist.
func (c *ConvexPolyhedron) ComputeNormals() {

	if len(c.FaceNormals) <= 0 {
		// this.faceNormals.length = this.faces.length;
		faceLen := len(c.Faces)

		// Generate normals
		for i := 0; i < faceLen; i++ {
			// Check so all vertices exists for this face
			for j := 0; j < len(c.Faces[i]); j++ {
				// if(!this.vertices[this.faces[i][j]]){
				if len(c.Vertices) > int(c.Faces[i][j]) {
					log.Printf("Vertex %v not found!", c.Faces[i][j])
				}
			}

			// var n = this.faceNormals[i] || new Vec3();
			n := algebra.Vector{}

			c.GetFaceNormal(uint32(i), &n)
			n.Negate(&n)
			c.FaceNormals[i] = n
			var vertex = c.Vertices[c.Faces[i][0]]

			if n.Dot(vertex) < 0 {
				log.Printf("FaceNormals[%v] = Vec3(%v) looks like it points into the shape? The vertices follow.", i, n)
				log.Printf("Make sure they are ordered CCW around the normal, using the right hand rule.")
				for j := 0; j < len(c.Faces[i]); j++ {
					log.Printf("Vertices[%v] = Vec3(%v)", c.Faces[i][j], c.Vertices[c.Faces[i][j]])
				}
			}
		}
	}
}

var (
	cliAabbmin = algebra.Vector{}
	cliAabbmax = algebra.Vector{}
)

// CalculateLocalInertia
func (c *ConvexPolyhedron) CalculateLocalInertia(mass float64, target *algebra.Vector) {
	// Approximate with box inertia
	// Exact inertia calculation is overkill, but
	// see http://geometrictools.com/Documentation/PolyhedralMassProperties.pdf for the correct way to do it
	c.ComputeLocalAABB(&cliAabbmin, &cliAabbmax)
	x := cliAabbmax.X - cliAabbmin.X
	y := cliAabbmax.Y - cliAabbmin.Y
	z := cliAabbmax.Z - cliAabbmin.Z
	target.X = 1.0 / 12.0 * mass * (2*y*2*y + 2*z*2*z)
	target.Y = 1.0 / 12.0 * mass * (2*x*2*x + 2*z*2*z)
	target.Z = 1.0 / 12.0 * mass * (2*y*2*y + 2*x*2*x)
}

// Volume Get approximate convex volume
func (c *ConvexPolyhedron) Volume() float64 {
	return 4.0 * math.Pi * c.BoundingSphereRadius / 3.0
}

// UpdateBoundingSphereRadius
func (c *ConvexPolyhedron) UpdateBoundingSphereRadius() {
	// Assume points are distributed with local (0,0,0) as center
	max2 := 0.0
	verts := c.Vertices
	nn := len(verts)
	for i := 0; i != nn; i++ {
		// var norm2 = verts[i].norm2()
		norm2 := verts[i].Dot(verts[i])
		if norm2 > max2 {
			max2 = norm2
		}
	}
	c.BoundingSphereRadius = math.Sqrt(max2)
}

// var computeLocalAABBWorldVert = algebra.Vector{}

// ComputeLocalAABB
func (c *ConvexPolyhedron) ComputeLocalAABB(aabbmin *algebra.Vector, aabbmax *algebra.Vector) {
	n := len(c.Vertices)
	vertices := c.Vertices
	// worldVert := &computeLocalAABBWorldVert

	aabbmin.Set(math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, 0)
	aabbmax.Set(-math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64, 0)

	for i := 0; i < n; i++ {
		var v = vertices[i]
		if v.X < aabbmin.X {
			aabbmin.X = v.X
		} else if v.X > aabbmax.X {
			aabbmax.X = v.X
		}
		if v.Y < aabbmin.Y {
			aabbmin.Y = v.Y
		} else if v.Y > aabbmax.Y {
			aabbmax.Y = v.Y
		}
		if v.Z < aabbmin.Z {
			aabbmin.Z = v.Z
		} else if v.Z > aabbmax.Z {
			aabbmax.Z = v.Z
		}
	}
}

var tempWorldVertex = algebra.Vector{}

// CalculateWorldAABB
func (c *ConvexPolyhedron) CalculateWorldAABB(pos *algebra.Vector, quat *algebra.Quaternion, min *algebra.Vector, max *algebra.Vector) {
	n := len(c.Vertices)
	verts := c.Vertices
	minx := min.X
	miny := min.Y
	minz := min.Z

	maxx := max.X
	maxy := max.Y
	maxz := max.Z

	for i := 0; i < n; i++ {
		tempWorldVertex.Copy(verts[i])
		// quat.vmult(tempWorldVertex, tempWorldVertex)
		quat.MulV(tempWorldVertex, &tempWorldVertex)
		// pos.vadd(tempWorldVertex, tempWorldVertex)
		pos.AddV(tempWorldVertex, &tempWorldVertex)

		var v = tempWorldVertex
		if v.X < minx {
			minx = v.X
		} else if v.X > maxx {
			maxx = v.X
		}

		if v.Y < miny {
			miny = v.Y
		} else if v.Y > maxy {
			maxy = v.Y
		}

		if v.Z < minz {
			minz = v.Z
		} else if v.Z > maxz {
			maxz = v.Z
		}
	}

	min.Set(minx, miny, minz, 0)
	max.Set(maxx, maxy, maxz, 0)
}
