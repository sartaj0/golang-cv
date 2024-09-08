package rotate

import (
	"gocv/num"
	"gocv/types"
	"math"
	"sync"
)

func RotateImage90(img_data types.ImageArray, clockwise bool) types.ImageArray {

	new_w, new_h, c := num.Shape(img_data)
	arr := num.CreateArray3D(new_h, new_w, c)

	for x := range img_data {
		for y := range img_data[0] {
			if !clockwise {
				arr[new_h-y-1][x] = img_data[x][y]
			} else {
				arr[y][new_w-x-1] = img_data[x][y]
			}
		}
	}

	return arr
}

func RotateImage180(img_data types.ImageArray) types.ImageArray {
	h, w, c := num.Shape(img_data)
	arr := num.CreateArray3D(h, w, c)
	for y := range img_data {
		for x := range img_data[0] {
			arr[y][x] = img_data[h-y-1][w-x-1]
		}
	}
	return arr

}

func RotateImageDegree(img_data types.ImageArray, degree float64) types.ImageArray {
	angle := (degree * math.Pi) / 180

	h, w, c := num.Shape(img_data)

	halfW, halfH := w/2, h/2

	x1, y1 := num.RotatePoints(-halfW, h-halfH, angle)
	x2, y2 := num.RotatePoints(w-halfW, h-halfH, angle)
	x3, y3 := num.RotatePoints(w-halfW, -halfH, angle)
	x4, y4 := num.RotatePoints(-halfW, -halfH, angle)

	new_adj_x := []float64{x1, x2, x3, x4}
	new_adj_y := []float64{y1, y2, y3, y4}

	new_w := int(math.Round(math.Abs(num.Max(new_adj_x) - num.Min(new_adj_x))))
	new_h := int(math.Round(math.Abs(num.Max(new_adj_y) - num.Min(new_adj_y))))

	arr := num.CreateArray3D(new_h, new_w, c)

	var wg sync.WaitGroup
	for y := range arr {

		wg.Add(1)

		go func(y int) {

			defer wg.Done()

			var newx, newy, adjx, adjy int
			var x_f, y_f float64

			for x := range arr[0] {
				adjx = x - (new_w / 2)
				adjy = new_h - (new_h / 2) - y

				x_f, y_f = num.RotatePoints(adjx, adjy, -angle)
				newx, newy = int(math.Round(x_f)), int(math.Round(y_f))

				newy = h - halfH - newy
				newx = halfW + newx

				if newy < 0 || newy >= len(img_data) || newx < 0 || newx >= len(img_data[0]) {
					continue
				}
				arr[y][x] = img_data[newy][newx]
			}
		}(y)
	}
	wg.Wait()

	return arr
}
