package num

import (
	"errors"
	"gocv/types"
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

/*
func processArray(arr interface{}) (bool, error) {
    switch v := arr.(type) {
    case types.ColorImage:
        return true, nil
    case types.GrayImage:
        return false, nil
        // Process 2D array
    default:
        return nil, errors.New("Unsupported type")
    }
}
*/

func IsRGB(arr interface{}) (bool, error) {
		
    switch arr.(type){
      
	case types.ColorImage:
		return true, nil


	case types.GrayImage:
		return false, nil

	default:
		return false, errors.New("unsupported variable") 
	}
}