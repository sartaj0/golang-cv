package resize 

import (
	"fmt"
	"math"
	"gocv/num"
	"gocv/types"
)

func nearest_neighbour_resize(img_data types.ColorImage, height int, width int)types.ColorImage{

	h, w, _ := num.Shape(img_data)

	arr := num.CreateArray(height, width, len(img_data[0][0]))
	fmt.Println(height, width, h, w)
	alpha_w :=  float64(width - 1) / float64(w - 1) 
	alpha_h :=  float64(height - 1) / float64(h - 1)

	var new_x, new_y int

	for y := range arr{
		for x := range arr[0]{
			new_x = int(math.Round(float64(x ) / alpha_w))
			new_y = int(math.Round(float64(y) / alpha_h))
			arr[y][x] = img_data[new_y][new_x]
		}
	}
	return arr
}

func Resize(img_data types.ColorImage, height int, width int) types.ColorImage {
	
	h, w, _ := num.Shape(img_data)


	if height == 0 && width == 0 {
		fmt.Println("Image can't be resize for both unknown length")
		return img_data
	}else if height == 0{
		height = int(math.Round((float64(width) * float64(h)) / (float64(w))))
	}else if width == 0{
		width = int(math.Round((float64(height) * float64(w)) / (float64(h))))
	}

	arr := nearest_neighbour_resize(img_data, height, width)
	return arr
}