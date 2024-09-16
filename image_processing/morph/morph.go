package morph

import (
	"gocv/num"
	"gocv/types"
	"sync"
)

const (
	erosion_type = "Erosion"
	dilation_type = "Dilation"
)

func morph(img_data types.ImageArray, kernel_size int, operation_type string) types.ImageArray{
	arr := num.CreateArray3D(num.Shape(img_data))

	kernel_radius := kernel_size / 2
	height, width, _ := num.Shape(arr)
	var wg sync.WaitGroup

	for y := range img_data {

		wg.Add(1)
		go func(y int){

			defer wg.Done()
			for x := range img_data[0] {

				var px, py int
				var found_weak_edge, found_strong_edge bool = false, false

				for i := 0; i < kernel_size; i++ {
					for j := 0; j < kernel_size; j++ {
						px = x + j - kernel_radius
						py = y + i - kernel_radius

						if px >= width || px < 0 || py >= height || py < 0 || y == py || px == x{
							continue
						}else if (img_data[py][px][0] == 0 && operation_type==erosion_type) {
							found_weak_edge = true
							break
						}else if (img_data[py][px][0] == 255 && operation_type==dilation_type) {
							found_strong_edge = true
							break
						}
					}
					if found_weak_edge || found_strong_edge {
						break
					}
				}
				if operation_type==erosion_type{
					if found_weak_edge{
						arr[y][x][0] = 0
					}else{
						arr[y][x][0] = 255
					}
				}else if operation_type==dilation_type{
					if found_strong_edge{
						arr[y][x][0] = 255
					}
				}

			}
		}(y)
	}

	wg.Wait()
	return arr
}

func Erosion(img_data types.ImageArray) types.ImageArray{
	return morph(img_data, 5, erosion_type)
}


func Dilation(img_data types.ImageArray) types.ImageArray{
	return morph(img_data, 5, dilation_type)
}

func Opening(img_data types.ImageArray) types.ImageArray{
	return morph(morph(img_data, 5, erosion_type), 5, dilation_type)
}

func Closing(img_data types.ImageArray) types.ImageArray{
	return morph(morph(img_data, 5, dilation_type), 5, erosion_type)
}