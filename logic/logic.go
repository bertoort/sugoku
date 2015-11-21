package logic

import (
	"math/rand"
	"time"
)

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

// Shuffle randomizes an array
func Shuffle(a []int) {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i)
		a[i], a[j] = a[j], a[i]
	}
}

// LevelNumber returns the number of values to remove to set a difficulty
// hard: <27, medium: 28-32, easy: 33>
func LevelNumber(dif string) int {
	hard := [5]int{23, 24, 25, 26, 27}
	med := [5]int{28, 29, 30, 31, 32}
	easy := [5]int{33, 34, 35, 36, 37}
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	r := rand.Intn(4)
	var n int
	if dif == "hard" {
		n = hard[r]
	} else if dif == "medium" {
		n = med[r]
	} else {
		n = easy[r]
	}
	return 81 - n
}

// RandomDif will return a random string: easy, medium, hard
func RandomDif() string {
	t := time.Now()
	rand.Seed(int64(t.Nanosecond()))
	r := rand.Intn(2)
	if r == 2 {
		return "hard"
	} else if r == 1 {
		return "medium"
	}
	return "easy"
}
