package algebra

import "math"

// Distance get the distance between two vectors
func Distance(p Vector, p2 Vector) float64 {
	dx := p2.X - p.X
	dy := p2.Y - p.Y
	dz := p2.Z - p.Z
	return dx*dx + dy*dy + dz*dz
}

// Direction get the direction from one vector towards another
func Direction(p Vector, p2 Vector) Vector {
	v := p.SubV(p2)
	return v.Normalized()
}

// Angle get the angle between two points (in radians)
func Angle(p Vector, p2 Vector) float64 {
	cross := p.Cross(p2)
	return math.Atan2(cross.Norm(), p.Dot(p2))
}

// DegToRad convert degrees to radians
func DegToRad(d float64) float64 {
	return d * math.Pi / 180
}
