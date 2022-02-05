package main

import (
	"fmt"
	"Lesson9/index"
)

func IsTrue(i int) bool {
	if i >= 0 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsTrue(1))
	fmt.Println(IsTrue(-1))

	s := []int {1,2,3,4,5}
	fmt.Println(index.Average(s))
}