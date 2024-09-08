package edge

import (
	"gocv/num"
	"gocv/types"
	"math"
	"sync"
)

func Sobel(img_data types.ImageArray, sobelx int, sobely int) types.ImageArray {
	kernel_size := 3
	kernel_radius := kernel_size / 2

	h, w, _ := num.Shape(img_data)
	kernelx := [][]float64{
		{-1, 0, 1}, 
		{-2, 0, 2}, 
		{-1, 0, 1}}

	kernely := [][]float64{
		{-1, -2, -1},
		{0,  0,  0},
		{1,  2,  1}}
	

	arr := num.CreateArray3D(h, w, 1)

	float_arr := make([][][]float64, h)
	for y := range float_arr{
		float_arr[y] = make([][]float64, w)
		for x := range float_arr[0]{
			float_arr[y][x] =  make([]float64, 1)
		}
	}

	var wg sync.WaitGroup

	for y := range img_data {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()

			for x := range img_data[0] {
				var px, py int
				var gx, gy float64

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
							gx += (float64(img_data[py][px][0]) * kernelx[i][j])
							gy += (float64(img_data[py][px][0]) * kernely[i][j])
						}else if sobelx == 1 {
							gx += (float64(img_data[py][px][0]) * kernelx[i][j])

						}else{
							gy += (float64(img_data[py][px][0]) * kernely[i][j])
						}

					}
				}
				if sobelx == 1 && sobely == 1{
					arr[y][x][0] = types.ImageType(math.Sqrt(gx * gx + gy * gy))
					// arr[y][x][0] = types.ImageType(math.Atan2(gy, gx))
					// arr[y][x][0] = types.ImageType(math.Abs(float64(gx)) + math.Abs(float64(gy)))
				}else if sobelx == 1 {
					arr[y][x][0] = types.ImageType(gx)
				}else{
					arr[y][x][0] = types.ImageType(gy)
				}
			}
		}(y)
	}
	wg.Wait()

	return arr
}