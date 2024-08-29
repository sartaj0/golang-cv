package num

import (
	"math"
)
func GetGaussianKernel(sigma float64, kernel_size int) [][]float64{

	var sum float64
	gaussian_kernel := make([][]float64, kernel_size)
	for i := range gaussian_kernel{
		gaussian_kernel[i] = make([]float64, kernel_size)
	}

	half := kernel_size / 2

	for y := -half; y <= half; y++{
		for x := -half; x <= half; x++{
			normal := 1 / (2.0 * math.Pi  * math.Pow(sigma, 2.0))
			exp_term := math.Exp(-(math.Pow(float64(x), 2.0) + math.Pow(float64(y), 2.0)) / (2.0 * math.Pow(sigma, 2.0)))

			gaussian_kernel[y+half][x+half] = normal * exp_term
			sum += normal * exp_term

		}
	}
	
	for y := range gaussian_kernel{
		for x:= range gaussian_kernel[0]{
			gaussian_kernel[y][x] /= sum
		}
	}

	return gaussian_kernel
}