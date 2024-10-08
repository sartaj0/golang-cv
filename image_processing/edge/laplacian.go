package edge

import (
	convolve "gocv/image_processing"
	"gocv/types"
)

func Laplacian(img_data types.ImageArray) (types.ImageArray, error){
	// laplacian_kernel := [][]float64{{-1, -1, -1}, {-1, 8, -1}, {-1, -1, -1}}
	laplacian_kernel := [][]float64{{0, 1, 0}, {1, -4, 1}, {0, 1, 0}}

	arr, err := convolve.ConvolveImage(img_data, 3, laplacian_kernel)
	return arr, err
}