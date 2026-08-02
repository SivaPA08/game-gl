package renderer

import (
	"errors"

	"github.com/go-gl/gl/v4.4-core/gl"
)

func Draw_Texture(textureId uint32, x int32, y int32, width int, hight int) error {
	return Draw_Texture_Flipped(textureId, x, y, width, hight, false)
}

func Draw_Texture_Flipped(textureId uint32, x int32, y int32, width int, hight int, flipX bool) error {
	if quadVAO == 0 {
		CreateQuad()
	}

	program, err := Shader()
	if err != nil {
		return errors.New("Draw_Texture failed to get shader from renderer/draw_texture.go")
	}
	gl.UseProgram(program)

	// get current viewport size to map screen coordinates to normalized device coordinates (NDC which i will use below as notation)
	var viewport [4]int32
	gl.GetIntegerv(gl.VIEWPORT, &viewport[0])
	winWidth := float32(viewport[2])
	winHeight := float32(viewport[3])

	if winWidth <= 0 || winHeight <= 0 {
		return errors.New("invalid viewport dimensions")
	}

	// NDC width and height (total screen range is 2.0: from -1.0 to 1.0)
	ndcWidth := 2.0 * float32(width) / winWidth
	ndcHeight := 2.0 * float32(hight) / winHeight

	if flipX {
		ndcWidth = -ndcWidth
	}

	// map top-left (x, y) coordinates to center of the quad in screen space, then to NDC 
	centerX := float32(x) + float32(width)/2.0
	centerY := float32(y) + float32(hight)/2.0

	translationX := (2.0 * centerX / winWidth) - 1.0
	translationY := 1.0 - (2.0 * centerY / winHeight)

	// set uniforms
	translationLoc := gl.GetUniformLocation(program, gl.Str("translation\x00"))
	gl.Uniform2f(translationLoc, translationX, translationY)

	scaleLoc := gl.GetUniformLocation(program, gl.Str("scale\x00"))
	gl.Uniform2f(scaleLoc, ndcWidth, ndcHeight)

	// bind texture
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, textureId)
	texLoc := gl.GetUniformLocation(program, gl.Str("texture1\x00"))
	gl.Uniform1i(texLoc, 0)

	// bind VAO and draw
	gl.BindVertexArray(quadVAO)
	gl.DrawElements(gl.TRIANGLES, 6, gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.BindVertexArray(0)

	return nil
}
