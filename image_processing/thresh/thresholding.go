package thresh

import (
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

func Thresholding(img_data types.GrayImage, threshold types.ImageType, max_val types.ImageType, thresh_type int) types.GrayImage {

	h, w := len(img_data), len(img_data[0])
	thresh_img := make(types.GrayImage, h)

	var wg sync.WaitGroup

	for i := range thresh_img {
		thresh_img[i] = make([]types.ImageType, w)


		wg.Add(1)
		go func(i int){
			defer wg.Done()
			for j := range img_data[0]{

				if thresh_type == THRESH_BINARY{
					if img_data[i][j] > threshold{
						thresh_img[i][j] = max_val
					}
				}else if thresh_type == THRESH_BINARY_INV{
					if img_data[i][j] > threshold{
						thresh_img[i][j] = 0
					}else{
						thresh_img[i][j] = max_val
					}
				}else if thresh_type == THRESH_TOZERO{
					if img_data[i][j] > threshold{
						thresh_img[i][j] = img_data[i][j]
					}else{
						thresh_img[i][j] = 0
					}
				}else if thresh_type == THRESH_TOZERO_INV{
					if img_data[i][j] > threshold{
						thresh_img[i][j] = 0
					}else{
						thresh_img[i][j] = img_data[i][j]
					}
				}else if thresh_type == THRESH_TRUNC{
					if img_data[i][j] > threshold{
						thresh_img[i][j] = threshold
					}else{
						thresh_img[i][j] = img_data[i][j]
					}
				}
			}
		}(i)

	}
	wg.Wait()
	return thresh_img
}