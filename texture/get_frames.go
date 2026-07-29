package texture

import (
	"errors"
	"game-gl/types"
	"image"
	"image/draw"
)

func Get_Frames_From_Sheet(path string, horizontal int32, vertical int32, indexes []types.Index) ([]uint32, error) {
	img, err := Get_Sprite(path)
	if err != nil {
		return nil, errors.New("Get_Frames_From_Sheet failed to get sprite from texture/get_sprite.go")
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	frame_width := width / int(horizontal)
	frame_height := height / int(vertical)

	frames := make([]*image.RGBA, len(indexes))

	for i, idx := range indexes {
		x0 := int(idx.X) * frame_width
		y0 := int(idx.Y) * frame_height
		rect := image.Rect(x0, y0, x0+frame_width, y0+frame_height)
		
		// Create a new independent RGBA image with bounds starting at (0,0)
		frame := image.NewRGBA(image.Rect(0, 0, frame_width, frame_height))
		// Copy the pixels from the sprite sheet sub-image region to the new frame
		draw.Draw(frame, frame.Bounds(), img, rect.Min, draw.Src)
		
		frames[i] = frame
	}

	textures := make([]uint32, len(frames))

	for i, frame := range frames {
		tex, err := Load_Texture(frame)
		if err != nil {
			return nil, errors.New("Get_Frames_From_Sheet failed to load texture from texture/load_texture.go")
		}
		textures[i] = tex
	}
	return textures, nil
}
