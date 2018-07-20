package algebra

import "math"

// Quaternion is a geometrical operator to represent the relationship
// (relative length and relative orientation) between two vectors in
// 3D space.
type Quaternion Vector

// FromVector sets the quats values given an axis and an angle
// (the angles are Rads *not* degrees)
func (q *Quaternion) FromVector(v *Vector, angle float64) *Quaternion {
	sinHalfAngle := math.Sin(angle * .5)
	cosHalfAngle := math.Cos(angle * .5)

	q.X = v.X * sinHalfAngle
	q.Y = v.Y * sinHalfAngle
	q.Z = v.Z * sinHalfAngle
	q.W = cosHalfAngle

	return q
}

// Length return the length of the quat
func (q *Quaternion) Length() float64 {
	return math.Sqrt(q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W)
}

// Normalized Quaternion but with norm (length) 1
func (q *Quaternion) Normalized() Quaternion {
	length := q.Length()
	return Quaternion{
		q.X / length,
		q.Y / length,
		q.Z / length,
		q.W / length,
	}
}

// Conjugate negate the Quaternion
func (q *Quaternion) Conjugate() Quaternion {
	return Quaternion{
		-q.X,
		-q.Y,
		-q.Z,
		q.W,
	}
}

// Mul multiply the quat by a number
func (q *Quaternion) Mul(r float64) Quaternion {
	return Quaternion{q.X * r, q.Y * r, q.Z * r, q.W * r}
}

// Dot dot product between two rotations.
func (q *Quaternion) Dot(r Quaternion) float64 {
	return q.X*r.X + q.Y*r.Y + q.Z*r.Z + q.W*r.W
}

// MulQ multiply a quat by a quat
func (q *Quaternion) MulQ(r Quaternion) Quaternion {
	w := q.W*r.W - q.X*r.X - q.Y*r.Y - q.Z*r.Z
	x := q.X*r.W + q.W*r.X + q.Y*r.Z - q.Z*r.Y
	y := q.Y*r.W + q.W*r.Y + q.Z*r.X - q.X*r.Z
	z := q.Z*r.W + q.W*r.Z + q.X*r.Y - q.Y*r.X
	return Quaternion{x, y, z, w}
}

// MulV multiply the quat by a vector
func (q *Quaternion) MulV(r Vector) Quaternion {
	w := -q.X*r.X - q.Y*r.Y - q.Z*r.Z
	x := q.W*r.X + q.Y*r.Z - q.Z*r.Y
	y := q.W*r.Y + q.Z*r.X - q.X*r.Z
	z := q.W*r.Z + q.X*r.Y - q.Y*r.X
	return Quaternion{x, y, z, w}
}

// SubQ subtract a quat from this quat
func (q *Quaternion) SubQ(r Quaternion) Quaternion {
	return Quaternion{q.X - r.X, q.Y - r.Y, q.Z - r.Z, q.W - r.W}
}

// AddQ add a quat to this quat
func (q *Quaternion) AddQ(r Quaternion) Quaternion {
	return Quaternion{q.X + r.X, q.Y + r.Y, q.Z + r.Z, q.W + r.W}
}
