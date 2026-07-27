package texture

import (
	"errors"
	"game-gl/types"
	"image"
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

	for _, idx := range indexes {
		x0 := int(idx.X) * frame_width
		y0 := int(idx.Y) * frame_height
		rect := image.Rect(x0, y0, x0+frame_width, y0+frame_height)
		frame := img.SubImage(rect).(*image.RGBA)
		frames = append(frames, frame)
	}

	textures := make([]uint32, 0, len(frames))

	for i, frame := range frames {
		tex, err := Load_Texture(frame)
		if err != nil {
			return nil, errors.New("Get_Frames_From_Sheet failed to load texture from texture/load_texture.go")
		}
		textures[i] = tex
	}
	return textures, nil
}
