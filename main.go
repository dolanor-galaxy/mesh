package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/therohans/mesh/algebra"
	"github.com/therohans/mesh/geometry"

	"github.com/therohans/mesh/render"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	winTitle  = "OpenGL Shader"
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
	window, err = sdl.CreateWindow(
		winTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	context, err = window.GLCreateContext()
	if err != nil {
		panic(err)
	}
	defer sdl.GLDeleteContext(context)

	render.InitOpenGl(winWidth, winHeight)

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

	poly, err := buildTestObject()
	if err != nil {
		panic("Can't load test object")
	}
	mesh := render.PolyToGpu(poly)

	// Load the program and bind all the locations
	program := render.UseProgram()

	running = true
	for running {

		for event = sdl.PollEvent(); event != nil; event =
			sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				// xrot = float32(t.Y) / 2
				// yrot = float32(t.X) / 2
				fmt.Printf("[%dms]MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
			}
		}
		render.StartDrawGl()
		render.DrawGl(mesh, program)
		render.EndDrawGl()

		window.GLSwap()

		///////////////////////////////////
		// Measure speed
		currentTime := sdl.GetTicks()
		nbFrames++
		if currentTime-lastTime >= 1.0 {
			log.Printf("%v ms / Frame\n", 1000.0/nbFrames)
			nbFrames = 0
			lastTime += 1.0
		}
		///////////////////////////////////
		// time.Sleep(50 * time.Millisecond)
	}
}

///////////////////////////////////////////////////////////

func buildTestObject() (geometry.Polyhedron, error) {
	modelData := phoCube
	vertLen := len(modelData)
	voffset := vertLen / 6

	verts := make([]geometry.Vertex, voffset, voffset)
	index := make([]uint16, voffset, voffset)

	poly := geometry.Polyhedron{
		Vertices: verts,
		Indices:  index,
	}

	numRow := vertLen / 6
	for i := 0; i < numRow; i++ {
		row := i * 6
		poly.Vertices[i] = geometry.Vertex{
			Pos: algebra.Vector{
				X: modelData[row+0],
				Y: modelData[row+1],
				Z: modelData[row+2],
				W: 0,
			},
			Color: algebra.Vector{
				X: modelData[row+3],
				Y: modelData[row+4],
				Z: modelData[row+5],
				W: 1,
			},
		}
		// fmt.Printf("%v\n", row)
		poly.Indices[i] = uint16(i)
		fmt.Printf("%v", poly.Vertices[i])
	}
	fmt.Printf("%v", poly.Indices)
	return poly, nil
}

var phoCube = []float64{
	-1.0, -1.0, -1.0, 0.583, 0.771, 0.014,
	-1.0, -1.0, 1.0, 0.609, 0.115, 0.436,
	-1.0, 1.0, 1.0, 0.327, 0.483, 0.844,
	1.0, 1.0, -1.0, 0.822, 0.569, 0.201,
	-1.0, -1.0, -1.0, 0.435, 0.602, 0.223,
	-1.0, 1.0, -1.0, 0.310, 0.747, 0.185,
	1.0, -1.0, 1.0, 0.597, 0.770, 0.761,
	-1.0, -1.0, -1.0, 0.559, 0.436, 0.730,
	1.0, -1.0, -1.0, 0.359, 0.583, 0.152,
	1.0, 1.0, -1.0, 0.483, 0.596, 0.789,
	1.0, -1.0, -1.0, 0.559, 0.861, 0.639,
	-1.0, -1.0, -1.0, 0.195, 0.548, 0.859,
	-1.0, -1.0, -1.0, 0.014, 0.184, 0.576,
	-1.0, 1.0, 1.0, 0.771, 0.328, 0.970,
	-1.0, 1.0, -1.0, 0.406, 0.615, 0.116,
	1.0, -1.0, 1.0, 0.676, 0.977, 0.133,
	-1.0, -1.0, 1.0, 0.971, 0.572, 0.833,
	-1.0, -1.0, -1.0, 0.140, 0.616, 0.489,
	-1.0, 1.0, 1.0, 0.997, 0.513, 0.064,
	-1.0, -1.0, 1.0, 0.945, 0.719, 0.592,
	1.0, -1.0, 1.0, 0.543, 0.021, 0.978,
	1.0, 1.0, 1.0, 0.279, 0.317, 0.505,
	1.0, -1.0, -1.0, 0.167, 0.620, 0.077,
	1.0, 1.0, -1.0, 0.347, 0.857, 0.137,
	1.0, -1.0, -1.0, 0.055, 0.953, 0.042,
	1.0, 1.0, 1.0, 0.714, 0.505, 0.345,
	1.0, -1.0, 1.0, 0.783, 0.290, 0.734,
	1.0, 1.0, 1.0, 0.722, 0.645, 0.174,
	1.0, 1.0, -1.0, 0.302, 0.455, 0.848,
	-1.0, 1.0, -1.0, 0.225, 0.587, 0.040,
	1.0, 1.0, 1.0, 0.517, 0.713, 0.338,
	-1.0, 1.0, -1.0, 0.053, 0.959, 0.120,
	-1.0, 1.0, 1.0, 0.393, 0.621, 0.362,
	1.0, 1.0, 1.0, 0.673, 0.211, 0.457,
	-1.0, 1.0, 1.0, 0.820, 0.883, 0.371,
	1.0, -1.0, 1.0, 0.982, 0.099, 0.879}
