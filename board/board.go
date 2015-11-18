package board

// Basic returns a standard input
func Basic() [9][9]int {
	var input [9][9]int
	input[0] = [9]int{1, 0, 3, 0, 0, 6, 0, 8, 0}
	input[1] = [9]int{0, 5, 0, 0, 8, 0, 1, 2, 0}
	input[2] = [9]int{7, 0, 9, 1, 0, 3, 0, 5, 6}
	input[3] = [9]int{0, 3, 0, 0, 6, 7, 0, 9, 0}
	input[4] = [9]int{5, 0, 7, 8, 0, 0, 0, 3, 0}
	input[5] = [9]int{8, 0, 1, 0, 3, 0, 5, 0, 7}
	input[6] = [9]int{0, 4, 0, 0, 7, 8, 0, 1, 0}
	input[7] = [9]int{6, 0, 8, 0, 0, 2, 0, 4, 0}
	input[8] = [9]int{0, 1, 2, 0, 4, 5, 0, 7, 8}
	return input
}

// Solved returns a solved basic standard input
func Solved() [9][9]int {
	var input [9][9]int
	input[0] = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	input[1] = [9]int{4, 5, 6, 7, 8, 9, 1, 2, 3}
	input[2] = [9]int{7, 8, 9, 1, 2, 3, 4, 5, 6}
	input[3] = [9]int{2, 3, 4, 5, 6, 7, 8, 9, 1}
	input[4] = [9]int{5, 6, 7, 8, 9, 1, 2, 3, 4}
	input[5] = [9]int{8, 9, 1, 2, 3, 4, 5, 6, 7}
	input[6] = [9]int{3, 4, 5, 6, 7, 8, 9, 1, 2}
	input[7] = [9]int{6, 7, 8, 9, 1, 2, 3, 4, 5}
	input[8] = [9]int{9, 1, 2, 3, 4, 5, 6, 7, 8}
	return input
}

// Broken returns an unsolvable standard input
func Broken() [9][9]int {
	var input [9][9]int
	input[0] = [9]int{1, 1, 3, 0, 0, 6, 0, 8, 0}
	input[1] = [9]int{0, 5, 0, 0, 8, 0, 1, 2, 0}
	input[2] = [9]int{7, 0, 9, 1, 0, 3, 0, 5, 6}
	input[3] = [9]int{0, 3, 0, 0, 6, 7, 0, 9, 0}
	input[4] = [9]int{5, 0, 7, 8, 0, 0, 0, 3, 0}
	input[5] = [9]int{8, 0, 1, 0, 3, 0, 5, 0, 7}
	input[6] = [9]int{0, 4, 0, 0, 7, 8, 0, 1, 0}
	input[7] = [9]int{6, 0, 8, 0, 0, 2, 0, 4, 0}
	input[8] = [9]int{0, 1, 2, 0, 4, 5, 0, 7, 8}
	return input
}

// Medium returns a standard input
func Medium() [9][9]int {
	var input [9][9]int
	input[0] = [9]int{9, 3, 8, 0, 0, 0, 1, 0, 5}
	input[1] = [9]int{0, 0, 4, 8, 6, 0, 0, 0, 0}
	input[2] = [9]int{6, 7, 0, 0, 0, 0, 0, 8, 0}
	input[3] = [9]int{0, 0, 5, 0, 8, 2, 0, 0, 0}
	input[4] = [9]int{0, 6, 0, 0, 0, 0, 0, 7, 0}
	input[5] = [9]int{0, 0, 0, 6, 9, 0, 5, 0, 0}
	input[6] = [9]int{0, 1, 0, 0, 0, 0, 0, 4, 2}
	input[7] = [9]int{0, 0, 0, 0, 4, 6, 8, 0, 0}
	input[8] = [9]int{8, 0, 3, 0, 0, 0, 6, 9, 7}
	return input
}

// Hard returns a standard input
func Hard() [9][9]int {
	var input [9][9]int
	input[0] = [9]int{0, 0, 8, 0, 0, 0, 2, 0, 0}
	input[1] = [9]int{0, 0, 3, 0, 0, 5, 0, 0, 1}
	input[2] = [9]int{0, 6, 0, 9, 0, 0, 0, 0, 8}
	input[3] = [9]int{0, 9, 0, 0, 5, 0, 8, 0, 0}
	input[4] = [9]int{6, 8, 0, 0, 7, 0, 0, 9, 4}
	input[5] = [9]int{0, 0, 2, 0, 1, 0, 0, 7, 0}
	input[6] = [9]int{7, 0, 0, 0, 0, 1, 0, 8, 0}
	input[7] = [9]int{3, 0, 0, 2, 0, 0, 5, 0, 0}
	input[8] = [9]int{0, 0, 6, 0, 0, 0, 4, 0, 0}
	return input
}
