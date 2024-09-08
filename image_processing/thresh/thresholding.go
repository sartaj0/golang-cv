package thresh

import (
	"gocv/num"
	"gocv/types"

	"sync"
)


const (
    THRESH_BINARY = iota
	THRESH_BINARY_INV
	THRESH_TRUNC
	THRESH_TOZERO
	THRESH_TOZERO_INV
)

func Thresholding(img_data types.ImageArray, threshold types.ImageType, max_val types.ImageType, thresh_type int) types.ImageArray {

	thresh_img := num.CreateArray3D(num.Shape(img_data))

	var wg sync.WaitGroup

	for i := range thresh_img {


		wg.Add(1)
		go func(i int){
			defer wg.Done()
			for j := range img_data[0]{

				if thresh_type == THRESH_BINARY{
					if img_data[i][j][0] > threshold{
						thresh_img[i][j][0] = max_val
					}
				}else if thresh_type == THRESH_BINARY_INV{
					if img_data[i][j][0] > threshold{
						thresh_img[i][j][0] = 0
					}else{
						thresh_img[i][j][0] = max_val
					}
				}else if thresh_type == THRESH_TOZERO{
					if img_data[i][j][0] > threshold{
						thresh_img[i][j][0] = img_data[i][j][0]
					}else{
						thresh_img[i][j][0] = 0
					}
				}else if thresh_type == THRESH_TOZERO_INV{
					if img_data[i][j][0] > threshold{
						thresh_img[i][j][0] = 0
					}else{
						thresh_img[i][j][0] = img_data[i][j][0]
					}
				}else if thresh_type == THRESH_TRUNC{
					if img_data[i][j][0] > threshold{
						thresh_img[i][j][0] = threshold
					}else{
						thresh_img[i][j][0] = img_data[i][j][0]
					}
				}
			}
		}(i)

	}
	wg.Wait()
	return thresh_img
}