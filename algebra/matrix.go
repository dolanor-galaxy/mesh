package algebra

import (
	"math"
)

// Matrix 4x4 matrix Row-Major Ordering
type Matrix [4][4]float64

// PerspectiveOptions used with InitPerspective
type PerspectiveOptions struct {
	Fov         float64
	AspectRatio float64
	ZNear       float64
	ZFar        float64
}

// InitIdentity create an identity matrix
//
//     0 1 2 3
//   ---------
//  0 | 1 0 0 0
//  1 | 0 1 0 0
//  2 | 0 0 1 0
//  3 | 0 0 0 1
func (m *Matrix) InitIdentity() {
	m[0][0] = 1
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[1][0] = 0
	m[1][1] = 1
	m[1][2] = 0
	m[1][3] = 0

	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = 1
	m[2][3] = 0

	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = 0
	m[3][3] = 1
}

// InitTranslation initialize the matrix with a translation point
//
//      0 1 2 3
//    ---------
//  0 | 1 0 0 0
//  1 | 0 1 0 0
//  2 | 0 0 1 0
//  3 | x y z 1
func (m *Matrix) InitTranslation(p Vector) {
	// row \  / col
	m[0][0] = 1
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[1][0] = 0
	m[1][1] = 1
	m[1][2] = 0
	m[1][3] = 0

	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = 1
	m[2][3] = 0

	m[3][0] = p.X
	m[3][1] = p.Y
	m[3][2] = p.Z
	m[3][3] = 1
}

// InitRotation Row-Major Ordering
func (m *Matrix) InitRotation(p Vector) {
	rx := Matrix{}
	ry := Matrix{}
	rz := Matrix{}

	x := DegToRad(p.X)
	y := DegToRad(p.Y)
	z := DegToRad(p.Z)

	// Rotate about the Z
	rz[0][0] = math.Cos(z)
	rz[0][1] = -math.Sin(z)
	rz[0][2] = 0
	rz[0][3] = 0
	rz[1][0] = math.Sin(z)
	rz[1][1] = math.Cos(z)
	rz[1][2] = 0
	rz[1][3] = 0
	rz[2][0] = 0
	rz[2][1] = 0
	rz[2][2] = 1
	rz[2][3] = 0
	rz[3][0] = 0
	rz[3][1] = 0
	rz[3][2] = 0
	rz[3][3] = 1

	// Rotate about the X
	rx[0][0] = 1
	rx[0][1] = 0
	rx[0][2] = 0
	rx[0][3] = 0
	rx[1][0] = 0
	rx[1][1] = math.Cos(x)
	rx[1][2] = -math.Sin(x)
	rx[1][3] = 0
	rx[2][0] = 0
	rx[2][1] = math.Sin(x)
	rx[2][2] = math.Cos(x)
	rx[2][3] = 0
	rx[3][0] = 0
	rx[3][1] = 0
	rx[3][2] = 0
	rx[3][3] = 1

	// Rotate about the Y
	ry[0][0] = math.Cos(y)
	ry[0][1] = 0
	ry[0][2] = -math.Sin(y)
	ry[0][3] = 0
	ry[1][0] = 0
	ry[1][1] = 1
	ry[1][2] = 0
	ry[1][3] = 0
	ry[2][0] = math.Sin(y)
	ry[2][1] = 0
	ry[2][2] = math.Cos(y)
	ry[2][3] = 0
	ry[3][0] = 0
	ry[3][1] = 0
	ry[3][2] = 0
	ry[3][3] = 1

	mul := Matrix{}
	ry.Mul(rx, &mul)
	rz.Mul(mul, &mul)
	// mul := rz.Mul(ry.Mul(rx))

	m[0] = mul[0]
	m[1] = mul[1]
	m[2] = mul[2]
	m[3] = mul[3]
}

