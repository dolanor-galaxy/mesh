package algebra_test

import (
	"testing"

	"github.com/robrohan/mesh/algebra"
)

func TestDistance(t *testing.T) {
	v := algebra.Vector{X: 1.0, Y: 1.0, Z: 1.0}
	v2 := algebra.Vector{X: 3.0, Y: 2.5, Z: 3.0}

	expected := 10.25 // algebra.Vector{X: 2.0, Y: 1.5, Z: 2.0}
	actual := algebra.Distance(v, v2)

	if actual != expected {
		t.Errorf("Distance: %v should be %v", actual, expected)
	}
}

func TestDirection(t *testing.T) {
	v := algebra.Vector{X: 1.0, Y: 1.0, Z: 1.0}
	v2 := algebra.Vector{X: 3.0, Y: 2.5, Z: 3.0}

	expected := algebra.Vector{
		X: -0.6246950475544243,
		Y: -0.4685212856658182,
		Z: -0.6246950475544243,
	}
	actual := algebra.Direction(v, v2)

	if actual != expected {
		t.Errorf("Direction: %v should be %v", actual, expected)
	}
}

func TestAngle(t *testing.T) {
	v := algebra.Vector{X: 1.0, Y: 1.0, Z: 1.0}
	v2 := algebra.Vector{X: 3.0, Y: 2.5, Z: 3.0}

	expected := 0.08299792509963226
	actual := algebra.Angle(v, v2)

	if actual != expected {
		t.Errorf("Angle: %v should be %v", actual, expected)
	}
}

func TestDegToRad(t *testing.T) {
	expected := 0.5934119456780721
	actual := algebra.DegToRad(34)

	if actual != expected {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}
