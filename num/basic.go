package num


func MinValue(values ...int) int {
	min_value := values[0]

	for _, element := range values[1:]{
		if element < min_value{
			min_value = element
		}
	}

	return min_value
}


func MaxValue(values ...int) (max_value int) {

	for _, element := range values{
		if element > max_value{
			max_value = element
		}
	}

	return max_value
}
