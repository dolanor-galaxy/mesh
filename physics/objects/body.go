package objects

import (
	"github.com/therohans/mesh/algebra"
)

const (
	// Dynamic A dynamic body is fully simulated. Can be moved manually by the user, but normally they move according to forces. A dynamic body can collide with all body types. A dynamic body always has finite, non-zero mass.
	Dynamic = 1
	// Static A static body does not move during simulation and behaves as if it has infinite mass. Static bodies can be moved manually by setting the position of the body. The velocity of a static body is always zero. Static bodies do not collide with other static or kinematic bodies.
	Static = 2
	// Kinematic A kinematic body moves under simulation according to its velocity. They do not respond to forces. They can be moved manually, but normally a kinematic body is moved by setting its velocity. A kinematic body behaves as if it has infinite mass. Kinematic bodies do not collide with other static or kinematic bodies.
	Kinematic = 4
)

const (
	Awake    = 0
	Sleepy   = 1
	Sleeping = 2
)

// BodyOptions options
type BodyOptions struct {
	Position        algebra.Vector
	Velocity        algebra.Vector
	AngularVelocity algebra.Vector
	Quaternion      algebra.Quaternion
	Mass            float32
	// Material             physics.Material
	Type                 int
	LinearDamping        float32
	AngularDamping       float32
	AllowSleep           bool
	SleepSpeedLimit      float32
	SleepTimeLimit       float32
	CollisionFilterGroup int
	CollisionFilterMast  int
	FixedRotation        bool
	LinearFactor         algebra.Vector
	AngularFactor        algebra.Vector
	// Shape                physics.Shape
}

// Body physics rigid body
type Body struct {
	// Options all the properties of this body
	Options BodyOptions
	// SleepState Current sleep state.
	SleepState             int
	wakeUpAfterNarrowphase bool
}

// WakeUp Wake the body up.
func (b *Body) WakeUp() {
	s := b.SleepState
	b.SleepState = Awake
	b.wakeUpAfterNarrowphase = false
	if s == Sleeping {
		// this.dispatchEvent(Body.wakeupEvent)
	}
}

// Sleep Force body sleep
func (b *Body) Sleep() {
	b.SleepState = Sleeping
	b.Options.Velocity.Set(0, 0, 0, 0)
	b.Options.AngularVelocity.Set(0, 0, 0, 0)
	b.wakeUpAfterNarrowphase = false
}
