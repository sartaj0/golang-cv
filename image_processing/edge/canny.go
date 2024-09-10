package edge

import (
	"gocv/image_processing/thresh"
	"gocv/num"
	"gocv/types"
	"math"
	"sync"
)

func non_max_suppression(arr types.ImageArray, theta_arr [][][]float64) types.ImageArray{
	h, w, _ := num.Shape(arr)

	var wg sync.WaitGroup
	Z := num.CreateArray3D(num.Shape(arr))
	for y := range Z{

		wg.Add(1)
		go func (y int)  {

			defer wg.Done()


			var adj_mag1, adj_mag2, pixel types.ImageType
			var degree float64
			var adj1px, adj1py, adj2px, adj2py int
			for x := range Z[0]{
				adj_mag1, adj_mag2 = 0, 0

				degree = num.RadianToDegree(theta_arr[y][x][0])
				if degree > 180{
					degree = math.Mod(degree, 180)
				}
				pixel = arr[y][x][0]

				if (degree >= 0 && degree <= 22.5) || (degree >= 157.5 && degree <= 180){
					adj1py, adj1px = y , x - 1
					adj2py, adj2px = y, x + 1
				}else if degree > 22.5 && degree <= 67.5{
					adj1py, adj1px = y - 1, x - 1
					adj2py, adj2px = y + 1, x + 1
				}else if degree > 67.5 && degree <= 112.5{
					adj1py, adj1px = y - 1, x 
					adj2py, adj2px = y + 1, x
				}else if degree > 112.5 && degree <= 157.5{
					adj1py, adj1px = y - 1, x + 1
					adj2py, adj2px = y + 1, x - 1
				}else{
					continue
				}


				if adj1px < 0 || adj1py < 0 || adj1px >= w || adj1py >= h{
					adj_mag1 = 255
				}else if adj2px < 0 || adj2py < 0 || adj2px >= w || adj2py >= h{
					adj_mag2 = 255
				}

				if (adj_mag1 != 255 && adj_mag2 != 255){
					adj_mag1 = arr[adj2py][adj2px][0]
					adj_mag2 = arr[adj2py][adj2px][0]
				}

				if pixel >= adj_mag1 && pixel >= adj_mag2{
					Z[y][x][0] = arr[y][x][0]
				}
				
			}
		}(y)
	}

	wg.Wait()
	return Z

}

func edge_tracking_hysteresis(img_data types.ImageArray) types.ImageArray{
	height, width, _ := num.Shape(img_data)
	var wg sync.WaitGroup

	for y := range img_data {

		wg.Add(1)
		go func(y int){

			defer wg.Done()
			for x := range img_data[0] {

				var px, py int
				var found_strong_edge bool = false

				if img_data[y][x][0] == 128{
					for i := 0; i < 3; i++ {
						for j := 0; j < 3; j++ {
							px = x + j - 1
							py = y + i - 1

							if px >= width || px < 0 || py >= height || py < 0 || y == py || px == x{
								continue
							}else if img_data[py][px][0] == 255{
								found_strong_edge = true
								break
							}
						}
						if found_strong_edge{
							break
						}
					}
					if found_strong_edge{
						img_data[y][x][0] = 255
					}else{
						img_data[y][x][0] = 0
					}
				}


			}
		}(y)
	}

	wg.Wait()
	return img_data
}

func Canny(img_data types.ImageArray, thresh1 types.ImageType, thresh2 types.ImageType) types.ImageArray {
	arr, theta_arr := SobelXY(img_data)
	arr = non_max_suppression(arr, theta_arr)
	arr = thresh.DoubleThresholding(arr, thresh1, thresh2)
	arr = edge_tracking_hysteresis(arr)
	return arr
}