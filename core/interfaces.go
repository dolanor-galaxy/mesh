package core

// Initializer can be initialized
type Initializer interface {
	Initialize()
}

// EntityHolder can hold other entities
type EntityHolder interface {
	Add(*Entity)
	Remove(*Entity)
	All() *[]Entity
}

// EntityIterator Something that iterates over a group of entities
type EntityIterator interface {
	HasNext() bool
	Next() *Entity
}

// ComponentHolder can hold components
type ComponentHolder interface {
	Attach(Componenter)
	Detach(Componenter)
}

// Child a child of an entity
type Child interface {
	Parent() *Entity
	SetParent(*Entity)
}

// EntitySystem an larger item that operates on entities
type EntitySystem interface {
	Configure(Settings)
	Initializer
}

type Componenter interface {
	GetParent() *Entity
	SetParent(*Entity)
}

//////////////////////////////////////////////////

// Renderer component that can render itself
// type Renderer interface {
// 	Render(RenderSystem)
// }

// Updater a component that can update itself
type Updater interface {
	Update(float64)
}

// Inputter a component that can take user input
type Inputter interface {
	Input(float64)
}

// type RenderComponent struct {
// 	Component
// 	Mesh     render.Mesh
// 	Material render.Material
// }

// func (rc RenderComponent) Renderer(rs RenderSystem, dt float64) {
// 	rs.Render(rc.Mesh, rc.Material)
// }
