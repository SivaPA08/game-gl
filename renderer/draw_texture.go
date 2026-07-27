package renderer

import (
	"errors"

	"github.com/go-gl/gl/v4.4-core/gl"
)

func Draw_Texture(textureId uint32, x int32, y int32, width int, hight int) error {
	program, err := Shader()
	if err != nil {
		return errors.New("Draw_Texture failed to get shader from renderer/draw_texture.go")
	}
	gl.UseProgram(program)
	return nil
}
