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
	VertBuffer  []float32
	IndexBuffer []uint16
}

// Mesh a polyhedron and metadata
type Mesh struct {
	Name     string
	Poly     geometry.Polyhedron
	Resource MeshResource
}

// CreateMesh send a polygon to the GPU
func CreateMesh(p geometry.Polyhedron) Mesh {
	indexLen := len(p.GetIndices())

	verts := VertexBuffer(p)
	index := IndexBuffer(p)

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

// VertexBuffer creates a buffer for sending to opengl
func VertexBuffer(p geometry.Polyhedron) []float32 {
	vertices := p.GetVerticies()
	vlen := len(vertices)
	buffer := make([]float32, vlen*int(geometry.VertexSize), vlen*int(geometry.VertexSize))

	for i := 0; i < vlen; i++ {
		row := int(i * int(geometry.VertexSize))
		buffer[row] = float32(vertices[i].Pos.X)
		buffer[row+1] = float32(vertices[i].Pos.Y)
		buffer[row+2] = float32(vertices[i].Pos.Z)

		buffer[row+3] = float32(vertices[i].Color.X)
		buffer[row+4] = float32(vertices[i].Color.Y)
		buffer[row+5] = float32(vertices[i].Color.Z)

		buffer[row+6] = float32(vertices[i].TexCoord.X)
		buffer[row+7] = float32(vertices[i].TexCoord.Y)

		buffer[row+8] = float32(vertices[i].Normal.X)
		buffer[row+9] = float32(vertices[i].Normal.Y)
		buffer[row+10] = float32(vertices[i].Normal.Z)

		buffer[row+11] = float32(vertices[i].Tangent.X)
		buffer[row+12] = float32(vertices[i].Tangent.Y)
		buffer[row+13] = float32(vertices[i].Tangent.Z)
	}
	return buffer
}

// IndexBuffer get the polygons index buffer
func IndexBuffer(p geometry.Polyhedron) []uint16 {
	return p.GetIndices()
}
