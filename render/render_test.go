package render_test

import (
	"testing"

	"github.com/therohans/mesh/core"
	"github.com/therohans/mesh/render"
)

func TestRenderSystemCreate(t *testing.T) {
	s := core.Settings{
		Width:  300,
		Height: 300,
	}

	mockOpenGl := func(w, h int32) error {
		if w != s.Width || h != s.Height {
			t.Errorf("InitSystem failed")
		}
		return nil
	}

	rs := render.System{}
	rs.InitSystem(s, mockOpenGl)
}

func TestRenderDraw(t *testing.T) {
	m := render.Mesh{
		ID: "mesh1",
	}
	mat := render.Material{
		Name: "mat1",
	}

	mockDraw := func(mesh Mesh, material Material) error {
		if w != s.Width || h != s.Height {
			t.Errorf("InitSystem failed")
		}
		return nil
	}

	rs := render.System{}
	rs.InitSystem(s, mockOpenGl)
}
