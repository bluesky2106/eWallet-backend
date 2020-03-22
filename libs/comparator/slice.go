package comparator

// SliceContainsFloat : used to check in int slice contains number
func SliceContainsFloat(floatSlice []float64, searchFloat float64) bool {
	if len(floatSlice) > 0 {
		for _, value := range floatSlice {
			if value == searchFloat {
				return true
			}
		}
	}
	return false
}

// SliceContainsInt : used to check in int slice contains number
func SliceContainsInt(intSlice []int, searchInt int) bool {
	if len(intSlice) > 0 {
		for _, value := range intSlice {
			if value == searchInt {
				return true
			}
		}
	}
	return false
}

// SliceContainsString : used to check in string slice contains string
func SliceContainsString(stringSlice []string, searchString string) bool {
	if len(stringSlice) > 0 {
		for _, value := range stringSlice {
			if value == searchString {
				return true
			}
		}
	}
	return false
}
