package algebra_test

import (
	"testing"

	"github.com/robrohan/mesh/algebra"
)

func TestSetFromVector(t *testing.T) {
	expected := algebra.Quaternion{
		X: 0.7071067811865475,
		Y: 1.414213562373095,
		Z: 3.5355339059327373,
		W: 0.7071067811865476,
	}
	q := algebra.Quaternion{}
	v := algebra.Vector{
		X: 1, Y: 2, Z: 5,
	}

	q.SetFromVector(&v, algebra.DegToRad(90))

	if q != expected {
		t.Errorf("FromVector did something odd %v", q)
	}
}

func TestQuatLength(t *testing.T) {
	q := algebra.Quaternion{X: 1.0, Y: 2.0, Z: 3.0, W: 2.0}

	expected := 4.242640687119285
	actual := q.Length()

	if actual != expected {
		t.Errorf("Length: %v should be %v", actual, expected)
	}
}

func TestQuatNormalzed(t *testing.T) {
	q := algebra.Quaternion{X: 2.9, Y: 55.8, Z: .898, W: 4.3}

	expected := algebra.Quaternion{
		X: 0.051741627257717866,
		Y: 0.9955802762002265,
		Z: 0.016022062509458843,
		W: 0.076720343864892,
	}
	actual := algebra.Quaternion{}
	q.Normalized(&actual)

	if actual != expected {
		t.Errorf("Normalized: %v should be %v", actual, expected)
	}
}

func TestQuatConjugate(t *testing.T) {
	q := algebra.Quaternion{X: 2.9, Y: 55.8, Z: .898, W: 4.3}

	expected := algebra.Quaternion{
		X: -2.9,
		Y: -55.8,
		Z: -.898,
		W: 4.3,
	}
	actual := algebra.Quaternion{}
	q.Conjugate(&actual)

	if actual != expected {
		t.Errorf("Conjugate: %v should be %v", actual, expected)
	}
}

func TestQuatMul(t *testing.T) {
	q := algebra.Quaternion{X: 2.9, Y: 55.8, Z: .898, W: 4.3}

	expected := algebra.Quaternion{
		X: 5.8,
		Y: 111.6,
		Z: 1.796,
		W: 8.6,
	}
	actual := algebra.Quaternion{}
	q.Mul(2, &actual)

	if actual != expected {
		t.Errorf("Mul: %v should be %v", actual, expected)
	}
}

func TestQuatDot(t *testing.T) {
	q := algebra.Quaternion{X: 2.9, Y: 55.8, Z: .898, W: 4.3}
	q2 := algebra.Quaternion{X: .9, Y: 5.8, Z: 2.2, W: 4.3}

	expected := 346.7156

	actual := q.Dot(q2)

	if actual != expected {
		t.Errorf("Dot: %v should be %v", actual, expected)
	}
}

func TestQuatMulQ(t *testing.T) {
	q := algebra.Quaternion{X: 2.9, Y: 55.8, Z: .898, W: 4.3}
	q2 := algebra.Quaternion{X: .9, Y: 5.8, Z: 2.2, W: 4.3}

	expected := algebra.Quaternion{
		X: 133.89159999999998,
		Y: 259.3082,
		Z: -20.078599999999998,
		W: -309.7356,
	}

	actual := algebra.Quaternion{}
	q.MulQ(q2, &actual)

	if actual != expected {
		t.Errorf("MulQ: %v should be %v", actual, expected)
	}
}

func TestQuatMulV(t *testing.T) {
	q := algebra.Quaternion{X: 2.9, Y: 55.8, Z: .898, W: 4.3}
	v := algebra.Vector{X: .9, Y: 5.8, Z: 2.2}

	expected := algebra.Vector{
		X: 121.42160000000001,
		Y: 19.368199999999998,
		Z: -23.939999999999998,
		W: -328.2256,
	}

	actual := algebra.Vector{}
	q.MulV(v, &actual)

	if actual != expected {
		t.Errorf("MulV: %v should be %v", actual, expected)
	}
}

func TestQuatSubQ(t *testing.T) {
	q := algebra.Quaternion{X: 2.9, Y: 55.8, Z: .898, W: 4.3}
	q2 := algebra.Quaternion{X: .9, Y: 5.8, Z: 2.2, W: 4.3}

	expected := algebra.Quaternion{X: 2, Y: 50, Z: -1.302, W: 0}

	actual := algebra.Quaternion{}
	q.SubQ(q2, &actual)

	if actual != expected {
		t.Errorf("SubQ: %v should be %v", actual, expected)
	}
}

func TestQuatAddQ(t *testing.T) {
	q := algebra.Quaternion{X: 2.9, Y: 55.8, Z: .898, W: 4.3}
	q2 := algebra.Quaternion{X: .9, Y: 5.8, Z: 2.2, W: 4.3}

	expected := algebra.Quaternion{
		X: 3.8, Y: 61.599999999999994, Z: 3.0980000000000003, W: 8.6}

	actual := algebra.Quaternion{}
	q.AddQ(q2, &actual)

	if actual != expected {
		t.Errorf("AddQ: %v should be %v", actual, expected)
	}
}
