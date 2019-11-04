package core_test

import (
	"testing"

	"github.com/robrohan/mesh/internal/core"
	"github.com/robrohan/mesh/internal/render"
)

func mockEntity() *core.Entity {
	entity := core.Entity{
		Transform: core.NewTransform(),
	}
	entity.Name = "Test Model"
	render := render.NewComponentRender()
	entity.Attach(&render)

	cameraComp := core.NewComponentCamera()
	entity.Attach(&cameraComp)

	return &entity
}

func TestGetComponent(t *testing.T) {
	e := mockEntity()

	comp := e.GetComponent(core.ComponentTypeRender)

	if comp == nil {
		t.Fatalf("Could not get component from entity")
	}
}

func TestGetComponentNil(t *testing.T) {
	e := mockEntity()

	comp := e.GetComponent("doesntexist")

	if comp != nil {
		t.Fatalf("expected comp to not be found")
	}
}
