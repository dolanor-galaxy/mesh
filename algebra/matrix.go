package algebra

import (
	"math"
)

// Matrix 4x4 matrix Row-Major Ordering
type Matrix [4][4]float64

// Identity create an identity matrix
//     0 1 2 3
//   ---------
//  0 | 1 0 0 0
//  1 | 0 1 0 0
//  2 | 0 0 1 0
//  3 | 0 0 0 1
func (m *Matrix) Identity() *Matrix {
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

	return m
}

// InitTranslation initialize the matrix with a translation point
//      0 1 2 3
//    ---------
//  0 | 1 0 0 0
//  1 | 0 1 0 0
//  2 | 0 0 1 0
//  3 | x y z 1
func (m *Matrix) InitTranslation(p *Vector) *Matrix {
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

	return m
}

// InitRotation Row-Major Ordering
func (m *Matrix) InitRotation(p *Vector) *Matrix {
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

	mul := rz.Mul(ry.Mul(&rx))

	m[0] = mul[0]
	m[1] = mul[1]
	m[2] = mul[2]
	m[3] = mul[3]

	return m
}

// InitScale initialize with scale
//      0 1 2 3
//    ---------
//  0 | x 0 0 0
//  1 | 0 y 0 0
//  2 | 0 0 z 0
//  3 | 0 0 0 1
func (m *Matrix) InitScale(p *Vector) *Matrix {
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

	return m
}

// Mul multiply two matries
func (m *Matrix) Mul(r *Matrix) *Matrix {
	res := Matrix{}

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

	res[0][0] = m[0][0]*r[0][0] +
		m[0][1]*r[1][0] +
		m[0][2]*r[2][0] +
		m[0][3]*r[3][0]
	res[0][1] = m[0][0]*r[0][1] +
		m[0][1]*r[1][1] +
		m[0][2]*r[2][1] +
		m[0][3]*r[3][1]
	res[0][2] = m[0][0]*r[0][2] +
		m[0][1]*r[1][2] +
		m[0][2]*r[2][2] +
		m[0][3]*r[3][2]
	res[0][3] = m[0][0]*r[0][3] +
		m[0][1]*r[1][3] +
		m[0][2]*r[2][3] +
		m[0][3]*r[3][3]

	res[1][0] = m[1][0]*r[0][0] +
		m[1][1]*r[1][0] +
		m[1][2]*r[2][0] +
		m[1][3]*r[3][0]
	res[1][1] = m[1][0]*r[0][1] +
		m[1][1]*r[1][1] +
		m[1][2]*r[2][1] +
		m[1][3]*r[3][1]
	res[1][2] = m[1][0]*r[0][2] +
		m[1][1]*r[1][2] +
		m[1][2]*r[2][2] +
		m[1][3]*r[3][2]
	res[1][3] = m[1][0]*r[0][3] +
		m[1][1]*r[1][3] +
		m[1][2]*r[2][3] +
		m[1][3]*r[3][3]

	res[2][0] = m[2][0]*r[0][0] +
		m[2][1]*r[1][0] +
		m[2][2]*r[2][0] +
		m[2][3]*r[3][0]
	res[2][1] = m[2][0]*r[0][1] +
		m[2][1]*r[1][1] +
		m[2][2]*r[2][1] +
		m[2][3]*r[3][1]
	res[2][2] = m[2][0]*r[0][2] +
		m[2][1]*r[1][2] +
		m[2][2]*r[2][2] +
		m[2][3]*r[3][2]
	res[2][3] = m[2][0]*r[0][3] +
		m[2][1]*r[1][3] +
		m[2][2]*r[2][3] +
		m[2][3]*r[3][3]

	res[3][0] = m[3][0]*r[0][0] +
		m[3][1]*r[1][0] +
		m[3][2]*r[2][0] +
		m[3][3]*r[3][0]
	res[3][1] = m[3][0]*r[0][1] +
		m[3][1]*r[1][1] +
		m[3][2]*r[2][1] +
		m[3][3]*r[3][1]
	res[3][2] = m[3][0]*r[0][2] +
		m[3][1]*r[1][2] +
		m[3][2]*r[2][2] +
		m[3][3]*r[3][2]
	res[3][3] = m[3][0]*r[0][3] +
		m[3][1]*r[1][3] +
		m[3][2]*r[2][3] +
		m[3][3]*r[3][3]

	return &res
}

// Transpose Changes the matrix from Row-Major to Column-Major
func (m *Matrix) Transpose() *Matrix {
	nm := Matrix{
		{m[0][0], m[1][0], m[2][0], m[3][0]},
		{m[0][1], m[1][1], m[2][1], m[3][1]},
		{m[0][2], m[1][2], m[2][2], m[3][2]},
		{m[0][3], m[1][3], m[2][3], m[3][3]},
	}

	m[0] = nm[0]
	m[1] = nm[1]
	m[2] = nm[2]
	m[3] = nm[3]

	return &nm
}

// AsArray return the matrix as a flat array
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
