package main

import "fmt"

// 式スイッチ

func main() {

	// 変数を評価する場合、下記のフォーマットが望ましい
	switch a := 1; a {
	case 0,1:
		fmt.Println(a)
	case 2,3:
		fmt.Println("???")
	case 4,5:
		fmt.Println("Error")
	default:
		fmt.Println("No")
	}

	n := 5

	switch {
	case n > 0 && n < 4:
		fmt.Println("Pass")
	case n > 3 && n < 7:
		fmt.Println("Fail")
	}

	// 下記フォーマットは型違いでエラーになる（(type bool) as type int）
	// switch p := 3; p {
	// case p > 0 && p < 4:
	// 	fmt.Println("Pass")
	// case p > 3 && p < 7:
	// 	fmt.Println("Fail")
	// }
}