package main

import "fmt"
import "strconv"

func main() {
	// エラーハンドリング
	var s string = "Hello"

	i, e := strconv.Atoi(s)
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("i = %T\n", i)
	}

	// 繰り返し処理
	a := 0
	for a < 10 {
		fmt.Println(a)
		a++
	}

	// 古典的な繰り返し文
	arrStr := [5]string {
		"あ",
		"い",
		"う",
		"え",
		"お",
	}

	for b := 0; b < 5; b++ {
		// bの値が2のとき、スキップする
		if b == 2 {
			continue
		// bの値が4のとき、処理中断
		} else if b == 4 {
			break
		} else {
		fmt.Println(arrStr[b])
		}
	}

	arrInt := [...]int{1,2,3,4,5}

	for i := 0; i < len(arrInt); i++ {
		if i == 1 {
			continue
		} else if i == 4 {
			break
		}
		fmt.Println(arrInt[i])
	}

	// １つ目にインデックス番号、２つ目に要素の値を返している
	for j,k := range arrInt {
		fmt.Println(j,k)
	}

	lang := []string{
		"Python",
		"PHP",
		"Golang",
	}
	for j,k := range lang {
		fmt.Println(j,k)
	}

	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k,v := range m {
		fmt.Println(k,v)
	}
}