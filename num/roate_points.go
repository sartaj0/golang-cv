package num

import "math"


func RotatePoints(x int, y int, angle float64) (new_x float64, new_y float64) {
	sin, cos := math.Sin(angle), math.Cos(angle) 
	new_x = cos * float64(x) -  sin * float64(y)

	new_y = sin * float64(x) + cos * float64(y)
	return new_x, new_y
}
