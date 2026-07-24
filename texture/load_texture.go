package texture

import (
	"image"

	"github.com/go-gl/gl/v4.4-core/gl"
)

func Load_Texture(img *image.RGBA) (uint32, error) {
	var texture uint32
	gl.GenTextures(1, &texture) //tells to create 1 texture on opengl
	gl.BindTexture(gl.TEXTURE_2D, texture)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)

	bounds := img.Bounds()

	gl.TexImage2D(
		gl.TEXTURE_2D,
		0,
		gl.RGBA8,
		int32(bounds.Dx()),
		int32(bounds.Dy()),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		gl.Ptr(img.Pix),
	)

	gl.GenerateMipmap(gl.TEXTURE_2D)

	return texture, nil
}
