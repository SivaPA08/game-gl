package renderer

import "github.com/go-gl/gl/v4.4-core/gl"

func Draw_Texture(textureId uint32, x int32, y int32, width int, hight int) {
	gl.BindTexture(gl.TEXTURE_2D, textureId)
}
