package geometry

import (
	gl "github.com/chsc/gogl/gl21"
)

// VertexBuffer creates a buffer for sending to opengl
func VertexBuffer(p Polyhedron) []gl.Float {
	vertices := p.GetVerticies()
	vlen := len(vertices)
	buffer := make([]gl.Float, vlen*int(VertexSize), vlen*int(VertexSize))

	for i := 0; i < vlen; i++ {
		row := int(i * int(VertexSize))
		buffer[row] = gl.Float(vertices[i].Pos.X)
		buffer[row+1] = gl.Float(vertices[i].Pos.Y)
		buffer[row+2] = gl.Float(vertices[i].Pos.Z)

		buffer[row+3] = gl.Float(vertices[i].Color.X)
		buffer[row+4] = gl.Float(vertices[i].Color.Y)
		buffer[row+5] = gl.Float(vertices[i].Color.Z)

		buffer[row+6] = gl.Float(vertices[i].TexCoord.X)
		buffer[row+7] = gl.Float(vertices[i].TexCoord.Y)

		buffer[row+8] = gl.Float(vertices[i].Normal.X)
		buffer[row+9] = gl.Float(vertices[i].Normal.Y)
		buffer[row+10] = gl.Float(vertices[i].Normal.Z)

		buffer[row+11] = gl.Float(vertices[i].Tangent.X)
		buffer[row+12] = gl.Float(vertices[i].Tangent.Y)
		buffer[row+13] = gl.Float(vertices[i].Tangent.Z)
	}
	return buffer
}

// IndexBuffer get the polygons index buffer
func IndexBuffer(p Polyhedron) []uint16 {
	return p.GetIndices()
}
