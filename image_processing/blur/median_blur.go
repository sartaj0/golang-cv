package blur

import (
	"errors"
	"gocv/num"
	"gocv/types"
	"sync"
)

func MedianBlur(img_data types.ColorImage, kernel_size int) (types.ColorImage, error ){
	h, w, c := num.Shape(img_data)

	arr := num.CreateArray(h, w, c)
	if kernel_size % 2 == 0 || kernel_size < 1{
		return nil, errors.New("kernel size cannot be even, 0 or negative number")
	}
	kernel_radius := kernel_size / 2
	
	var wg sync.WaitGroup

	for y := range img_data {
		wg.Add(1)
		go func (y int){
			defer wg.Done()


			for x := range img_data[0]{
				var r, g, b []types.ImageType
				for i := y - kernel_radius; i <= y+kernel_radius; i++{
					for j := x - kernel_radius; j <= x+kernel_radius; j++{
						if j < 0 || j >=w{
							continue
						}

						if i < 0 || i >=h {
							continue
						}

						r = append(r, img_data[i][j][0])
						g = append(g, img_data[i][j][1])
						b = append(b, img_data[i][j][2])

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


