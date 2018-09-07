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

func TestBoxVolume(t *testing.T) {
	half := algebra.Vector{X: 2.5, Y: 2.5, Z: 2.5}
	box := physics.NewBox(half)

	expected := 125.0
	actual := box.Volume()

	if actual != expected {
		t.Errorf("Box Volume not correct %v", actual)
	}
}

func TestBoxUpdateBoundingSphereRadius(t *testing.T) {
	half := algebra.Vector{X: 1.5, Y: 1.5, Z: 1.5}
	box := physics.NewBox(half)

	expected := 2.598076211353316
	box.UpdateBoundingSphereRadius()
	actual := box.BoundingSphereRadius

	if actual != expected {
		t.Errorf("BoundingSphereRadius Volume not correct %v", actual)
	}
}

func TestBoxCalculateWorldAABB(t *testing.T) {
	expectedMin := algebra.Vector{
		X: -2.896139826697846, Y: 7.103860173302154,
		Z: -7.462019382530521, W: -0.4341204441673258,
	}
	expectedMax := algebra.Vector{
		X: 2.896139826697846, Y: 12.896139826697846,
		Z: -2.5379806174694797, W: -0.4341204441673258,
	}

	half := algebra.Vector{X: 2.5, Y: 2.5, Z: 2.5}
	box := physics.NewBox(half)

	// Pretend world position
	worldPos := algebra.Vector{X: 0, Y: 10, Z: -5}
	// Rotate 20deg about the Zed
	rot := algebra.Quaternion{}
	rot.SetFromVector(&algebra.AxisZ, algebra.DegToRad(20))

	// Should be our values
	min := algebra.Vector{}
	max := algebra.Vector{}

	box.CalculateWorldAABB(&worldPos, &rot, &min, &max)

	if expectedMin != min || expectedMax != max {
		t.Errorf("Bounding min max not correct %v %v", min, max)
	}
}
