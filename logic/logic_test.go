package logic

import (
	"fmt"
	"testing"
)

func Test_duplicate(t *testing.T) {
	list := []int{1, 2, 3, 5, 6, 7, 8, 9}
	if !Duplicate(list, 5) {
		fmt.Println("list: ", list, "num: 5", Duplicate(list, 5))
		t.Error("should equal true")
	} else if Duplicate(list, 4) {
		fmt.Println("list: ", list, "num: 4", Duplicate(list, 4))
		t.Error("should equal false")
	}
}

func Test_addValue(t *testing.T) {
	list := []int{1, 2, 3, 5, 6, 7, 8, 9}
	if AddValue(list) != 4 {
		fmt.Println("list: ", list, AddValue(list))
		t.Error("should equal 4")
	}
}

func Test_numInSlice(t *testing.T) {
	list := []int{1, 2, 3, 5, 6, 7, 8, 9}
	if !numInSlice(5, list) {
		fmt.Println("list: ", list, "num: 5", numInSlice(5, list))
		t.Error("should equal true")
	} else if numInSlice(4, list) {
		fmt.Println("list: ", list, "num: 4", numInSlice(4, list))
		t.Error("should equal false")
	}
}
