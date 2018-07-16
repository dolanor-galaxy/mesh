package algebra_test

import (
	"testing"

	"github.com/therohans/mesh/algebra"
)

func TestConsts(t *testing.T) {
	if algebra.Forward != algebra.AxisZ {
		t.Errorf("Forward direction is not Z")
	}
	if algebra.Up != algebra.AxisY {
		t.Errorf("Up direction is not Y")
	}
	if algebra.Right != algebra.AxisX {
		t.Errorf("Right direction is not X")
	}
}

func TestLength(t *testing.T) {
	v := algebra.Vector{X: 1.0, Y: 2.0, Z: 3.0}

	expected := 3.7416573867739413
	actual := v.Length()

	if actual != expected {
		t.Errorf("Length: %v should be %v", actual, expected)
	}
}

func TestMax(t *testing.T) {
	v := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}

	expected := 55.8
	actual := v.Max()

	if actual != expected {
		t.Errorf("Max: %v should be %v", actual, expected)
	}
}

func TestScale(t *testing.T) {
	v := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}

	expected := algebra.Vector{X: 5.8, Y: 111.6, Z: 1.796}
	actual := v.Scale(2)

	if actual != expected {
		t.Errorf("Scale: %v should be %v", actual, expected)
	}
}

func TestMaxV(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := algebra.Vector{X: 9.3, Y: 5.8, Z: 33.898}

	expected := algebra.Vector{X: 9.3, Y: 55.8, Z: 33.898}
	actual := a.MaxV(&b)

	if actual != expected {
		t.Errorf("MaxV: %v should be %v", actual, expected)
	}
}

func TestDot(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := algebra.Vector{X: 9.3, Y: 5.8, Z: 33.898}

	expected := 381.050404
	actual := a.Dot(&b)

	if actual != expected {
		t.Errorf("Dot: %v should be %v", actual, expected)
	}
}

func TestCross(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := algebra.Vector{X: 9.3, Y: 5.8, Z: 33.898}

	expected := algebra.Vector{
		X: 1886.3000000000002,
		Y: -89.95280000000001,
		Z: -502.12000000000006,
		W: 0,
	}
	actual := a.Cross(&b)

	if actual != expected {
		t.Errorf("Cross: %v should be %v", actual, expected)
	}
}

func TestNormalzed(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}

	expected := algebra.Vector{
		X: 0.051894578693085794,
		Y: 0.9985232727842025,
		Z: 0.016069424712548637,
	}
	actual := a.Normalized()

	if actual != expected {
		t.Errorf("Normalized: %v should be %v", actual, expected)
	}
}

func TestAddV(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := algebra.Vector{X: 9.3, Y: 5.8, Z: 33.898}

	expected := algebra.Vector{
		X: 12.200000000000001,
		Y: 61.599999999999994,
		Z: 34.79600000000001,
	}
	actual, _ := a.AddV(&b)

	if actual != expected {
		t.Errorf("AddV: %v should be %v", actual, expected)
	}
}

func TestAdd(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := 4.0

	expected := algebra.Vector{
		X: 6.9,
		Y: 59.8,
		Z: 4.898,
	}
	actual, _ := a.Add(b)

	if actual != expected {
		t.Errorf("Add: %v should be %v", actual, expected)
	}
}

func TestSubV(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := algebra.Vector{X: 9.3, Y: 5.8, Z: 33.898}

	expected := algebra.Vector{
		X: -6.4,
		Y: 50,
		Z: -33,
	}
	actual, _ := a.SubV(&b)

	if actual != expected {
		t.Errorf("SubV: %v should be %v", actual, expected)
	}
}

func TestSub(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := 4.0

	expected := algebra.Vector{
		X: -1.1,
		Y: 51.8,
		Z: -3.102,
	}
	actual, _ := a.Sub(b)

	if actual != expected {
		t.Errorf("Sub: %v should be %v", actual, expected)
	}
}

func TestMulV(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := algebra.Vector{X: 9.3, Y: 5.8, Z: 33.898}

	expected := algebra.Vector{
		X: 26.970000000000002,
		Y: 323.64,
		Z: 30.440404000000004,
	}
	actual, _ := a.MulV(&b)

	if actual != expected {
		t.Errorf("MulV: %v should be %v", actual, expected)
	}
}

func TestMul(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := 4.0

	expected := algebra.Vector{
		X: 11.6,
		Y: 223.2,
		Z: 3.592,
	}
	actual, _ := a.Mul(b)

	if actual != expected {
		t.Errorf("Mul: %v should be %v", actual, expected)
	}
}

func TestDivV(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := algebra.Vector{X: 9.3, Y: 5.8, Z: 33.898}

	expected := algebra.Vector{
		X: 0.31182795698924726,
		Y: 9.620689655172413,
		Z: 0.02649123842114579,
	}
	actual, _ := a.DivV(&b)

	if actual != expected {
		t.Errorf("DivV: %v should be %v", actual, expected)
	}
}

func TestDiv(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := 4.0

	expected := algebra.Vector{
		X: 0.725,
		Y: 13.95,
		Z: 0.2245,
	}
	actual, _ := a.Div(b)

	if actual != expected {
		t.Errorf("Div: %v should be %v", actual, expected)
	}
}

func TestDivV0(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := algebra.Vector{X: .0, Y: .0, Z: .0}

	actual, err := a.DivV(&b)

	if err == nil {
		t.Errorf("Div by 0 should error")
	}
	if actual != algebra.VectorIdentity {
		t.Errorf("Should return the Identity vector")
	}
}

func TestDiv0(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := .0

	actual, err := a.Div(b)

	if err == nil {
		t.Errorf("Div by 0 should error")
	}
	if actual != algebra.VectorIdentity {
		t.Errorf("Should return the Identity vector")
	}
}