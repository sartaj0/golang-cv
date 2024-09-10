package edge

import (
	convolve "gocv/image_processing"
	"gocv/num"
	"gocv/types"
	"math"
	"sync"
)


func SobelXY(img_data types.ImageArray) (types.ImageArray, [][][]float64){
	h, w, _ := num.Shape(img_data)

	kernel_size := 3
	kernel_radius := kernel_size / 2
	
	kernelx :=  num.GetSobelXKernel(kernel_size)
	kernely := num.GetSobelYKernel(kernel_size)
	

	arr := num.CreateArray3D(h, w, 1)

	float_arr := make([][][]float64, h)
	for y := range float_arr{
		float_arr[y] = make([][]float64, w)
		for x := range float_arr[0]{
			float_arr[y][x] =  make([]float64, 1)
		}
	}

	theta_arr := make([][][]float64, h)
	for y := range theta_arr{
		theta_arr[y] = make([][]float64, w)
		for x := range theta_arr[0]{
			theta_arr[y][x] =  make([]float64, 1)
		}
	}

	var wg sync.WaitGroup
	min_val, max_val := math.Inf(1), 0.0
	for y := range img_data {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()

			for x := range img_data[0] {
				var px, py int
				var gx, gy, value float64

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
						gx += (float64(img_data[py][px][0]) * kernelx[i][j])
						gy += (float64(img_data[py][px][0]) * kernely[i][j])


					}
				}
				
				value = math.Sqrt(gx * gx + gy * gy)

				if value > max_val{
					max_val = value
				}else if value < min_val{
					min_val = value
				}

				float_arr[y][x][0] = value
				theta_arr[y][x][0] = math.Atan2(gy, gx)

			}
		}(y)
	
	}
	wg.Wait()

	if max_val != 0{
		for y := range float_arr{
			wg.Add(1)
			go func(y int){
				defer wg.Done()
	
				for x := range float_arr[0]{
					arr[y][x][0] = types.ImageType((float_arr[y][x][0] * 255) / max_val)
				}
			}(y)
		}
		wg.Wait()	
	}


	return arr, theta_arr
}

func Sobel(img_data types.ImageArray, sobelx int, sobely int) types.ImageArray {

	kernel_size := 3
	var arr types.ImageArray
	// kernel_radius := kernel_size / 2

	if sobelx == 1 && sobely == 1 {
		arr, _ = SobelXY(img_data)
		
	}else if sobelx == 1{
		arr, _ = convolve.ConvolveImage(img_data, kernel_size, num.GetSobelXKernel(kernel_size))
	}else if sobely == 1{
		arr, _ = convolve.ConvolveImage(img_data, kernel_size, num.GetSobelYKernel(kernel_size))
	}

	return arr
}