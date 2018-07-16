package algebra

import (
	"math"
)

// DegToRad convert degrees to radians
func DegToRad(d float64) float64 {
	return d * math.Pi / 180
}