// InitScale initialize with scale
//
//      0 1 2 3
//    ---------
//  0 | x 0 0 0
//  1 | 0 y 0 0
//  2 | 0 0 z 0
//  3 | 0 0 0 1
func (m *Matrix) InitScale(p Vector) {
	m[0][0] = p.X
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	m[1][0] = 0
	m[1][1] = p.Y
	m[1][2] = 0
	m[1][3] = 0

	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = p.Z
	m[2][3] = 0

	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = 0
	m[3][3] = 1
}

// InitPerspective init the matrix to a perspective projection matrix
func (m *Matrix) InitPerspective(o PerspectiveOptions) {
	tanHalfFOV := math.Tan(math.Pi*0.5 - 0.5*o.Fov)
	zRange := 1 / (o.ZNear - o.ZFar)

	// row 0
	m[0][0] = 1 / (tanHalfFOV * o.AspectRatio)
	m[0][1] = 0
	m[0][2] = 0
	m[0][3] = 0

	// row 1
	m[1][0] = 0
	m[1][1] = 1 / tanHalfFOV
	m[1][2] = 0
	m[1][3] = 0

	// row 2
	m[2][0] = 0
	m[2][1] = 0
	m[2][2] = (o.ZFar + o.ZNear) * zRange
	m[2][3] = -1

	// row 3
	m[3][0] = 0
	m[3][1] = 0
	m[3][2] = o.ZNear * o.ZFar * zRange * 2
	m[3][3] = 0
}

// Mul multiply two matries returning a new Matrix
func (m *Matrix) Mul(r Matrix, out *Matrix) {
	// This looks nasty, but it's just an unwinding of the below
	// loop.  This needs to be really fast.
	// for (let i = 0; i < 4; i++) {
	//   for (let j = 0; j < 4; j++) {
	//     res.set(i, j,
	//       this.m[i][0] * r.get(0, j) +
	//       this.m[i][1] * r.get(1, j) +
	//       this.m[i][2] * r.get(2, j) +
	//       this.m[i][3] * r.get(3, j)
	//     );
	//   }
	// }

	out[0][0] = m[0][0]*r[0][0] +
		m[0][1]*r[1][0] +
		m[0][2]*r[2][0] +
		m[0][3]*r[3][0]
	out[0][1] = m[0][0]*r[0][1] +
		m[0][1]*r[1][1] +
		m[0][2]*r[2][1] +
		m[0][3]*r[3][1]
	out[0][2] = m[0][0]*r[0][2] +
		m[0][1]*r[1][2] +
		m[0][2]*r[2][2] +
		m[0][3]*r[3][2]
	out[0][3] = m[0][0]*r[0][3] +
		m[0][1]*r[1][3] +
		m[0][2]*r[2][3] +
		m[0][3]*r[3][3]

	out[1][0] = m[1][0]*r[0][0] +
		m[1][1]*r[1][0] +
		m[1][2]*r[2][0] +
		m[1][3]*r[3][0]
	out[1][1] = m[1][0]*r[0][1] +
		m[1][1]*r[1][1] +
		m[1][2]*r[2][1] +
		m[1][3]*r[3][1]
	out[1][2] = m[1][0]*r[0][2] +
		m[1][1]*r[1][2] +
		m[1][2]*r[2][2] +
		m[1][3]*r[3][2]
	out[1][3] = m[1][0]*r[0][3] +
		m[1][1]*r[1][3] +
		m[1][2]*r[2][3] +
		m[1][3]*r[3][3]

	out[2][0] = m[2][0]*r[0][0] +
		m[2][1]*r[1][0] +
		m[2][2]*r[2][0] +
		m[2][3]*r[3][0]
	out[2][1] = m[2][0]*r[0][1] +
		m[2][1]*r[1][1] +
		m[2][2]*r[2][1] +
		m[2][3]*r[3][1]
	out[2][2] = m[2][0]*r[0][2] +
		m[2][1]*r[1][2] +
		m[2][2]*r[2][2] +
		m[2][3]*r[3][2]
	out[2][3] = m[2][0]*r[0][3] +
		m[2][1]*r[1][3] +
		m[2][2]*r[2][3] +
		m[2][3]*r[3][3]

	out[3][0] = m[3][0]*r[0][0] +
		m[3][1]*r[1][0] +
		m[3][2]*r[2][0] +
		m[3][3]*r[3][0]
	out[3][1] = m[3][0]*r[0][1] +
		m[3][1]*r[1][1] +
		m[3][2]*r[2][1] +
		m[3][3]*r[3][1]
	out[3][2] = m[3][0]*r[0][2] +
		m[3][1]*r[1][2] +
		m[3][2]*r[2][2] +
		m[3][3]*r[3][2]
	out[3][3] = m[3][0]*r[0][3] +
		m[3][1]*r[1][3] +
		m[3][2]*r[2][3] +
		m[3][3]*r[3][3]
}

