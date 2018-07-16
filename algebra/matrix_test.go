package algebra_test

import (
	"testing"

	"github.com/therohans/mesh/algebra"
)

func TestInitIdentity(t *testing.T) {
	m := algebra.Matrix{}
	m.Identity()
	if m[0][0] != 1 {
		t.Errorf("Init did something odd")
	}
}

func TestInitTranslation(t *testing.T) {
	m := algebra.Matrix{}
	p := algebra.Vector{X: 2.0, Y: 3.0, Z: 4.0, W: .0}

	m.InitTranslation(&p)

	if m[3][0] != 2.0 && m[3][1] != 3.0 && m[3][2] != 4.0 {
		t.Errorf("InitTranslate did something odd")
	}
}

func TestInitRotation(t *testing.T) {
	expected := algebra.Matrix{
		{0.9961969233988566, -0.07153602925877484, -0.04974219867014598, 0},
		{0.06966087492121549, 0.9968289510967708, -0.03846303108859671, 0},
		{0.05233595624294383, 0.034851668155187324, 0.9980211966240684, 0},
		{0, 0, 0, 1},
	}

	m := algebra.Matrix{}
	p := algebra.Vector{X: 2.0, Y: 3.0, Z: 4.0, W: .0}

	m.InitRotation(&p)

	if m != expected {
		t.Errorf("InitRotate did something odd: %v", m)
	}
}

func TestMatrixMul(t *testing.T) {
	expected := algebra.Matrix{
		{10, 20, 30, 40},
		{10, 20, 30, 40},
		{10, 20, 30, 40},
		{10, 20, 30, 40},
	}

	m1 := algebra.Matrix{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}
	m2 := algebra.Matrix{
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
		{1, 2, 3, 4},
	}

	m3 := m1.Mul(&m2)

	if *m3 != expected {
		t.Errorf("Expected %v to equal %v", expected, m3)
	}
}

func TestMatrixAsArray(t *testing.T) {
	m1 := algebra.Matrix{
		{10, 20, 30, 40},
		{10, 20, 30, 40},
		{10, 20, 30, 40},
		{10, 20, 30, 40},
	}
	expected := [16]float32{
		10, 20, 30, 40,
		10, 20, 30, 40,
		10, 20, 30, 40,
		10, 20, 30, 40,
	}

	actual := m1.AsArray()

	if actual != expected {
		t.Errorf("AsArray %v to equal %v", expected, actual)
	}
}
