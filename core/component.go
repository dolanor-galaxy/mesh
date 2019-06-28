package core

// Component a behaviour that is usually attached to an Entity
type Component struct {
	Parent *Entity
	Initializer
	Child
}

// GetParent get the entity this component is attached
func (rc *Component) GetParent() *Entity {
	return rc.Parent
}

// SetParent sets this components parent
func (rc *Component) SetParent(e *Entity) {
	rc.Parent = e
}
