package animation

import (
	"errors"
	"game-gl/renderer"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func Animate(sprites []uint32, fps int, key glfw.Action) error {
	lastTime := glfw.GetTime()
	frameTime := 0.0
	frameTimeDuratoin := 1.0 / float32(fps)
	index := 0
	for glfw.Press == key {
		currentTime := glfw.GetTime()
		deltaTime := currentTime - lastTime
		lastTime = currentTime
		frameTime += deltaTime
		if frameTime >= float64(frameTimeDuratoin) {
			index = (index + 1) % len(sprites)
			frameTime = 0.0
		}
		err := renderer.Draw_Texture(sprites[index], 300, 200, 200, 200)
		if err != nil {
			return errors.New("animation.go:Animate:Draw_Texture:err")

		}
	}
	return nil
}
