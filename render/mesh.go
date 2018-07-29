package render

import (
	"log"

	gl "github.com/chsc/gogl/gl21"

	"github.com/therohans/mesh/geometry"
)

const (
	// SizeOfFloat Go's sizeOf(float)
	SizeOfFloat = 4
	// SizeOfInt Go's sizeOf(int)
	SizeOfInt = 2
)

// MeshResource a recipt from teh GPU to point to the buffers
type MeshResource struct {
	Vbo         gl.Uint
	Ibo         gl.Uint
	Size        uint
	VertBuffer  []gl.Float
	IndexBuffer []uint16
}

// Mesh a polyhedron and metadata
type Mesh struct {
	Poly     geometry.Polyhedron
	Resource MeshResource
}

// PolyToGpu send a polygon to the GPU
func PolyToGpu(p geometry.Polyhedron) Mesh {
	indexLen := len(p.GetIndices())

	verts := geometry.VertexBuffer(p)
	index := geometry.IndexBuffer(p)

	var vertexBuffer gl.Uint
	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER,
		gl.Sizeiptr(len(verts)*SizeOfFloat),
		gl.Pointer(&verts[0]),
		gl.STATIC_DRAW)

	if gl.GetError() != gl.NO_ERROR {
		log.Printf("Vertex bind buffer")
	}

	var indexBuffer gl.Uint
	gl.GenBuffers(1, &indexBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER,
		gl.Sizeiptr(len(index)*SizeOfInt),
		gl.Pointer(&index[0]),
		gl.STATIC_DRAW)

	if gl.GetError() != gl.NO_ERROR {
		log.Printf("Index bind buffer")
	}

	return Mesh{
		Poly: p,
		Resource: MeshResource{
			Vbo:         vertexBuffer,
			Ibo:         indexBuffer,
			Size:        uint(indexLen),
			VertBuffer:  verts,
			IndexBuffer: index,
		},
	}
}
