package blur

import (
	"errors"
	"gocv/num"
	"gocv/types"
	"sync"
)

func GaussianBlur(img_data types.ColorImage, kernel_size int) (types.ColorImage, error) {
	h, w, c := num.Shape(img_data)

	gaussian_kernel := num.GetGaussianKernel(1, kernel_size)


	arr := make(types.ColorImage, h)
	if kernel_size%2 == 0 {
		return nil, errors.New("kernel size cannot be even number")
	}
	kernel_radius := kernel_size / 2

	var wg sync.WaitGroup

	for y := range img_data {
		arr[y] = make([][]types.ImageType, w)

		wg.Add(1)
		go func(y int){

			defer wg.Done()
			for x := range img_data[0] {
				arr[y][x] = make([]types.ImageType, c)

				var r, g, b float64
				var px, py int

				for i := 0; i < kernel_size; i++ {
					for j := 0; j < kernel_size; j++ {
						px = x + j - kernel_radius
						py = y + i - kernel_radius

						if px >= w || px < 0 {
							px = x
						}

						if py >= h || py < 0 {
							py = y
						}

						r += (float64(img_data[py][px][0]) * gaussian_kernel[i][j])
						g += (float64(img_data[py][px][1]) * gaussian_kernel[i][j])
						b += (float64(img_data[py][px][2]) * gaussian_kernel[i][j])
						
					}
				}

				arr[y][x][0] = types.ImageType(r)
				arr[y][x][1] = types.ImageType(g)
				arr[y][x][2] = types.ImageType(b)

			}
		}(y)
	}

	wg.Wait()
	return arr, nil
}
