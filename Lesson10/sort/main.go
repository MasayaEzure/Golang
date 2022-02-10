package main

import (
	"fmt"
	"sort"
)

// sort
type Entry struct {
	Name string
	Value int
}

func main() {
	i := []int{5,4,3,2,1,7,8,6,9,0}
	s := []string{"a", "y", "g"}

	fmt.Println(i, s)

	sort.Ints(i)
	sort.Strings(s)
	// ソートされた順番で出力
	fmt.Println(i, s)

	el := []Entry {
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}

	fmt.Println(el)

	// Slice
	sort.Slice(el, func(i, j int) bool {return el[i].Name < el[j].Name})
	fmt.Println(el)

	// SliceStable
	sort.SliceStable(el, func(i, j int) bool {return el[i].Name < el[j].Name})
	fmt.Println(el)

}