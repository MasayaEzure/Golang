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

type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// カスタムする関数
func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name)
	} else {
		return (l[i].Value < l[j].Value)
	}
}

func main() {
	/*
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
	*/

	m := map[string]int{
		"S": 1,
		"J": 4,
		"A": 3,
		"M": 3,
	}

	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

	sort.Sort(lt)
	fmt.Println(lt)

	// リバース
	sort.Sort(sort.Reverse(lt))
	fmt.Println(lt)

}