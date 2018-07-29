package render

import (
	gl "github.com/chsc/gogl/gl21"
)

// System system used to render to the screen OpenGL 2.1
type System struct {
}

// Init initialize the render system
func (r *System) Init(width int32, height int32) {
	initOpenGl(width, height)
}

// Render render a mesh
func (r *System) Render(mesh Mesh, program Program) {
	gl.ClearColor(1, 1, 1, 1)
	// Swap program if needed...
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	drawGl(mesh, program)
}
