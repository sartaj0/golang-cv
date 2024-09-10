package thresh

import (
	"gocv/cvt_color"
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

	if len(img_data[0][0]) == 3{
		img_data = cvt_color.RGBToGray(img_data)
	}

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

func DoubleThresholding(img_data types.ImageArray, thresh1 types.ImageType, thresh2 types.ImageType) types.ImageArray{

	
	var wg sync.WaitGroup

	for y := range img_data{

		wg.Add(1)

		go func (y int)  {

			defer wg.Done()

			for x := range img_data[0]{
				if img_data[y][x][0] >= thresh2{
					img_data[y][x][0] = 255
				}else if img_data[y][x][0] < thresh2 &&  img_data[y][x][0] > thresh1{
					img_data[y][x][0] = 128
				}else{
					img_data[y][x][0] = 0
				}
			}
		}(y)
	}

	wg.Wait()
	return img_data
}