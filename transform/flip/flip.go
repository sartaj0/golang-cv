package flip

import (
	"gocv/num"
	"gocv/types"
	"sync"
)

func FlipImage(img_data types.ImageArray, horizontal_flip bool, vertical_flip bool) types.ImageArray {
	h, w, c := num.Shape(img_data)

	arr := num.CreateArray3D(h, w, c)

	var wg sync.WaitGroup
	for y := range img_data{
		wg.Add(1)

		go func(y int){
			defer wg.Done()
			for x := range img_data[0]{
				if horizontal_flip && vertical_flip{
					arr[y][x] = img_data[h-1-y][w-1-x]
				}else if horizontal_flip{
					arr[y][x] = img_data[y][w-1-x]
				}else if vertical_flip{
					arr[y][x] = img_data[h-1-y][x]
				}else{
					arr[y][x] = img_data[y][x]
				}
			}
		}(y)
	}
	wg.Wait()

	return arr
}