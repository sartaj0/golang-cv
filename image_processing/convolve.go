package convolve

import (
	"sync"
	"errors"
	"gocv/num"
	"gocv/types"
)

func ConvolveImage(img_data types.ImageArray, kernel_size int, kernel [][]float64) (types.ImageArray, error) {

	height, width, channels := num.Shape(img_data)
	arr := make(types.ImageArray, height)
	if kernel_size%2 == 0 {
		return nil, errors.New("kernel size cannot be even number")
	}
	kernel_radius := kernel_size / 2

	var wg sync.WaitGroup

	for y := range img_data {
		arr[y] = make([][]types.ImageType, width)

		wg.Add(1)
		go func(y int){

			defer wg.Done()
			for x := range img_data[0] {
				arr[y][x] = make([]types.ImageType, channels)

				var r, g, b, gray float64
				var px, py int

				for i := 0; i < kernel_size; i++ {
					for j := 0; j < kernel_size; j++ {
						px = x + j - kernel_radius
						py = y + i - kernel_radius


						if px >= width || px < 0 {
							if px >= width{
								px = width - 1
							}else{
								px = 0
							}
						}

						if py >= height || py < 0 {
							if py >= height{
								py = height - 1
							}else{
								py = 0
							}
						}
						
						if channels == 3{
							r += (float64(img_data[py][px][0]) * kernel[i][j])
							g += (float64(img_data[py][px][1]) * kernel[i][j])
							b += (float64(img_data[py][px][2]) * kernel[i][j])
						}else if channels == 1{
							gray += float64(img_data[py][px][0]) * kernel[i][j]
						}
					}
				}
				if channels == 3{
					arr[y][x][0] = types.ImageType(r)
					arr[y][x][1] = types.ImageType(g)
					arr[y][x][2] = types.ImageType(b)
				}else if channels == 1{
					arr[y][x][0] = types.ImageType(gray)
				}


			}
		}(y)
	}

	wg.Wait()
	return arr, nil
}