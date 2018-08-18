package render

import (
	"errors"
	"fmt"
	"log"
	"math"

	gl "github.com/chsc/gogl/gl21"
	"github.com/therohans/mesh/core"
)

// RenderInitializer initialize the render framework (opengl)
type RenderInitializer func(width int32, height int32) error
type RenderDrawer func(mesh Mesh, material Material) error

// System system used to render to the screen OpenGL 2.1
type System struct {
	core.EntitySystem
}

// Initialize the system
func (r *System) Initialize(s core.Settings) {
	r.InitSystem(s, initOpenGl)
}

// Configure the system
func (r *System) Configure(s core.Settings) {
	r.Initialize(s)
}

// InitSystem configure and startup the system
func (r *System) InitSystem(s core.Settings, fn RenderInitializer) {
	err := fn(s.Width, s.Height)
	if err != nil {
		panic("Initialize render system failed")
	}
}

// Render render a mesh
func (r *System) Render(mesh Mesh, material Material) error {
	return r.Draw(mesh, material, drawGl)
}

// Draw call the opengl draw code directly
func (r *System) Draw(mesh Mesh, mat Material, fn RenderDrawer) error {
	return fn(mesh, mat)
}

//////////////////////////////////////////////////////////////

// InitOpenGl startup OpenGl
func initOpenGl(width, height int32) error {
	gl.Init()
	version := gl.GoStringUb(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)
	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))

	// OpenGL flags
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	if gl.GetError() != gl.NO_ERROR {
		return errors.New("Initialzation failed")
	}

	return nil
}

var (
	uniRoll  float32
	uniYaw   float32 = 1.0
	uniPitch float32
	uniscale float32 = 0.3
	yrot     float32 = 20.0
	zrot     float32
	xrot     float32
)

// DrawGl draw gl
func drawGl(mesh Mesh, material Material) error {
	gl.ClearColor(1, 1, 1, 1)
	// Swap program if needed...
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	uniYaw = yrot * (math.Pi / 180.0)
	yrot = yrot - 1.0
	uniPitch = zrot * (math.Pi / 180.0)
	zrot = zrot - 0.5
	uniRoll = xrot * (math.Pi / 180.0)
	xrot = xrot - 0.2

	gl.Uniform4f(material.Shader.Program.UniScale,
		gl.Float(uniRoll),
		gl.Float(uniYaw),
		gl.Float(uniPitch),
		gl.Float(uniscale))

	err := gl.GetError()
	if err != gl.NO_ERROR {
		return fmt.Errorf("Uniform failed: %v", err)
	}

	gl.DrawElements(gl.TRIANGLES,
		gl.Sizei(mesh.Resource.Size),
		gl.UNSIGNED_SHORT,
		gl.Offset(nil, 0))

	err = gl.GetError()
	if err != gl.NO_ERROR {
		return fmt.Errorf("Draw elements failed: %v", err)
	}

	return nil
}
