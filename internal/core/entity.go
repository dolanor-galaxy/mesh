package core

import "fmt"

// Entity a gameobject, something in the system
type Entity struct {
	ID         string
	Name       string
	Transform  *Transform
	children   []*Entity
	components []Componenter
	EntityHolder
	ComponentHolder
	Initializer
}

// Add add a sub entity to this entity
func (ge *Entity) Add(e *Entity) {
	ge.children = append(ge.children, e)
}

func (ge *Entity) Remove(e *Entity) {
	panic("Not implemented")
}

// Attach add a component to this entity
func (ge *Entity) Attach(cmp Componenter) {
	cmp.SetParent(ge)
	ge.components = append(ge.components, cmp)
}

// Detach a component from this entity
func (ge *Entity) Detach(cmp Componenter) {
	panic("Not implemented")
}

// GetComponent get component by name
func (ge *Entity) GetComponent(t string) Componenter {
	tt := ""
	for q := 0; q < len(ge.components); q++ {
		// ew.
		tt = fmt.Sprintf("%T", ge.components[q])
		switch tt {
		case t:
			return ge.components[q]
		}
	}
	return nil
}

// func (ge *Entity) GetComponent(i interface{}) Componenter {
// 	for q := 0; q < len(ge.components); q++ {
// 		c := ge.components[q]

// 		// Nup.
// 		// convert, found := c.(i)

// 		// Nup.
// 		switch c.(type) {
// 		case i:
// 			return c
// 		default:
// 			return nil
// 		}
// 	}

// 	return nil
// }
