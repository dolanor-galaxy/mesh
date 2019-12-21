package main

import (
	"log"
	"math"
	"os"
	"runtime"

	"github.com/robrohan/mesh/internal/algebra"
	"github.com/robrohan/mesh/internal/core"
	"github.com/robrohan/mesh/internal/model"
	"github.com/robrohan/mesh/internal/render"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winTitle  = "Mesh Test"
	winWidth  = 800
	winHeight = 600
)

func main() {
	if err := run(); err != nil {
		log.Printf("error: %s", err)
		os.Exit(1)
	}
}

func run() error {
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
		sdl.WINDOW_OPENGL|sdl.WINDOW_RESIZABLE) // |sdl.WINDOW_FULLSCREEN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	context, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}
	defer sdl.GLDeleteContext(context)

	return GameLoop(window)
}

// GameLoop main game loop
func GameLoop(window *sdl.Window) error {
	var running bool
	var event sdl.Event
	///////////////////////////////////
	lastTime := sdl.GetTicks()
	nbFrames := 0
	///////////////////////////////////

	w := int32(winWidth)
	h := int32(winHeight)

	// Much cooler...
	// mode, err := sdl.GetCurrentDisplayMode(0)
	// if err == nil {
	// 	log.Printf("Using CurrentDisplayMode\n")
	// 	w = mode.W
	// 	h = mode.H
	// }
	log.Printf("W: %v H: %v\n", w, h)

	settings := core.Settings{
		Width:  w,
		Height: h,
	}

	renderSystem := render.System{}
	renderSystem.Initialize(settings)

	///////////////////////////////////
	scene, camera := buildTestScene(&settings)
	///////////////////////////////////

	tempQ := algebra.Quaternion{}

	running = true
	for running {
		///////////////////////////////////
		// Get input
		for event = sdl.PollEvent(); event != nil; event =
			sdl.PollEvent() {
			switch t := event.(type) {
			// switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				xrot := float32(t.X) / 2
				yrot := float32(t.Y) / 2
				log.Printf("x: %v y: %v", xrot, yrot)
				// camera.Parent.Transform.Position.X += float64(xrot)
				// camera.Parent.Transform.Position.Y += float64(yrot)
			}
		}

		tempQ.SetFromVector(&algebra.AxisZ, algebra.DegToRad((float64(lastTime))))
		cameraEntity := camera.GetParent()
		cameraEntity.Transform.Position.Y = 3 * math.Sin(float64(lastTime)/10)
		cameraEntity.Transform.Position.Z = 3 * math.Sin(float64(lastTime)/10)
		// TODO: bad name.
		cameraEntity.Transform.RotationMatrix(&tempQ)

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

	return nil
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
	entity.Transform.Position.Z = -8
	entity.Name = "Test Model"
	render := render.NewComponentRender()
	render.Mesh = mesh
	render.Material = material
	entity.Attach(&render)

	camera := core.Entity{
		Transform: core.NewTransform(),
	}
	camera.Transform.Position.Y = 2
	camera.Name = "Test Camera"
	cameraComp := core.NewComponentCamera()
	cameraComp.UpdatePerspective(s.Width, s.Height, algebra.PerspectiveOptions{
		Fov:        120,
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