// Transform Applies the current matrix to the given vector
// Expects the matrix to be in Row-Major Ordering
//
//                    | a b c t1 |    x * a + y * d + z * g
//     [x, y, z, w] * | d e f t2 | =  x * b + y * e + z * h
//                    | g h i t3 |    x * c + y * f + z * i
//                    | j k l t4 |
func (m *Matrix) Transform(r Vector, out *Vector) {

	//        row col
	out.X = (r.X * m[0][0]) + (r.Y * m[1][0]) + (r.Z * m[2][0]) + (r.W * m[3][0])
	out.Y = (r.X * m[0][1]) + (r.Y * m[1][1]) + (r.Z * m[2][1]) + (r.W * m[3][1])
	out.Z = (r.X * m[0][2]) + (r.Y * m[1][2]) + (r.Z * m[2][2]) + (r.W * m[3][2])
	out.W = (r.X * m[0][3]) + (r.Y * m[1][3]) + (r.Z * m[2][3]) + (r.W * m[3][3])

}

// Transpose Changes the matrix from Row-Major to Column-Major
func (m *Matrix) Transpose(out *Matrix) {
	out[0] = [4]float64{m[0][0], m[1][0], m[2][0], m[3][0]}
	out[1] = [4]float64{m[0][1], m[1][1], m[2][1], m[3][1]}
	out[2] = [4]float64{m[0][2], m[1][2], m[2][2], m[3][2]}
	out[3] = [4]float64{m[0][3], m[1][3], m[2][3], m[3][3]}
}

