package render

import (
	"github.com/therohans/mesh/core"
)

// ComponentRender draw an object on screen
type ComponentRender struct {
	*core.Component
	Mesh     Mesh
	Material Material
}

func NewComponentRender() ComponentRender {
	return ComponentRender{
		Component: &core.Component{
			Parent: &core.Entity{},
		},
	}
}
