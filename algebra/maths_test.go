package algebra_test

import (
	"testing"

	"github.com/therohans/mesh/algebra"
)

func TestDegToRad(t *testing.T) {
	expected := 0.5934119456780721
	actual := algebra.DegToRad(34)

	if actual != expected {
		t.Errorf("Expected %v got %v", expected, actual)
	}
}
