package renderer

import (
	"errors"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func GLFW_Window(width int32, height int32) (*glfw.Window, error) {
	if glfw.Init() != nil {
		return nil, errors.New("glfw init failed")
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 4)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	window, err := glfw.CreateWindow(int(width), int(height), "game", nil, nil)
	window.MakeContextCurrent()
	if err != nil {
		return nil, errors.New("glfw create window failed in renderer/glfw_window.go")
	}
	if GLInit(width, height) != nil {
		return nil, errors.New("GLInit failed to initilize from renderer/glfw_window.go")
	}
	return window, nil
}
