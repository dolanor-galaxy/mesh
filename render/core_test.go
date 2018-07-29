package render_test

import (
	"os"
	"testing"

	"github.com/therohans/mesh/render"
)

func testChdir(t *testing.T, dir string) func() {
	old, err := os.Getwd()
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("err: %s", err)
	}
	return func() {
		if err := os.Chdir(old); err != nil {
			t.Fatalf("err: %s", err)
		}
	}
}

func TestReadVertexShader(t *testing.T) {
	defer testChdir(t, "../")()

	txt := render.ReadVertexShader("Demo.glsl")

	if txt == "" {
		t.Errorf("Could not read file %v", txt)
	}
}

func TestReadFragmentShader(t *testing.T) {
	defer testChdir(t, "../")()

	txt := render.ReadFragmentShader("Demo.glsl")

	if txt == "" {
		t.Errorf("Could not read file %v", txt)
	}
}

// func TestCreateProgram(t *testing.T) {
// 	defer testChdir(t, "../")()

// 	frag := render.ReadFragmentShader("Demo.glsl")
// 	vert := render.ReadVertexShader("Demo.glsl")

// 	prog, err := render.CreateProgram(vert, frag)
// 	if err != nil {
// 		t.Errorf("Error creating %v", err)
// 	}

// 	fmt.Printf("%v", prog)
// }