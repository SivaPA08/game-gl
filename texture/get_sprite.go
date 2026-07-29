package texture

import (
	"errors"
	"image"
	"image/draw"
	_ "image/png"
	"os"
)

func Get_Sprite(path string) (*image.RGBA, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("Sprite didnt created in animation/sprite.go")
	}
	defer file.Close()
	img, _, err := image.Decode(file) //_ is format will be filling if i want
	if err != nil {
		return nil, errors.New("Sprite didnt decoded in animation/sprite.go")
	}
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src) //copying piels from img to rgba
	return rgba, nil
}
