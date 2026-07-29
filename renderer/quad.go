package renderer

import (
	"unsafe"

	"github.com/go-gl/gl/v4.4-core/gl"
)

var quadVAO uint32
var quadVBO uint32
var quadEBO uint32

var ver = []float32{
	// x      y      z      u    v
	0.5, 0.5, 0.0, 1.0, 1.0, // top right
	0.5, -0.5, 0.0, 1.0, 0.0, // bottom right
	-0.5, -0.5, 0.0, 0.0, 0.0, // bottom left
	-0.5, 0.5, 0.0, 0.0, 1.0, // top left
}

// for ebo
var indices = []uint32{
	0, 1, 3,
	1, 2, 3,
}

func CreateQuad() {
	gl.GenVertexArrays(1, &quadVAO)
	gl.GenBuffers(1, &quadVBO)
	gl.GenBuffers(1, &quadEBO)

	//vao
	gl.BindVertexArray(quadVAO)

	//vbo
	gl.BindBuffer(gl.ARRAY_BUFFER, quadVBO)
	gl.BufferData(
		gl.ARRAY_BUFFER,
		len(ver)*int(unsafe.Sizeof(ver[0])),
		gl.Ptr(ver),
		gl.STATIC_DRAW,
	)

	//ebo
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, quadEBO)
	gl.BufferData(
		gl.ELEMENT_ARRAY_BUFFER,
		len(indices)*int(unsafe.Sizeof(indices[0])),
		gl.Ptr(indices),
		gl.STATIC_DRAW,
	)

	// Position attribute
	gl.VertexAttribPointer(
		0,
		3,
		gl.FLOAT,
		false,
		5*4,
		gl.PtrOffset(0),
	)
	gl.EnableVertexAttribArray(0)

	// Texture coordinate attribute
	gl.VertexAttribPointer(
		1,
		2,
		gl.FLOAT,
		false,
		5*4,
		gl.PtrOffset(3*4),
	)
	gl.EnableVertexAttribArray(1)

	gl.BindVertexArray(0)
}