// Inverse Because with matrices we don't divide!
// But we can multiply by an inverse, which achieves the same thing.
func (m *Matrix) Inverse(out *Matrix) {
	m00 := m[0][0]
	m01 := m[0][1]
	m02 := m[0][2]
	m03 := m[0][3]

	m10 := m[1][0]
	m11 := m[1][1]
	m12 := m[1][2]
	m13 := m[1][3]

	m20 := m[2][0]
	m21 := m[2][1]
	m22 := m[2][2]
	m23 := m[2][3]

	m30 := m[3][0]
	m31 := m[3][1]
	m32 := m[3][2]
	m33 := m[3][3]

	tmp0 := m22 * m33
	tmp1 := m32 * m23
	tmp2 := m12 * m33
	tmp3 := m32 * m13
	tmp4 := m12 * m23
	tmp5 := m22 * m13
	tmp6 := m02 * m33
	tmp7 := m32 * m03
	tmp8 := m02 * m23
	tmp9 := m22 * m03
	tmp10 := m02 * m13
	tmp11 := m12 * m03
	tmp12 := m20 * m31
	tmp13 := m30 * m21
	tmp14 := m10 * m31
	tmp15 := m30 * m11
	tmp16 := m10 * m21
	tmp17 := m20 * m11
	tmp18 := m00 * m31
	tmp19 := m30 * m01
	tmp20 := m00 * m21
	tmp21 := m20 * m01
	tmp22 := m00 * m11
	tmp23 := m10 * m01

	t0 := (tmp0*m11 + tmp3*m21 + tmp4*m31) -
		(tmp1*m11 + tmp2*m21 + tmp5*m31)
	t1 := (tmp1*m01 + tmp6*m21 + tmp9*m31) -
		(tmp0*m01 + tmp7*m21 + tmp8*m31)
	t2 := (tmp2*m01 + tmp7*m11 + tmp10*m31) -
		(tmp3*m01 + tmp6*m11 + tmp11*m31)
	t3 := (tmp5*m01 + tmp8*m11 + tmp11*m21) -
		(tmp4*m01 + tmp9*m11 + tmp10*m21)

	d := 1.0 / (m00*t0 + m10*t1 + m20*t2 + m30*t3)

	out[0][0] = d * t0
	out[0][1] = d * t1
	out[0][2] = d * t2
	out[0][3] = d * t3

	out[1][0] = d * ((tmp1*m10 + tmp2*m20 + tmp5*m30) -
		(tmp0*m10 + tmp3*m20 + tmp4*m30))
	out[1][1] = d * ((tmp0*m00 + tmp7*m20 + tmp8*m30) -
		(tmp1*m00 + tmp6*m20 + tmp9*m30))
	out[1][2] = d * ((tmp3*m00 + tmp6*m10 + tmp11*m30) -
		(tmp2*m00 + tmp7*m10 + tmp10*m30))
	out[1][3] = d * ((tmp4*m00 + tmp9*m10 + tmp10*m20) -
		(tmp5*m00 + tmp8*m10 + tmp11*m20))

	out[2][0] = d * ((tmp12*m13 + tmp15*m23 + tmp16*m33) -
		(tmp13*m13 + tmp14*m23 + tmp17*m33))
	out[2][1] = d * ((tmp13*m03 + tmp18*m23 + tmp21*m33) -
		(tmp12*m03 + tmp19*m23 + tmp20*m33))
	out[2][2] = d * ((tmp14*m03 + tmp19*m13 + tmp22*m33) -
		(tmp15*m03 + tmp18*m13 + tmp23*m33))
	out[2][3] = d * ((tmp17*m03 + tmp20*m13 + tmp23*m23) -
		(tmp16*m03 + tmp21*m13 + tmp22*m23))

	out[3][0] = d * ((tmp14*m22 + tmp17*m32 + tmp13*m12) -
		(tmp16*m32 + tmp12*m12 + tmp15*m22))
	out[3][1] = d * ((tmp20*m32 + tmp12*m02 + tmp19*m22) -
		(tmp18*m22 + tmp21*m32 + tmp13*m02))
	out[3][2] = d * ((tmp18*m12 + tmp23*m32 + tmp15*m02) -
		(tmp22*m32 + tmp14*m02 + tmp19*m12))
	out[3][3] = d * ((tmp22*m22 + tmp16*m02 + tmp21*m12) -
		(tmp20*m12 + tmp23*m22 + tmp17*m02))
}

// Clone duplicate the matrix
func (m *Matrix) Clone(out *Matrix) {
	out[0] = m[0]
	out[1] = m[1]
	out[2] = m[2]
	out[3] = m[3]
}

// AsArray returns the matrix as a flat float32 array
// it's 32 bit because opengl will need that. This may not
// be the right place for this
func (m *Matrix) AsArray() [16]float32 {
	return [16]float32{
		float32(m[0][0]), float32(m[0][1]), float32(m[0][2]), float32(m[0][3]),
		float32(m[1][0]), float32(m[1][1]), float32(m[1][2]), float32(m[1][3]),
		float32(m[2][0]), float32(m[2][1]), float32(m[2][2]), float32(m[2][3]),
		float32(m[3][0]), float32(m[3][1]), float32(m[3][2]), float32(m[3][3]),
	}
}
