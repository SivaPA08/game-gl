package animation

import (
	"game-gl/renderer"
)

//it returns single frame at a time 
func DrawFrame(textureId uint32, x, y int32, width, height int, flipX bool) error {
	return renderer.Draw_Texture_Flipped(textureId, x, y, width, height, flipX)
}
