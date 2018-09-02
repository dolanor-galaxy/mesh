package physics

import physics "github.com/therohans/mesh/physics/world"

type Broadphaser interface {
	CollisionPairs(physics.World)
}
