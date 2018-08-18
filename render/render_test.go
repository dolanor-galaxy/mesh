package render_test

import (
	"testing"

	"github.com/therohans/mesh/core"
	"github.com/therohans/mesh/render"
)

func InitMockRenderSystem(t *testing.T) render.System {
	s := core.Settings{
		Width:  300,
		Height: 300,
	}

	mockInit := func(w, h int32) error {
		if w != s.Width || h != s.Height {
			t.Errorf("InitSystem failed")
		}
		return nil
	}

	rs := render.System{}
	rs.InitSystem(s, mockInit)

	return rs
}
func TestRenderDraw(t *testing.T) {
	m := render.Mesh{
		Name: "mesh1",
	}
	mat := render.Material{
		Name: "mat1",
	}

	rs := InitMockRenderSystem(t)

	mockDraw := func(mesh render.Mesh, material render.Material) error {
		return nil
	}

	err := rs.Draw(m, mat, mockDraw)
	if err != nil {
		t.Errorf("Draw call failed")
	}

}
