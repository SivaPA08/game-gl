package renderer

import (
	"errors"
	"github.com/go-gl/gl/v4.4-core/gl"
)

const vertex_shader = `
#version 410

layout (location = 0) in vec2 position;
layout (location = 1) in vec2 texCoord;

out vec2 TexCoord;

void main()
{
    gl_Position = vec4(position, 0.0, 1.0);
    TexCoord = texCoord;
}
`

const fragment_shader = `
#version 410

in vec2 TexCoord;

out vec4 color;

uniform sampler2D texture1;

void main()
{
    color = texture(texture1, TexCoord);
}
`

func Shader() (uint32, error) {

	var status int32

	vs := gl.CreateShader(gl.VERTEX_SHADER)
	vsSource, free := gl.Strs(vertex_shader + "\x00")
	defer free()
	gl.ShaderSource(vs, 1, vsSource, nil)
	gl.CompileShader(vs)
	gl.GetShaderiv(vs, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		return 0, errors.New("Vertex shader failed to compile in renderer/compile_link_shader.go")
	}

	fs := gl.CreateShader(gl.FRAGMENT_SHADER)
	fsSource, free := gl.Strs(fragment_shader + "\x00")
	defer free()
	gl.ShaderSource(fs, 1, fsSource, nil)
	gl.CompileShader(fs)
	gl.GetShaderiv(fs, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		return 0, errors.New("Fragment_shader failed to compile in renderer/compile_link_shader.go")
	}

	//creating program
	program := gl.CreateProgram()
	gl.AttachShader(program, vs)
	gl.AttachShader(program, fs)
	gl.LinkProgram(program)

	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		return 0, errors.New("Program failed to link in renderer/compile_link_shader.go")
	}
	gl.DeleteShader(vs)
	gl.DeleteShader(fs)

	return program, nil
}
