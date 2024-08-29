package blur

import (
	"errors"
	"gocv/num"
	"gocv/types"
	"sync"
)

func MedianBlur(img_data types.ColorImage, kernel_size int) (types.ColorImage, error) {
	h, w, c := num.Shape(img_data)

	arr := num.CreateArray3D(h, w, c)
	if kernel_size%2 == 0 || kernel_size < 1 {
		return nil, errors.New("kernel size cannot be even, 0 or negative number")
	}
	kernel_radius := kernel_size / 2

	var wg sync.WaitGroup

	for y := range img_data {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()

			for x := range img_data[0] {
				var r, g, b []types.ImageType
				var px, py int

				for i := 0; i < kernel_size; i++ {
					for j := 0; j < kernel_size; j++ {
						px = x + j - kernel_radius
						py = y + i - kernel_radius

						if px >= w || px < 0 || py >= h || py < 0 {
							continue
						}
						r = append(r, img_data[py][px][0])
						g = append(g, img_data[py][px][1])
						b = append(b, img_data[py][px][2])

					}
				}
				arr[y][x][0] = num.Median(r)
				arr[y][x][1] = num.Median(g)
				arr[y][x][2] = num.Median(b)
			}
		}(y)
	}
	wg.Wait()
	return arr, nil
}
