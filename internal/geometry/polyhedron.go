package geometry

// Polyhedron a mesh
type Polyhedron struct {
	Vertices []Vertex
	Indices  []uint16
}

// GetVertices get the array of verts for this mesh
func (p *Polyhedron) GetVertices() []Vertex {
	return p.Vertices
}

// GetIndices get the indices of this polygon
func (p *Polyhedron) GetIndices() []uint16 {
	return p.Indices
}
