package num

import (
	"math"
)


func Min(arr []float64) (min_value float64) {
	min_value = math.Inf(1)

	for _, element := range arr{
		if element < min_value{
			min_value = element
		}
	}

	return min_value
}



func Max(arr []float64) (max_value float64) {
	for _, element := range arr{
		if element > max_value{
			max_value = element
		}
	}
	return max_value
}