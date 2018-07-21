package geometry

// VertexBuffer creates a buffer for sending to opengl
func VertexBuffer(p Polyhedron) []float32 {
	vertices := p.GetVerticies()
	vlen := len(vertices)
	buffer := make([]float32, vlen*int(VertexSize), vlen*int(VertexSize))

	for i := 0; i < vlen; i++ {
		row := int(i * int(VertexSize))
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
func IndexBuffer(p Polyhedron) []uint16 {
	return p.GetIndices()
}
