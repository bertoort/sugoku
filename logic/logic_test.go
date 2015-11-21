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

func Test_levelNumber(t *testing.T) {
	h := LevelNumber("hard")
	m := LevelNumber("medium")
	e := LevelNumber("easy")
	if h < 81-27 || h > 81-23 {
		fmt.Printf("hard: %v", h)
		t.Error("hard difficulty number was not in range")
	} else if m < 81-32 && m > 81-28 {
		fmt.Printf("medium: %v", m)
		t.Error("medium difficulty number was not in range")
	} else if e < 81-37 && e > 81-33 {
		fmt.Printf("easy: %v", e)
		t.Error("easy difficulty number was not in range")
	}
}

func Test_randomDif(t *testing.T) {
	dif := RandomDif()
	err := true
	if dif == "hard" || dif == "medium" {
		err = false
	} else if dif == "easy" {
		err = false
	}
	if err {
		fmt.Printf("dif: %v", dif)
		t.Error("should equal hard, medium, or easy")
	}
}
