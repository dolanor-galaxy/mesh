package core

// Component a behavior that is usually attached to an Entity
type Component struct {
	parent Entity
	Initializer
	Child
}

// Parent get the entity this component is attached
func (rc Component) Parent() Entity {
	return rc.parent
}

// SetParent sets this components parent
func (rc Component) SetParent(e Entity) {
	rc.parent = e
}
