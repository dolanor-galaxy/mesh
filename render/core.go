package render

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	gl "github.com/chsc/gogl/gl21"
	"github.com/therohans/mesh/geometry"
)

// Program wrapper object for an OpenGL Program with Attribute Locations
type Program struct {
	Program     gl.Uint
	PosLoc      gl.Uint
	ColorLoc    gl.Uint
	TexCoordLoc gl.Uint
	NormalLoc   gl.Uint
	TangentLoc  gl.Uint
	UniScale    gl.Int
}

// ReadVertexShader read a vertex shader from disk
func ReadVertexShader(path string) string {
	return readShader("vertex", path)
}

// ReadFragmentShader read a fragment shader from disk
func ReadFragmentShader(path string) string {
	return readShader("fragment", path)
}

func readShader(shaderType string, path string) string {
	fullPath, err := filepath.Abs(
		filepath.Join("assets", "shaders", shaderType, path))
	if err != nil {
		panic(err)
	}

	b, err := ioutil.ReadFile(fullPath)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// CreateProgram creates an opengl program (OpenGL 2.1)
func CreateProgram(vertexSource string, fragmentSource string) (gl.Uint, error) {
	// Vertex shader
	vs := gl.CreateShader(gl.VERTEX_SHADER)
	vsSource := gl.GLStringArray(vertexSource)
	defer gl.GLStringArrayFree(vsSource)
	gl.ShaderSource(vs, 1, &vsSource[0], nil)

	gl.CompileShader(vs)

	status, err := compileStatus(vs)
	if err != nil {
		return status, err
	}

	// Fragment shader
	fs := gl.CreateShader(gl.FRAGMENT_SHADER)
	fsSource := gl.GLStringArray(fragmentSource)
	defer gl.GLStringArrayFree(fsSource)
	gl.ShaderSource(fs, 1, &fsSource[0], nil)
	gl.CompileShader(fs)

	status, err = compileStatus(fs)
	if err != nil {
		return status, err
	}

	// create program
	program := gl.CreateProgram()
	gl.AttachShader(program, vs)
	gl.AttachShader(program, fs)

	gl.LinkProgram(program)
	var linkstatus gl.Int
	gl.GetProgramiv(program, gl.LINK_STATUS, &linkstatus)
	if linkstatus == gl.FALSE {
		return gl.FALSE, errors.New("Program link failed")
	}

	return program, nil
}

// Checks a shader compile status to see if it errored and
// also tries to get out the log as to why it died
func compileStatus(shader gl.Uint) (gl.Uint, error) {
	var status gl.Int
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength gl.Int
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		chary := gl.GLStringArray(log)
		defer gl.GLStringArrayFree(chary)
		gl.GetShaderInfoLog(shader, gl.Sizei(logLength), nil, chary[0])
		logOut := gl.GoString(chary[0])

		return gl.FALSE, fmt.Errorf("failed to compile %v", logOut)
	}

	return gl.TRUE, nil
}

// UseProgram uses a program
func UseProgram() Program {
	program, err := CreateProgram(
		ReadVertexShader("Demo.glsl"),
		ReadFragmentShader("Demo.glsl"))
	if err != nil {
		panic(err)
	}

	gl.UseProgram(program)

	// Uniforms
	unistring := gl.GLString("scaleMove")
	uniScale := gl.GetUniformLocation(program, unistring)
	if gl.GetError() != gl.NO_ERROR {
		log.Printf("GetUniformLocation not found.")
	}

	// Attributes
	posLoc := gl.GetAttribLocation(program, gl.GLString("Pos"))
	if posLoc == -1 {
		log.Printf("Position attribute not found.")
	}
	colorLoc := gl.GetAttribLocation(program, gl.GLString("Color"))
	if colorLoc == -1 {
		log.Printf("Color attribute not found.")
	}
	texCoordLoc := gl.GetAttribLocation(program, gl.GLString("TexCoord"))
	if texCoordLoc == -1 {
		log.Printf("TexCoord attribute not found.")
	}
	normalLoc := gl.GetAttribLocation(program, gl.GLString("Normal"))
	if normalLoc == -1 {
		log.Printf("Normal attribute not found.")
	}
	tangentLoc := gl.GetAttribLocation(program, gl.GLString("Tangent"))
	if tangentLoc == -1 {
		log.Printf("Tangent attribute not found.")
	}

	gl.EnableVertexAttribArray(gl.Uint(posLoc))
	gl.EnableVertexAttribArray(gl.Uint(colorLoc))
	gl.EnableVertexAttribArray(gl.Uint(texCoordLoc))
	gl.EnableVertexAttribArray(gl.Uint(normalLoc))
	gl.EnableVertexAttribArray(gl.Uint(tangentLoc))
	if gl.GetError() != gl.NO_ERROR {
		log.Printf("EnableVertexAttribArray failed.")
	}

	bpe := gl.Sizei(geometry.VertexSize * 4)
	gl.VertexAttribPointer(gl.Uint(posLoc), 3, gl.FLOAT, gl.FALSE, bpe, gl.Offset(nil, uintptr(0)))
	gl.VertexAttribPointer(gl.Uint(colorLoc), 3, gl.FLOAT, gl.FALSE, bpe, gl.Offset(nil, uintptr(3*4)))
	gl.VertexAttribPointer(gl.Uint(texCoordLoc), 2, gl.FLOAT, gl.TRUE, bpe, gl.Offset(nil, uintptr(6*4)))
	gl.VertexAttribPointer(gl.Uint(normalLoc), 3, gl.FLOAT, gl.TRUE, bpe, gl.Offset(nil, uintptr(8*4)))
	gl.VertexAttribPointer(gl.Uint(tangentLoc), 3, gl.FLOAT, gl.TRUE, bpe, gl.Offset(nil, uintptr(11*4)))
	if gl.GetError() != gl.NO_ERROR {
		log.Printf("VertexAttribPointer failed.")
	}

	return Program{
		Program:     program,
		PosLoc:      gl.Uint(posLoc),
		ColorLoc:    gl.Uint(colorLoc),
		TexCoordLoc: gl.Uint(texCoordLoc),
		NormalLoc:   gl.Uint(normalLoc),
		TangentLoc:  gl.Uint(tangentLoc),
		UniScale:    uniScale,
	}
}
