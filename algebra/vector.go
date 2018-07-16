package algebra

import (
	"errors"
	"math"
)

// Vector a point in 3d space (or 2d)
type Vector struct {
	X float64
	Y float64
	Z float64
	W float64
}

var (
	// AxisX a vector representing X direction
	AxisX = Vector{X: 1}
	// AxisY a vector representing Y direction
	AxisY = Vector{Y: 1}
	// AxisZ a vector representing Z direction
	AxisZ = Vector{Z: 1}
	// AxisAll a vector all directions
	AxisAll = Vector{X: 1, Y: 1, Z: 1}
	// VectorIdentity 0,0,0
	VectorIdentity = Vector{}
	// Forward a vector representing forward direction
	Forward = AxisZ
	// Up a vector representing up direction
	Up = AxisY
	// Right a vector representing right direction
	Right = AxisX
)

// Length return the length of the vector
func (v *Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Max return the longest axis value
func (v *Vector) Max() float64 {
	return math.Max(v.X, math.Max(v.Y, v.Z))
}

// Scale scales the vector
func (v *Vector) Scale(s float64) Vector {
	return Vector{v.X * s, v.Y * s, v.Z * s, v.W}
}

// MaxV return a new vector with the largest parts of each
func (v *Vector) MaxV(o *Vector) Vector {
	return Vector{
		X: math.Max(o.X, v.X),
		Y: math.Max(o.Y, v.Y),
		Z: math.Max(o.Z, v.Z),
		W: math.Max(o.W, v.W),
	}
}

// Dot dot product of two vectors
func (v *Vector) Dot(o *Vector) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

// Cross a vector that is perpendicular to both a and b
func (v *Vector) Cross(o *Vector) Vector {
	return Vector{
		v.Y*o.Z - v.Z*o.Y,
		v.Z*o.X - v.X*o.Z,
		v.X*o.Y - v.Y*o.X,
		0.0,
	}
}

// Normalized vector in the same direction but with norm (length) 1
func (v *Vector) Normalized() Vector {
	len := v.Length()
	return Vector{
		X: v.X / len,
		Y: v.Y / len,
		Z: v.Z / len,
	}
}

// AddV adds two vectors together
func (v *Vector) AddV(r *Vector) (Vector, error) {
	return Vector{
		X: v.X + r.X,
		Y: v.Y + r.Y,
		Z: v.Z + r.Z,
	}, nil
}

// Add add a value to a vector
func (v *Vector) Add(r float64) (Vector, error) {
	return Vector{
		X: v.X + r,
		Y: v.Y + r,
		Z: v.Z + r,
	}, nil
}

// SubV subtracts two vectors
func (v *Vector) SubV(r *Vector) (Vector, error) {
	return Vector{
		X: v.X - r.X,
		Y: v.Y - r.Y,
		Z: v.Z - r.Z,
	}, nil
}

// Sub subtracts a value to a vector
func (v *Vector) Sub(r float64) (Vector, error) {
	return Vector{
		X: v.X - r,
		Y: v.Y - r,
		Z: v.Z - r,
	}, nil
}

// MulV multiply one vector with another
func (v *Vector) MulV(r *Vector) (Vector, error) {
	return Vector{
		X: v.X * r.X,
		Y: v.Y * r.Y,
		Z: v.Z * r.Z,
	}, nil
}

// Mul multiply a vector by a number
func (v *Vector) Mul(r float64) (Vector, error) {
	return Vector{
		X: v.X * r,
		Y: v.Y * r,
		Z: v.Z * r,
	}, nil
}

// DivV divide a vector by another vector
func (v *Vector) DivV(r *Vector) (Vector, error) {
	if r.X == .0 || r.Y == .0 || r.Z == 0 {
		return VectorIdentity, errors.New("Can not divide by zero")
	}
	return Vector{
		X: v.X / r.X,
		Y: v.Y / r.Y,
		Z: v.Z / r.Z,
	}, nil
}

// Div divide a vector by a number
func (v *Vector) Div(r float64) (Vector, error) {
	if r == .0 {
		return VectorIdentity, errors.New("Can not divide by zero")
	}
	return Vector{
		X: v.X / r,
		Y: v.Y / r,
		Z: v.Z / r,
	}, nil
}
