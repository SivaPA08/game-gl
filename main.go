package main

import (
	"game-gl/animation"
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

	//animation frames
	front_idle, err := texture.Get_Frames_From_Sheet("assets/player.png", 6, 10, []types.Index{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}})
	if err != nil {
		panic(err)
	}
	// side_idle, err := texture.Get_Frames_From_Sheet("assets/player.png", 6, 10, []types.Index{{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1}})
	// if err != nil {
	// 	panic(err)
	// }
	// back_idle,err:=texture.Get_Frames_From_Sheet("assets/player.png",6,10,[]types.Index{{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}, {5, 2}})
	// if err!=nil {
	// 	panic(err)
	// }
	front_walk, err := texture.Get_Frames_From_Sheet("assets/player.png", 6, 10, []types.Index{{0, 3}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3}})
	if err != nil {
		panic(err)
	}
	//animation frames ends

	for !window.ShouldClose() {
		gl.ClearColor(0.12, 0.12, 0.14, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		if window.GetKey(glfw.KeyLeft) == glfw.Press {
			err = animation.Animate(front_walk, 3, glfw.Press)
		} else {
			err = animation.Animate(front_idle, 3, glfw.Release)
		}
		if err != nil {
			panic(err)
		}
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
