package blur

import (
	"errors"
	"gocv/num"
	"gocv/types"
	"sync"
)



func GaussianBlur(img_data types.ColorImage, kernel_size int) (types.ColorImage, error ){
	h, w, c := num.Shape(img_data)

	// isrgb, _ := num.IsRGB(img_data)
	gaussian_kernel := num.GetGaussianKernel(1, kernel_size)


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

				var r, g, b float64
				var i_k, j_k int

				
				for i := y - kernel_radius; i <= y+kernel_radius; i++{
					j_k = 0
					for j := x - kernel_radius; j <= x+kernel_radius; j++{
						
						if j < 0 || j >=w || i < 0 || i >=h{
							continue
						}
						r += float64(img_data[i][j][0]) * gaussian_kernel[i_k][j_k]
						g += float64(img_data[i][j][1]) * gaussian_kernel[i_k][j_k]
						b += float64(img_data[i][j][2]) * gaussian_kernel[i_k][j_k]

						j_k += 1
					}
					i_k += 1
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