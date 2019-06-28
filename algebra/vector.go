package algebra

import (
	"math"
)

// Normalize vector in the same direction but with norm (length) 1
func Normalize(out *Vector) {
	len := out.Length()
	out.X = out.X / len
	out.Y = out.Y / len
	out.Z = out.Z / len
}

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
	// Precision used with 'Almost Equals' when comparing vectors
	Precision = 1e-6
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
func (v *Vector) MaxV(o Vector, out *Vector) {
	out.X = math.Max(o.X, v.X)
	out.Y = math.Max(o.Y, v.Y)
	out.Z = math.Max(o.Z, v.Z)
	out.W = math.Max(o.W, v.W)
}

// Dot dot product of two vectors
func (v *Vector) Dot(o Vector) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

// Cross a vector that is perpendicular to both a and b
func (v *Vector) Cross(o Vector, out *Vector) {
	out.X = v.Y*o.Z - v.Z*o.Y
	out.Y = v.Z*o.X - v.X*o.Z
	out.Z = v.X*o.Y - v.Y*o.X
	out.W = 0.0
}

// Normalized vector in the same direction but with norm (length) 1
func (v *Vector) Normalized(out *Vector) {
	len := v.Length()
	if len == 0 {
		out.X = 0
		out.Y = 0
		out.Z = 0
		return
	}
	out.X = v.X / len
	out.Y = v.Y / len
	out.Z = v.Z / len
}

// Norm assign a strictly positive length or size
func (v *Vector) Norm() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2) + math.Pow(v.Z, 2))
}

// AddV adds two vectors together
func (v *Vector) AddV(r Vector, out *Vector) {
	out.X = v.X + r.X
	out.Y = v.Y + r.Y
	out.Z = v.Z + r.Z
}

// Add add a value to a vector
func (v *Vector) Add(r float64, out *Vector) {
	out.X = v.X + r
	out.Y = v.Y + r
	out.Z = v.Z + r
}

// SubV subtracts two vectors
func (v *Vector) SubV(r Vector, out *Vector) {
	out.X = v.X - r.X
	out.Y = v.Y - r.Y
	out.Z = v.Z - r.Z
}

// Sub subtracts a value to a vector
func (v *Vector) Sub(r float64, out *Vector) {
	out.X = v.X - r
	out.Y = v.Y - r
	out.Z = v.Z - r
}

// MulV multiply one vector with another
func (v *Vector) MulV(r Vector, out *Vector) {
	out.X = v.X * r.X
	out.Y = v.Y * r.Y
	out.Z = v.Z * r.Z
}

// Mul multiply a vector by a number
func (v *Vector) Mul(r float64, out *Vector) {
	out.X = v.X * r
	out.Y = v.Y * r
	out.Z = v.Z * r
}

// DivV divide a vector by another vector
func (v *Vector) DivV(r Vector, out *Vector) {
	if r.X == .0 || r.Y == .0 || r.Z == 0 {
		panic("Can not divide vector by vector with a zero")
	}
	out.X = v.X / r.X
	out.Y = v.Y / r.Y
	out.Z = v.Z / r.Z
}

// Div divide a vector by a number
func (v *Vector) Div(r float64, out *Vector) {
	if r == .0 {
		panic("Can not divide vector by zero")
	}
	out.X = v.X / r
	out.Y = v.Y / r
	out.Z = v.Z / r
}

// Abs absolute value of this vector
func (v *Vector) Abs(out *Vector) {
	out.X = math.Abs(v.X)
	out.Y = math.Abs(v.Y)
	out.Z = math.Abs(v.Z)
}

// Magnitude the magnitude of this vector
func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

// Set sets this vector's elements
func (v *Vector) Set(x, y, z, w float64) {
	v.X = x
	v.Y = y
	v.Z = z
	v.W = w
}

// IsZero returns true if xy and z are zero
func (v *Vector) IsZero() bool {
	return v.X == 0 && v.Y == 0 && v.Z == 0
}

// Negate Make the vector point in the opposite direction.
func (v *Vector) Negate(target *Vector) {
	target.X = -v.X
	target.Y = -v.Y
	target.Z = -v.Z
	target.W = -v.W
}

// AlmostEquals Check if a vector is almost equal to another one.
func (v *Vector) AlmostEquals(nv *Vector) bool {
	if math.Abs(v.X-nv.X) > Precision ||
		math.Abs(v.Y-nv.Y) > Precision ||
		math.Abs(v.Z-nv.Z) > Precision {
		return false
	}
	return true
}

// Copy sets the properties of this vector to the given vector
func (v *Vector) Copy(vec Vector) {
	v.X = vec.X
	v.Y = vec.Y
	v.Z = vec.Z
	v.W = vec.W
}

// Clone return a new instance of this vector
func (v *Vector) Clone(out *Vector) {
	out.X = v.X
	out.Y = v.Y
	out.Z = v.Z
	out.W = v.W
}
