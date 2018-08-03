package core

// Entity a gameobject, something in the system
type Entity struct {
	ID         string
	Name       string
	transform  Transform
	children   []Entity
	components []Component
	EntityHolder
	ComponentHolder
	Initializer
}

// Transform this entities position, rotation, etc
func (ge Entity) Transform() Transform {
	return ge.transform
}

// Add add a sub entity to this entity
func (ge Entity) Add(e Entity) {
	ge.children = append(ge.children, e)
}

func (ge Entity) Remove(e Entity) {

}

// Attach add a component to this entity
func (ge Entity) Attach(cmp Component) {
	cmp.SetParent(ge)
	ge.components = append(ge.components, cmp)
}

func (ge Entity) Detach(cmp Component) {

}
