package core

import "github.com/robrohan/mesh/algebra"

// Transform position in space
type Transform struct {
	Position algebra.Vector
	Rotation algebra.Quaternion
	Scale    algebra.Vector
	Forward  algebra.Vector
	Up       algebra.Vector
	Right    algebra.Vector
}

// NewTransform create a new transform at 0,0,0 with scale 1
func NewTransform() *Transform {
	t := Transform{
		Position: algebra.Vector{},
		Rotation: algebra.Quaternion{},
		Scale:    algebra.Vector{X: 1, Y: 1, Z: 1},
		Forward:  algebra.Forward,
		Up:       algebra.Up,
		Right:    algebra.Right,
	}
	return &t
}

func (t *Transform) GetTransformation() *algebra.Matrix {
	translationMatrix := algebra.Matrix{}
	translationMatrix.InitTranslation(&t.Position)
	rotationMatrix := t.RotationMatrix(&t.Rotation)
	scaleMatrix := algebra.Matrix{}
	scaleMatrix.InitScale(&t.Scale)
	// scaleMatrix := new Matrix4().initScale(this.scale);

	// “first translate, then rotate, then scale”.
	// This first one will rotate about the world space
	// return this.getParentMatrix().mul(translationMatrix.mul(rotationMatrix.mul(scaleMatrix)));
	// Where as this one rotates at the object space
	tr := algebra.Matrix{}
	translationMatrix.Mul(*rotationMatrix, &tr)
	trs := algebra.Matrix{}
	scaleMatrix.Mul(tr, &trs)

	// return this.getParentMatrix().mul(scaleMatrix.mul(rotationMatrix.mul(translationMatrix)));
	return &trs
}

func (t *Transform) RotationMatrix(rot *algebra.Quaternion) *algebra.Matrix {
	t.Forward = algebra.Vector{
		X: 2.0 * (rot.X*rot.Z - rot.W*rot.Y),
		Y: 2.0 * (rot.Y*rot.Z + rot.W*rot.X),
		Z: 1.0 - 2.0*(rot.X*rot.X+rot.Y*rot.Y),
	}
	t.Up = algebra.Vector{
		X: 2.0 * (rot.X*rot.Y + rot.W*rot.Z),
		Y: 1.0 - 2.0*(rot.X*rot.X+rot.Z*rot.Z),
		Z: 2.0 * (rot.Y*rot.Z - rot.W*rot.X),
	}
	t.Right = algebra.Vector{
		X: 1.0 - 2.0*(rot.Y*rot.Y+rot.Z*rot.Z),
		Y: 2.0 * (rot.X*rot.Y - rot.W*rot.Z),
		Z: 2.0 * (rot.X*rot.Z + rot.W*rot.Y),
	}
	mat := algebra.Matrix{}
	mat.InitFUR(&t.Forward, &t.Up, &t.Right)
	// return new Matrix4().initRotationFUR(forward, up, right);
	return &mat
}
