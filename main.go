package main

import (
	"game-gl/renderer"
	"game-gl/texture"

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

	tex, err := texture.Get_Frames_From_Sheet("assets/player.png", 4, 4, []texture.Index{{0, 0}, {1, 0}, {2, 0}, {3, 0}})
	if err != nil {
		panic(err)
	}
	index := 0
	MOD = len(tex)
	for !window.ShouldClose() {
		gl.ClearColor(1.0, 0.0, 0.0, 1.0) //set red color
		gl.Clear(gl.COLOR_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
