package algebra_test

import (
	"testing"

	"github.com/therohans/mesh/algebra"
)

func TestInitIdentity(t *testing.T) {
	m := algebra.Matrix{}
	m.InitIdentity()
	if m[0][0] != 1 {
		t.Errorf("Init did something odd")
	}
}

func TestInitTranslation(t *testing.T) {
	m := algebra.Matrix{}
	p := algebra.Vector{X: 2.0, Y: 3.0, Z: 4.0, W: .0}

	m.InitTranslation(p)

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

	m.InitRotation(p)

	if m != expected {
		t.Errorf("InitRotate did something odd: %v", m)
	}
}

func TestInitScale(t *testing.T) {
	expected := algebra.Matrix{
		{2.0, 0.0, 0.0, 0.0},
		{0.0, 3.0, 0.0, 0.0},
		{0.0, 0.0, 4.0, 0.0},
		{0.0, 0.0, 0.0, 1.0},
	}

	m := algebra.Matrix{}
	p := algebra.Vector{X: 2.0, Y: 3.0, Z: 4.0, W: .0}

	m.InitScale(p)

	if m != expected {
		t.Errorf("InitScale did something odd: %v", m)
	}
}

func TestMatrixInitPerspective(t *testing.T) {
	opts := algebra.PerspectiveOptions{
		Fov:         algebra.DegToRad(45),
		AspectRatio: (1024 / 768),
		ZNear:       .1,
		ZFar:        1000,
	}

	expected := algebra.Matrix{
		{0.4142135623730951, 0, 0, 0},
		{0, 0.4142135623730951, 0, 0},
		{0, 0, -1.0002000200020003, -1},
		{0, 0, -0.2000200020002, 0},
	}

	m := algebra.Matrix{}
	m.InitPerspective(opts)

	if m != expected {
		t.Errorf("Perspective did something odd: %v", m)
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

	m3 := m1.Mul(m2)

	if m3 != expected {
		t.Errorf("Expected %v to equal %v", expected, m3)
	}
}

func TestTranspose(t *testing.T) {
	expected := algebra.Matrix{
		{1.0, 0.0, 0.0, 2.0},
		{0.0, 1.0, 0.0, 3.0},
		{0.0, 0.0, 1.0, 4.0},
		{0.0, 0.0, 0.0, 1.0},
	}

	m := algebra.Matrix{}
	p := algebra.Vector{X: 2.0, Y: 3.0, Z: 4.0, W: .0}

	m.InitTranslation(p)
	m.Transpose()

	if m != expected {
		t.Errorf("Transpose did something odd: %v", m)
	}
}

func TestMatrixTransform(t *testing.T) {
	expected := algebra.Vector{
		X: 82.0, Y: 79.0, Z: 56.0, W: 37.0,
	}

	m := algebra.Matrix{
		{3.0, 2.0, 7.0, 2.0},
		{4.0, 9.0, 2.0, 3.0},
		{12.0, 4.0, 1.0, 4.0},
		{2.0, 4.0, 4.0, 1.0},
	}
	p := algebra.Vector{X: 2.0, Y: 3.0, Z: 4.0, W: 8.0}

	p2 := m.Transform(p)

	if p2 != expected {
		t.Errorf("Transform did something odd: %v", p2)
	}
}

func TestMatrixInverse(t *testing.T) {
	expected := algebra.Matrix{
		{-0.16339869281045752, -0.16993464052287582, 0.12418300653594773, 0.33986928104575165},
		{-0.1525054466230937, 0.04139433551198257, -0.017429193899782137, 0.2505446623093682},
		{0.07843137254901962, -0.07843137254901962, -0.019607843137254905, 0.15686274509803924},
		{0.6230936819172114, 0.48801742919389984, -0.10021786492374729, -1.309368191721133},
	}
	m := algebra.Matrix{
		{3.0, 2.0, 7.0, 2.0},
		{4.0, 9.0, 2.0, 3.0},
		{12.0, 4.0, 1.0, 4.0},
		{2.0, 4.0, 4.0, 1.0},
	}

	m2 := m.Inverse()

	if m2 != expected {
		t.Errorf("Inverse did something odd: %v", m2)
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

func TestMatrixCopy(t *testing.T) {
	m1 := algebra.Matrix{
		{10, 20, 30, 40},
		{10, 20, 30, 40},
		{10, 20, 30, 40},
		{10, 20, 30, 40},
	}

	m2 := m1.Copy()

	if m1 != m2 {
		t.Errorf("Copy %v to equal %v", m1, m2)
	}
	if &m1 == &m2 {
		t.Errorf("Copy %p same address %p", &m1, &m2)
	}

	m2[1][1] = 99

	if m1 == m2 {
		t.Errorf("Modified both %v and %v", m1, m2)
	}
}
