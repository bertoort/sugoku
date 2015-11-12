package board

import (
	"fmt"
	"testing"
)

func Test_boardBasic(t *testing.T) {
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
	if Basic() != input {
		fmt.Println(Basic())
		t.Error("should output a 2D matrix input")
	}
}
