package render_test

import (
	"testing"

	"github.com/robrohan/mesh/render"
)

func TestShaderCreate(t *testing.T) {
	s := render.Shader{
		Name: "My Shader",
	}

	if s.Name != "My Shader" {
		t.Errorf("Expected a shader got %v", s)
	}
}
