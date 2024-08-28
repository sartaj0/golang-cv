package cvt_color

import (
	"sync"
	"gocv/num"
	"gocv/types"
)
func RGBToGray(arr types.ColorImage) types.GrayImage {
	h, w, _ := num.Shape(arr)
	gray := make(types.GrayImage, h)

	var wg sync.WaitGroup
	for y := range arr{
		gray[y] = make([]types.ImageType, w)
		wg.Add(1)
		go func(y int){
			defer wg.Done()
			for x := range arr[0]{
				gray[y][x] = types.ImageType(0.299 * float64(arr[y][x][0]) + 0.587  * float64(arr[y][x][1]) + 0.114  * float64(arr[y][x][2]))
			}
		}(y)
		
	}
	return gray
}