package read

import (
	"errors"
	"fmt"
	"gocv/types"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"gocv/num"
)

func read_image_from_path(img_path string) (image.Image, error) {
	f, err := os.Open(img_path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer f.Close()

	img, _, err := image.Decode(f)

	if err != nil {
		fmt.Println("Decoding error:", err.Error())
		return nil, err
	}

	return img, nil
}

func rgb_to_pixel(r uint32, g uint32, b uint32, a uint32) []types.ImageType {
	alpha := float64(a>>8) / float64(255)
	adjusted_r := types.ImageType(float64(r>>8) * alpha)
	adjusted_g := types.ImageType(float64(g>>8) * alpha)
	adjusted_b := types.ImageType(float64(b>>8) * alpha)
	return []types.ImageType{adjusted_r, adjusted_g, adjusted_b}
}

func ReadImage(img_path string) (types.ImageArray, error) {

	img, err := read_image_from_path(img_path)
	if err != nil {
		return nil, errors.New("You got error bro")
	}
	size := img.Bounds().Size()
	H, W := size.Y, size.X
	var pixel = []types.ImageType{0, 0, 0}

	pixels := num.CreateArray3D(H, W, 3)

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			pixel = rgb_to_pixel(img.At(j, i).RGBA())
			pixels[i][j] = pixel
		}
	}
	return pixels, nil
}
