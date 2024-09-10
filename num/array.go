package num

import (
	"gocv/types"
	"sort"
)

func CreateArray3D(h int, w int, c int) types.ImageArray {
	arr := make(types.ImageArray, h)
	for y := range arr {
		arr[y] = make([][]types.ImageType, w)
		for x := range arr[y] {
			arr[y][x] = make([]types.ImageType, c)
		}
	}
	return arr
}


func Shape(img_data types.ImageArray) (int, int, int) {
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

func UniqueValueArray(img_data types.ImageArray) []types.ImageType{
	
	var arr []types.ImageType
	arr_track := make(map[types.ImageType]bool)
	for y := range img_data{
		for x := range img_data[0]{
			if arr_track[img_data[y][x][0]] {
				continue
			}
			arr_track[img_data[y][x][0]] = true
			arr = append(arr, img_data[y][x][0])
		}
	}
	arr =SortArray(arr)
	return arr
}