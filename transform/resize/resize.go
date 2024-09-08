package resize

import (
	"fmt"
	"gocv/num"
	"gocv/types"
	"math"
	"sync"
)

func nearest_neighbour_resize(img_data types.ImageArray, height int, width int) types.ImageArray {

	h, w, _ := num.Shape(img_data)

	arr := num.CreateArray3D(height, width, len(img_data[0][0]))
	alpha_w := float64(width-1) / float64(w-1)
	alpha_h := float64(height-1) / float64(h-1)

	var new_x, new_y int

	for y := range arr {
		for x := range arr[0] {
			new_x = int(math.Round(float64(x) / alpha_w))
			new_y = int(math.Round(float64(y) / alpha_h))
			arr[y][x] = img_data[new_y][new_x]
		}
	}
	return arr
}

func nearest_neighbour_resize_optimized(img_data types.ImageArray, height int, width int) types.ImageArray {
	h, w, _ := num.Shape(img_data)
	var wg sync.WaitGroup
	arr := num.CreateArray3D(height, width, len(img_data[0][0]))

	alpha_w := (w - 1) * 1000 / (width - 1)
	alpha_h := (h - 1) * 1000 / (height - 1)

	for y := range arr {

		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := range arr[0] {
				new_x := x * alpha_w / 1000
				new_y := y * alpha_h / 1000
				arr[y][x] = img_data[new_y][new_x]
			}
		}(y)
	}
	wg.Wait()
	return arr
}

func Resize(img_data types.ImageArray, height int, width int) types.ImageArray {

	h, w, _ := num.Shape(img_data)

	if height == 0 && width == 0 {
		fmt.Println("Image can't be resize for both unknown length")
		return img_data
	} else if height == 0 {
		height = int(math.Round((float64(width) * float64(h)) / (float64(w))))
	} else if width == 0 {
		width = int(math.Round((float64(height) * float64(w)) / (float64(h))))
	}

	arr := nearest_neighbour_resize_optimized(img_data, height, width)
	return arr
}
