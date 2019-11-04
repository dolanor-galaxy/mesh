package render_test

import (
	"os"
	"testing"

	"github.com/robrohan/mesh/internal/render"
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
	txt := render.ReadVertexShader("./testdata", "Simple.glsl")

	if txt == "" {
		t.Errorf("Could not read file %v", txt)
	}
}

func TestReadFragmentShader(t *testing.T) {
	txt := render.ReadFragmentShader("./testdata", "Simple.glsl")

	if txt == "" {
		t.Errorf("Could not read file %v", txt)
	}
}

// func TestCreateProgram(t *testing.T) {
// 	defer testChdir(t, "../")()

// 	frag := render.ReadFragmentShader("Simple.glsl")
// 	vert := render.ReadVertexShader("Simple.glsl")

// 	_, err := render.CreateProgram(vert, frag)
// 	if err != nil {
// 		t.Errorf("Error creating %v", err)
// 	}

// 	// fmt.Printf("%v", prog)
// }
