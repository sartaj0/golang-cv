package out 

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"gocv/types"
)

/*
func processArray(arr interface{}) {
    switch v := arr.(type) {
    case [][][]uint8:
        fmt.Println("Received a 3D array")
        // Process 3D array
    case [][]uint8:
        fmt.Println("Received a 2D array")
        // Process 2D array
    default:
        fmt.Println("Unsupported type")
    }
}
*/

func ImWrite(filename string, image_data types.ColorImage){
	height := len(image_data)
	width := len(image_data[0])

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b := image_data[y][x][0], image_data[y][x][1], image_data[y][x][2]
			img.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), 255})
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


func ImWriteGray(filename string, image_data types.GrayImage){
	height := len(image_data)
	width := len(image_data[0])

	img := image.NewGray(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.Gray{uint8(image_data[y][x])})
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