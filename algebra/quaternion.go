package algebra

import "math"

// Quaternion is a geometrical operator to represent the relationship
// (relative length and relative orientation) between two vectors in
// 3D space.
type Quaternion Vector

// SetFromVector sets the quats values given an axis and an angle (the angles are Rads *not* degrees)
func (q *Quaternion) SetFromVector(v *Vector, angle float64) {
	sinHalfAngle := math.Sin(angle * .5)
	cosHalfAngle := math.Cos(angle * .5)

	q.X = v.X * sinHalfAngle
	q.Y = v.Y * sinHalfAngle
	q.Z = v.Z * sinHalfAngle
	q.W = cosHalfAngle
}

// Length return the length of the quat
func (q *Quaternion) Length() float64 {
	return math.Sqrt(q.X*q.X + q.Y*q.Y + q.Z*q.Z + q.W*q.W)
}

// Normalized Quaternion but with norm (length) 1
func (q *Quaternion) Normalized(out *Quaternion) {
	length := q.Length()
	out.X = q.X / length
	out.Y = q.Y / length
	out.Z = q.Z / length
	out.W = q.W / length
}

// Conjugate negate the Quaternion
func (q *Quaternion) Conjugate(out *Quaternion) {
	out.X = -q.X
	out.Y = -q.Y
	out.Z = -q.Z
	out.W = q.W
}

// Mul multiply the quat by a number
func (q *Quaternion) Mul(r float64, out *Quaternion) {
	out.X = q.X * r
	out.Y = q.Y * r
	out.Z = q.Z * r
	out.W = q.W * r
}

// Dot dot product between two rotations.
func (q *Quaternion) Dot(r Quaternion) float64 {
	return q.X*r.X + q.Y*r.Y + q.Z*r.Z + q.W*r.W
}

// MulQ multiply a quat by a quat
func (q *Quaternion) MulQ(r Quaternion, out *Quaternion) {
	out.W = q.W*r.W - q.X*r.X - q.Y*r.Y - q.Z*r.Z
	out.X = q.X*r.W + q.W*r.X + q.Y*r.Z - q.Z*r.Y
	out.Y = q.Y*r.W + q.W*r.Y + q.Z*r.X - q.X*r.Z
	out.Z = q.Z*r.W + q.W*r.Z + q.X*r.Y - q.Y*r.X
}

// MulV multiply the quat by a vector
func (q *Quaternion) MulV(r Vector, out *Vector) {
	out.W = -q.X*r.X - q.Y*r.Y - q.Z*r.Z
	out.X = q.W*r.X + q.Y*r.Z - q.Z*r.Y
	out.Y = q.W*r.Y + q.Z*r.X - q.X*r.Z
	out.Z = q.W*r.Z + q.X*r.Y - q.Y*r.X
}

// SubQ subtract a quat from this quat
func (q *Quaternion) SubQ(r Quaternion, out *Quaternion) {
	out.X = q.X - r.X
	out.Y = q.Y - r.Y
	out.Z = q.Z - r.Z
	out.W = q.W - r.W
}

// AddQ add a quat to this quat
func (q *Quaternion) AddQ(r Quaternion, out *Quaternion) {
	out.X = q.X + r.X
	out.Y = q.Y + r.Y
	out.Z = q.Z + r.Z
	out.W = q.W + r.W
}
