package main

import "fmt"

// 配列型

func main() {
	// 明示的な宣言
	// var 配列名 [数字]型 = [数字]型{要素}
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	// インデックス番号を取得し、配列の値を表示する
	fmt.Println(arr[0])
	arr[1] = 6
	fmt.Println(arr[1])

	// 暗黙的な定義
	// 配列名 := [数字]型{要素}
	// ...：要素数を自動で取得するもの
	arrStr := [...]string{
		"Hello",
		"World",
		"Golang",
		"Yeah",
	}
	fmt.Println(arrStr)
	// fmt.Printf("%T\n", arrStr)

	arrStr[3] = "Great"
	fmt.Println(arrStr[3])
	
	// 配列の要素数を調べる
	fmt.Println(len(arrStr))
}