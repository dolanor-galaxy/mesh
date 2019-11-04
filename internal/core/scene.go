package core

type Scene struct {
	ActiveCamera *Entity
	children     []*Entity
	EntityHolder
	EntityIterator
}

// Add add a sub entity to this entity
func (s *Scene) Add(e *Entity) {
	s.children = append(s.children, e)
}

func (s *Scene) Remove(e *Entity) {
	panic("Not implemented")
}

func (s *Scene) All() []*Entity {
	return s.children
}
