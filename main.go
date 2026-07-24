package main

import (
	"game-gl/renderer"

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

	for !window.ShouldClose() {
		gl.ClearColor(1.0, 0.0, 0.0, 1.0) //set red color
		gl.Clear(gl.COLOR_BUFFER_BIT)
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
