package edge

import (
	"gocv/num"
	"gocv/types"
	"math"
	"sync"
)

func Sobel(img_data types.GrayImage, sobelx int, sobely int) types.GrayImage {
	kernel_size := 3
	kernel_radius := kernel_size / 2

	h, w := len(img_data), len(img_data[0])
	kernelx := [][]int{
		{-1, 0, 1}, 
		{-2, 0, 2}, 
		{-1, 0, 1}}

	kernely := [][]int{
		{-1, -2, -1},
		{0,  0,  0},
		{1,  2,  1}}
	

	arr := num.CreateArray2D(h, w)

	var wg sync.WaitGroup

	for y := range img_data {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()

			for x := range img_data[0] {
				var px, py, gx, gy int

				for i := 0; i < kernel_size; i++ {
					for j := 0; j < kernel_size; j++ {
						px = x + j - kernel_radius
						py = y + i - kernel_radius

						if px >= w || px < 0 {
							if px >= w{
								px = w - 1
							}else{
								px = 0
							}
						}

						if py >= h || py < 0 {
							if py >= h{
								py = h - 1
							}else{
								py = 0
							}
						}
						if sobelx == 1 && sobely == 1{
							gx += (int(img_data[py][px]) * kernelx[i][j])
							gy += (int(img_data[py][px]) * kernely[i][j])
						}else if sobelx == 1 {
							gx += (int(img_data[py][px]) * kernelx[i][j])

						}else{
							gy += (int(img_data[py][px]) * kernely[i][j])
						}

					}
				}
				if sobelx == 1 && sobely == 1{
					arr[y][x] = types.ImageType(math.Sqrt(math.Pow(float64(gx), 2.0) + math.Pow(float64(gy), 2.0)))
				}else if sobelx == 1 {
					arr[y][x] = types.ImageType(gx)
				}else{
					arr[y][x] = types.ImageType(gy)
				}
			}
		}(y)
	}
	wg.Wait()

	return arr
}