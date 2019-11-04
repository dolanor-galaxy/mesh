package render

import (
	"github.com/robrohan/mesh/internal/core"
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
