package num 

import (
	"sort"
	"gocv/types"
)
func CreateArray(h int, w int, c int) types.ColorImage {
	arr := make(types.ColorImage, h)
	for y := range arr {
		arr[y] = make([][]types.ImageType, w)
		for x := range arr[y] {
			arr[y][x] = make([]types.ImageType, c)
		}
	}
	return arr
}


func Shape(img_data types.ColorImage) (int, int, int) {
	return len(img_data), len(img_data[0]), len(img_data[0][0])

}

func SortArray(arr []types.ImageType) []types.ImageType {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}

func Median(arr []types.ImageType) types.ImageType {
	arr = SortArray(arr)
	i := (len(arr) + 1) / 2 
	return arr[i]
}