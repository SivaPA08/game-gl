package renderer

import (
	"errors"
	"github.com/go-gl/gl/v4.4-core/gl"
)

func GLInit(width int32, height int32) error {
	if gl.Init() != nil {
		return errors.New("gl init failed :came from renderer/gl_init.go")
	}
	gl.Viewport(0, 0, width, height)
	return nil
}
