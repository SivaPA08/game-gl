package main

import (
	"game-gl/renderer"
	"game-gl/texture"
	"game-gl/types"

	"github.com/go-gl/gl/v4.4-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func main() {
	window, err := renderer.GLFW_Window(800, 600)
	if err != nil {
		panic(err)
	}

	defer glfw.Terminate()
	defer window.Destroy()

	shaderProgram, err := renderer.Shader()
	if err != nil {
		panic(err)
	}
	gl.UseProgram(shaderProgram)

	// Enable VSync
	glfw.SwapInterval(1)

	tex, err := texture.Get_Frames_From_Sheet("assets/player.png", 6, 10, []types.Index{{0, 0}, {1, 0}, {2, 0}, {3, 0}})
	if err != nil {
		panic(err)
	}

	index := 0
	MOD := len(tex)

	lastTime := glfw.GetTime()
	frameTime := 0.0
	frameDuration := 0.15 // 150ms per frame

	for !window.ShouldClose() {
		// Use a nice dark grey background instead of bright red
		gl.ClearColor(0.12, 0.12, 0.14, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		currentTime := glfw.GetTime()
		deltaTime := currentTime - lastTime
		lastTime = currentTime

		frameTime += deltaTime
		if frameTime >= frameDuration {
			index = (index + 1) % MOD
			frameTime = 0.0
		}

		// Draw the current sprite frame at (300, 200) with size 200x200
		err = renderer.Draw_Texture(tex[index], 300, 200, 200, 200)
		if err != nil {
			panic(err)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

