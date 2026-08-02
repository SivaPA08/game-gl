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

	// Animation frames for all directions (Idle & Walk)
	// Row 0: Front Idle
	front_idle, err := texture.Get_Frames_From_Sheet("assets/player.png", 6, 10, []types.Index{
		{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0},
	})
	if err != nil {
		panic(err)
	}

	// Row 1: Side Idle
	side_idle, err := texture.Get_Frames_From_Sheet("assets/player.png", 6, 10, []types.Index{
		{0, 1}, {1, 1}, {2, 1}, {3, 1}, {4, 1}, {5, 1},
	})
	if err != nil {
		panic(err)
	}

	// Row 2: Back Idle
	back_idle, err := texture.Get_Frames_From_Sheet("assets/player.png", 6, 10, []types.Index{
		{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}, {5, 2},
	})
	if err != nil {
		panic(err)
	}

	// Row 3: Front Walk
	front_walk, err := texture.Get_Frames_From_Sheet("assets/player.png", 6, 10, []types.Index{
		{0, 3}, {1, 3}, {2, 3}, {3, 3}, {4, 3}, {5, 3},
	})
	if err != nil {
		panic(err)
	}

	// Row 4: Side Walk
	side_walk, err := texture.Get_Frames_From_Sheet("assets/player.png", 6, 10, []types.Index{
		{0, 4}, {1, 4}, {2, 4}, {3, 4}, {4, 4}, {5, 4},
	})
	if err != nil {
		panic(err)
	}

	// Row 5: Back Walk
	back_walk, err := texture.Get_Frames_From_Sheet("assets/player.png", 6, 10, []types.Index{
		{0, 5}, {1, 5}, {2, 5}, {3, 5}, {4, 5}, {5, 5},
	})
	if err != nil {
		panic(err)
	}

	// Player configuration
	playerX := float32(300.0)
	playerY := float32(200.0)
	playerSpeed := float32(180.0) // pixels per second
	playerWidth := 150
	playerHeight := 150

	// Animation State
	currentFrameTime := 0.0
	frameDuration := 1.0 / 8.0 // 8 FPS
	currentFrameIndex := 0

	// 0 = Down (Front), 1 = Up (Back), 2 = Right (Side), 3 = Left (Side Flipped)
	facingDirection := 0
	isMoving := false

	lastTime := glfw.GetTime()

	for !window.ShouldClose() {
		// Calculate delta time
		currentTime := glfw.GetTime()
		deltaTime := currentTime - lastTime
		lastTime = currentTime

		// Avoid extreme delta time on window drag/resize
		if deltaTime > 0.1 {
			deltaTime = 0.1
		}

		gl.ClearColor(0.12, 0.12, 0.14, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)

		// Input processing
		isMoving = false
		var dx, dy float32

		if window.GetKey(glfw.KeyLeft) == glfw.Press || window.GetKey(glfw.KeyA) == glfw.Press {
			dx = -1
			isMoving = true
			facingDirection = 3
		} else if window.GetKey(glfw.KeyRight) == glfw.Press || window.GetKey(glfw.KeyD) == glfw.Press {
			dx = 1
			isMoving = true
			facingDirection = 2
		}

		if window.GetKey(glfw.KeyUp) == glfw.Press || window.GetKey(glfw.KeyW) == glfw.Press {
			dy = -1
			isMoving = true
			// If not moving horizontally, set facingDirection to Up
			if dx == 0 {
				facingDirection = 1
			}
		} else if window.GetKey(glfw.KeyDown) == glfw.Press || window.GetKey(glfw.KeyS) == glfw.Press {
			dy = 1
			isMoving = true
			// If not moving horizontally, set facingDirection to Down
			if dx == 0 {
				facingDirection = 0
			}
		}

		// Diagonal movement speed normalization
		if dx != 0 && dy != 0 {
			dx *= 0.707107
			dy *= 0.707107
		}

		// Update position
		playerX += dx * playerSpeed * float32(deltaTime)
		playerY += dy * playerSpeed * float32(deltaTime)

		// Keep player on screen
		if playerX < -30 {
			playerX = -30
		}
		if playerX > 800 - float32(playerWidth) + 30 {
			playerX = 800 - float32(playerWidth) + 30
		}
		if playerY < -30 {
			playerY = -30
		}
		if playerY > 600 - float32(playerHeight) + 30 {
			playerY = 600 - float32(playerHeight) + 30
		}

		// Animation frame update
		currentFrameTime += deltaTime
		if currentFrameTime >= frameDuration {
			currentFrameIndex = (currentFrameIndex + 1) % 6
			currentFrameTime = 0.0
		}

		// Select current animation frames based on state
		var activeFrames []uint32
		flipX := false

		switch facingDirection {
		case 0: // Down (Front)
			if isMoving {
				activeFrames = front_walk
			} else {
				activeFrames = front_idle
			}
		case 1: // Up (Back)
			if isMoving {
				activeFrames = back_walk
			} else {
				activeFrames = back_idle
			}
		case 2: // Right (Side)
			if isMoving {
				activeFrames = side_walk
			} else {
				activeFrames = side_idle
			}
		case 3: // Left (Side Flipped)
			if isMoving {
				activeFrames = side_walk
			} else {
				activeFrames = side_idle
			}
			flipX = true
		}

		// Render current frame
		err = animation.DrawFrame(activeFrames[currentFrameIndex], int32(playerX), int32(playerY), playerWidth, playerHeight, flipX)
		if err != nil {
			panic(err)
		}

		window.SwapBuffers()
		glfw.PollEvents()
	}
}
