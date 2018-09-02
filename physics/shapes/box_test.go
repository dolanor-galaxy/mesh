package physics_test

import (
	"testing"

	"github.com/therohans/mesh/algebra"
	physics "github.com/therohans/mesh/physics/shapes"
)

func TestBoxCalculateLocalInertia(t *testing.T) {
	actual := algebra.Vector{}
	expected := algebra.Vector{
		X: 1.6666666666666665,
		Y: 1.6666666666666665,
		Z: 1.6666666666666665,
	}
	half := algebra.Vector{
		X: .5,
		Y: .5,
		Z: .5,
	}
	box := physics.NewBox(half)

	box.CalculateLocalInertia(10, &actual)

	if actual != expected {
		t.Errorf("CalculateLocalInertia did something odd %v", actual)
	}
}

func TestBoxGetSideNormals(t *testing.T) {
	half := algebra.Vector{
		X: .5,
		Y: .5,
		Z: .5,
	}
	box := physics.NewBox(half)
	quat := algebra.Quaternion{}
	quat.SetFromVector(&algebra.AxisZ, algebra.DegToRad(90))

	actual := [6]algebra.Vector{}
	expected := [6]algebra.Vector{
		algebra.Vector{X: 0.3535533905932738, Y: 0.35355339059327373, Z: 0, W: -0},
		algebra.Vector{X: -0.35355339059327373, Y: 0.3535533905932738, Z: 0, W: -0},
		algebra.Vector{X: 0, Y: 0, Z: 0.3535533905932738, W: -0.35355339059327373},
		algebra.Vector{X: -0.3535533905932738, Y: -0.35355339059327373, Z: 0, W: 0},
		algebra.Vector{X: 0.35355339059327373, Y: -0.3535533905932738, Z: 0, W: 0},
		algebra.Vector{X: 0, Y: 0, Z: -0.5, W: 0},
	}

	box.GetSideNormals(&actual, quat)

	if actual != expected {
		t.Errorf("CalculateLocalInertia did something odd %v", actual)
	}
}
