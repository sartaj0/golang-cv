package num

import (
	"math"
)
func MinValue(values ...int) int {
	min_value := values[0]

	for _, element := range values[1:] {
		if element < min_value {
			min_value = element
		}
	}

	return min_value
}

func MaxValue(values ...int) (max_value int) {

	for _, element := range values {
		if element > max_value {
			max_value = element
		}
	}

	return max_value
}

func RadianToDegree(radian float64) float64 {
	degree := math.Mod(radian * 180 / math.Pi, 360)
	if degree < 0{
		degree += 360
	}
	return degree
}

func DegreeToRadian(degree float64) float64 {
	angle := math.Mod(degree, 360)
	if angle < 0{
		angle += 360
	}
	angle = angle * (math.Pi / 180)

	return angle
}