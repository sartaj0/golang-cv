package blur

import (
	"errors"
	"gocv/num"
	"gocv/types"
	"sync"
)

func AverageBlur(img_data types.ColorImage, kernel_size int) (types.ColorImage, error ){
	h, w, c := num.Shape(img_data)

	arr := make(types.ColorImage, h)
	if kernel_size % 2 == 0{
		return nil, errors.New("kernel size cannot be even number")
	}
	kernel_radius := kernel_size / 2
	
	var wg sync.WaitGroup

	for y := range img_data {
		wg.Add(1)
		go func (y int){
			defer wg.Done()

			arr[y] = make([][]types.ImageType, w)

			for x := range img_data[0]{
				arr[y][x] = make([]types.ImageType, c)
				var r, g, b, num_cell uint
				var px, py int
				for i := 0; i < kernel_size; i++ {
					for j := 0; j < kernel_size; j++ {
						px = x + j - kernel_radius
						py = y + i - kernel_radius

						if px >= w || px < 0 || py >= h || py < 0 {
							continue
						}

						r += uint(img_data[py][px][0])
						g += uint(img_data[py][px][1])
						b += uint(img_data[py][px][2])

						num_cell += 1
					}
				}
				arr[y][x][0] = types.ImageType(r / num_cell)
				arr[y][x][1] = types.ImageType(g / num_cell)
				arr[y][x][2] = types.ImageType(b / num_cell)
			}
		}(y)
	}
	wg.Wait()
	return arr, nil
}