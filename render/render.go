package render

import (
	"errors"
	"fmt"
	"log"

	"github.com/therohans/mesh/algebra"
	"github.com/therohans/mesh/core"

	gl "github.com/chsc/gogl/gl21"
)

type RenderCommand struct {
	// Mesh         *Mesh
	// Material     *Material
	// ModelToWorld *algebra.Matrix
	Render *ComponentRender
	Camera *core.ComponentCamera
}

// RenderInitializer initialize the render framework (opengl)
type RenderInitializer func(width int32, height int32) error
type RenderDrawer func(mesh *Mesh, material *Material) error

//////////////////////////////////////////////////////////////////////////////////////

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

func (r *System) RenderScene(s *core.Scene) error {
	log.Printf("Start render scene...\n")

	camera := s.ActiveCamera
	if camera == nil {
		panic("Scene has no active camera")
	}
	ccomp := s.ActiveCamera.GetComponent(core.ComponentTypeCamera)
	cc, ok := ccomp.(*core.ComponentCamera)
	if !ok {
		panic("Camera has no camera component")
	}

	entities := s.All()
	for t := 0; t < len(entities); t++ {
		e := entities[t]
		// fmt.Printf("Entity: %v\n", e)

		comp := e.GetComponent(core.ComponentTypeRender)
		// fmt.Printf("%v\n", comp.(core.Component).ParentEntity)
		// fmt.Printf("Component: %v\n", comp)

		if rc, ok := comp.(*ComponentRender); ok {
			r.Render(RenderCommand{
				Render: rc,
				Camera: cc,
			})
			// fmt.Printf("%v\n", e)
		} else {
			// log.Printf("%v has no render component: %v\n", e, ok)
		}
	}
	log.Printf("Done render scene.\n")
	return nil
}

// Render render a mesh
func (r *System) Render(command RenderCommand) error {
	mesh := &command.Render.Mesh
	material := &command.Render.Material

	entity := command.Render.GetParent()
	if entity == nil {
		panic("Trying to render an unattached component")
	}

	modelToWorld := entity.Transform.GetTransformation()

	// These need to come from the camera
	view := command.Camera.GetView()
	proj := command.Camera.GetProjection()

	mtw := matrixAsArray(modelToWorld)
	viewa := matrixAsArray(view)
	proja := matrixAsArray(proj)

	// fmt.Printf("MTW: %v - %v\n", mtw, material.Shader.Program.UniWorld)
	// fmt.Printf("VIEW: %v - %v\n", viewa, material.Shader.Program.UniView)
	// fmt.Printf("PROJ: %v - %v\n", proja, material.Shader.Program.UniProject)

	gl.UniformMatrix4fv(material.Shader.Program.UniWorld, gl.Sizei(1), gl.FALSE, &mtw[0])
	gl.UniformMatrix4fv(material.Shader.Program.UniView, gl.Sizei(1), gl.FALSE, &viewa[0])
	gl.UniformMatrix4fv(material.Shader.Program.UniProject, gl.Sizei(1), gl.FALSE, &proja[0])

	return r.Draw(mesh, material, drawGl)
}

// Draw call the opengl draw code directly
func (r *System) Draw(mesh *Mesh, mat *Material, fn RenderDrawer) error {
	return fn(mesh, mat)
}

//////////////////////////////////////////////////////////////

func matrixAsArray(matrix *algebra.Matrix) [16]gl.Float {
	m := matrix
	asArray := [16]gl.Float{
		gl.Float(m[0][0]), gl.Float(m[0][1]), gl.Float(m[0][2]), gl.Float(m[0][3]),
		gl.Float(m[1][0]), gl.Float(m[1][1]), gl.Float(m[1][2]), gl.Float(m[1][3]),
		gl.Float(m[2][0]), gl.Float(m[2][1]), gl.Float(m[2][2]), gl.Float(m[2][3]),
		gl.Float(m[3][0]), gl.Float(m[3][1]), gl.Float(m[3][2]), gl.Float(m[3][3]),
	}
	return asArray
}

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

// DrawGl draw gl
func drawGl(mesh *Mesh, material *Material) error {
	gl.ClearColor(1, 1, 1, 1)
	// Swap program if needed...
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

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
