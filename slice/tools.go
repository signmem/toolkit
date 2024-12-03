package slice


func StringInSlice(a string, list []string) bool {

	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IntInSlice(a int, list []int) bool {

	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}


func Int64InSlice(a int64, list []int64) bool {

	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

