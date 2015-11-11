package logic

// Duplicate checks if a value is not in an array
func Duplicate(not []int, n int) bool {
	var d bool
	for _, num := range not {
		if num == n {
			d = true
		}
	}
	return d
}

// AddValue will return the number not in the array
func AddValue(a []int) int {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var n int
	for _, v := range list {
		if !numInSlice(v, a) {
			n = v
		}
	}
	return n
}

// numInSlice checks if value is not in the array
func numInSlice(num int, list []int) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}
