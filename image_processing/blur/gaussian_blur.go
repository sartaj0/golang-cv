package blur

import (
	convolve "gocv/image_processing"
	"gocv/num"
	"gocv/types"
)

func GaussianBlur(img_data types.ImageArray, kernel_size int, sigma float64) (types.ImageArray, error) {

	gaussian_kernel := num.GetGaussianKernel(sigma, kernel_size)

	img, err := convolve.ConvolveImage(img_data, kernel_size, gaussian_kernel)
	return img, err
}
