package main

import (
	"log"
	"runtime"

	"github.com/therohans/mesh/algebra"
	"github.com/therohans/mesh/core"
	"github.com/therohans/mesh/model"
	"github.com/therohans/mesh/render"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winTitle  = "Mesh Test"
	winWidth  = 640
	winHeight = 480
)

func main() {
	var window *sdl.Window
	var context sdl.GLContext

	var err error
	runtime.LockOSThread()
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	sdl.GLSetAttribute(sdl.GL_RED_SIZE, 5)
	sdl.GLSetAttribute(sdl.GL_GREEN_SIZE, 5)
	sdl.GLSetAttribute(sdl.GL_BLUE_SIZE, 5)
	sdl.GLSetAttribute(sdl.GL_DEPTH_SIZE, 16)
	sdl.GLSetAttribute(sdl.GL_DOUBLEBUFFER, 1)

	window, err = sdl.CreateWindow(
		winTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight,
		sdl.WINDOW_OPENGL) //|sdl.WINDOW_FULLSCREEN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	context, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}
	defer sdl.GLDeleteContext(context)

	GameLoop(window)
}

// GameLoop main game loop
func GameLoop(window *sdl.Window) {
	var running bool
	var event sdl.Event
	///////////////////////////////////
	lastTime := sdl.GetTicks()
	nbFrames := 0
	///////////////////////////////////

	settings := core.Settings{
		Width:  winWidth,
		Height: winHeight,
	}

	renderSystem := render.System{}
	renderSystem.Initialize(settings)

	///////////////////////////////////
	scene, camera := buildTestScene(&settings)
	///////////////////////////////////

	running = true
	for running {
		///////////////////////////////////
		// Get input
		for event = sdl.PollEvent(); event != nil; event =
			sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				// xrot = float32(t.Y) / 2
				// yrot = float32(t.X) / 2
				log.Printf("[%dms]MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
			}
		}

		///////////////////////////////////
		// Render
		camera.UpdateViewMatrix()
		renderSystem.RenderScene(scene)
		window.GLSwap()

		///////////////////////////////////
		// Measure speed
		currentTime := sdl.GetTicks()
		nbFrames++
		if currentTime-lastTime >= 1.0 {
			// log.Printf("%v ms / Frame\n", 1000.0/nbFrames)
			nbFrames = 0
			lastTime += 1.0
		}
		///////////////////////////////////
		// time.Sleep(50 * time.Millisecond)
	}
}

func buildTestScene(s *core.Settings) (*core.Scene, *core.ComponentCamera) {
	///////////////////////////////////
	scene := core.Scene{}

	poly, err := model.CreateTestPoly()
	if err != nil {
		panic("Can't load test object")
	}
	// Send the object the GPU (create buffers)
	mesh := render.CreateMesh(poly)
	shader := render.Shader{
		Name:    "default",
		Program: render.UseProgram(),
	}
	material := render.Material{
		Shader: shader,
	}

	entity := core.Entity{
		Transform: core.NewTransform(),
	}
	entity.Transform.Position.Z = -5
	entity.Name = "Test Model"
	render := render.NewComponentRender()
	render.Mesh = mesh
	render.Material = material
	entity.Attach(&render)

	camera := core.Entity{
		Transform: core.NewTransform(),
	}
	camera.Transform.Position.Y = 2
	camera.Transform.Position.Z = 5
	camera.Name = "Test Camera"
	cameraComp := core.NewComponentCamera()
	cameraComp.UpdatePerspective(s.Width, s.Height, algebra.PerspectiveOptions{
		Fov:        70,
		Near:       0.1,
		Far:        1000,
		PixelRatio: 1,
	})
	camera.Attach(&cameraComp)

	scene.Add(&camera)
	scene.Add(&entity)
	scene.ActiveCamera = &camera
	///////////////////////////////////

	return &scene, &cameraComp
}
