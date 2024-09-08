package cvt_color

import (
	"sync"
	"gocv/num"
	"gocv/types"
)
func RGBToGray(arr types.ImageArray) types.ImageArray {
	h, w, _ := num.Shape(arr)
	gray := num.CreateArray3D(h, w, 1)

	var wg sync.WaitGroup
	for y := range arr{
		wg.Add(1)
		go func(y int){
			defer wg.Done()
			for x := range arr[0]{
				gray[y][x][0] = types.ImageType(0.299 * float64(arr[y][x][0]) + 0.587  * float64(arr[y][x][1]) + 0.114  * float64(arr[y][x][2]))
			}
		}(y)
		
	}
	return gray
}