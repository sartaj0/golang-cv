package out

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"gocv/num"
	"gocv/types"
)



func ImWrite(filename string, image_data types.ImageArray){
	height, width, channels := num.Shape(image_data)

	var img image.Image

	if channels == 3{
		img = image.NewRGBA(image.Rect(0, 0, width, height))
	}else if channels == 1{
		img = image.NewGray(image.Rect(0, 0, width, height))
	}
	
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if channels == 3{
				r, g, b := image_data[y][x][0], image_data[y][x][1], image_data[y][x][2]
				img.(*image.RGBA).Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
			}else if channels == 1 {
				img.(*image.Gray).Set(x, y, color.Gray{uint8(image_data[y][x][0])})
			}
			
		}
	}
	f, err := os.Create(filename)
	if err != nil {
        fmt.Println("You got error bro")
        return
    }
    defer f.Close()
    png.Encode(f, img)
}