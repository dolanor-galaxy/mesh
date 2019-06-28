package algebra_test

import (
	"fmt"
	"testing"

	"github.com/robrohan/mesh/algebra"
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
	actual := algebra.Vector{}
	a.MaxV(b, &actual)

	if actual != expected {
		t.Errorf("MaxV: %v should be %v", actual, expected)
	}
}

func TestDot(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := algebra.Vector{X: 9.3, Y: 5.8, Z: 33.898}

	expected := 381.050404
	actual := a.Dot(b)

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
	actual := algebra.Vector{}
	a.Cross(b, &actual)

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
	actual := algebra.Vector{}
	a.Normalized(&actual)

	if actual != expected {
		t.Errorf("Normalized: %v should be %v", actual, expected)
	}
}

func TestNormalzedZero(t *testing.T) {
	a := algebra.Vector{X: 0, Y: 0, Z: 0}

	expected := algebra.Vector{X: 0, Y: 0, Z: 0}
	actual := algebra.Vector{}
	a.Normalized(&actual)

	if actual != expected {
		t.Errorf("Normalized: %v should be %v", actual, expected)
	}
}

func TestNorm(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}

	expected := 55.88252324296031
	actual := a.Norm()

	if actual != expected {
		t.Errorf("Norm: %v should be %v", actual, expected)
	}
}

func TestAlmostEquals(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := algebra.Vector{X: 2.900000001, Y: 55.8000000001, Z: .8980000001}

	expected := true
	actual := a.AlmostEquals(&b)

	if actual != expected {
		t.Errorf("Almost Equals: %v", actual)
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
	actual := algebra.Vector{}
	a.AddV(b, &actual)

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
	actual := algebra.Vector{}
	a.Add(b, &actual)

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
	actual := algebra.Vector{}
	a.SubV(b, &actual)

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
	actual := algebra.Vector{}
	a.Sub(b, &actual)

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
	actual := algebra.Vector{}
	a.MulV(b, &actual)

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
	actual := algebra.Vector{}
	a.Mul(b, &actual)

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
	actual := algebra.Vector{}
	a.DivV(b, &actual)

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
	actual := algebra.Vector{}
	a.Div(b, &actual)

	if actual != expected {
		t.Errorf("Div: %v should be %v", actual, expected)
	}
}

func TestDivV0(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}
	b := algebra.Vector{X: .0, Y: .0, Z: .0}

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should panic")
		}
	}()
	actual := algebra.Vector{}
	a.DivV(b, &actual)
	fmt.Printf("%v", actual)

	if actual != a {
		t.Errorf("Should not get here")
	}
}

func TestDiv0(t *testing.T) {
	a := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}

	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("Should panic")
		}
	}()
	actual := algebra.Vector{}
	a.Div(.0, &actual)
	fmt.Printf("%v", actual)

	if actual != a {
		t.Errorf("Should not get here")
	}
}

func TestVecAbs(t *testing.T) {
	expected := algebra.Vector{X: 2.9, Y: 55.8, Z: .898}

	a := algebra.Vector{X: -2.9, Y: 55.8, Z: -.898}

	actual := algebra.Vector{}
	a.Abs(&actual)

	if actual != expected {
		t.Errorf("Abs %v should be %v", actual, expected)
	}
}

func TestVecMagnitude(t *testing.T) {
	expected := 55.88252324296031

	a := algebra.Vector{X: -2.9, Y: 55.8, Z: -.898}

	actual := a.Magnitude()

	if actual != expected {
		t.Errorf("Magnitude %v should be %v", actual, expected)
	}
}

func TestCopy(t *testing.T) {
	a := algebra.Vector{X: -2.9, Y: 55.8, Z: -.898}
	b := algebra.Vector{X: 21, Y: 43, Z: 134}

	b.Copy(a)

	if &b == &a {
		t.Errorf("Copy used the same memory address")
	}
	if a.X != b.X || a.Y != b.Y || a.Z != b.Z {
		t.Errorf("Copy values differ %v %v", a, b)
	}
}

func TestClone(t *testing.T) {
	input := algebra.Vector{X: -2.9, Y: 55.8, Z: -.898}
	actual := algebra.Vector{}

	input.Clone(&actual)

	if &input == &actual {
		t.Errorf("Clone used the same memory address")
	}
	if input.X != actual.X || input.Y != actual.Y || input.Z != actual.Z {
		t.Errorf("Clone values differ %v %v", input, actual)
	}
}

func TestIsZero(t *testing.T) {
	input := algebra.Vector{X: 0, Y: 0, Z: 0}
	expected := true

	actual := input.IsZero()

	if actual != expected {
		t.Errorf("Is Zero is wrong")
	}
}

func TestIsZero2(t *testing.T) {
	input := algebra.Vector{X: 0, Y: 0.1, Z: 0}
	expected := false

	actual := input.IsZero()

	if actual != expected {
		t.Errorf("Is Zero is wrong")
	}
}

func TestNegate(t *testing.T) {
	input := algebra.Vector{X: 21, Y: 3.1, Z: 38, W: 123}
	expected := algebra.Vector{X: -21, Y: -3.1, Z: -38, W: -123}
	actual := algebra.Vector{}

	input.Negate(&actual)

	if actual != expected {
		t.Errorf("Negate is wrong")
	}
}
