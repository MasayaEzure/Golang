package main

import "fmt"

// 関数外で定義
var Number int = 1200

func main() {
	fmt.Println(Number)

	// 明示的な変数の定義(関数外でも定義可)
	// var 変数名 型 = 値

	// 数値
	var i int = 100
	fmt.Println(i)

	// 文字列
	var s string = "Hello Go"
	fmt.Println(s)

	// 論理値型
	var t,f bool = true, false
	fmt.Println(t,f)

	// 異なる型を定義
	// 変数の初期値（int：0、string：空文字）が格納されている
	var (
		num int
		str string
	)
	fmt.Println(num, str)

	// 値を代入
	num = 200
	str = "Golang"
	fmt.Println(num, str)

	// 暗黙的な変数の定義(関数内のみ定義可)
	// 変数名 := 値

	// 数値
	Num := 100
	fmt.Println(Num)

	// 値の更新
	Num = 150
	fmt.Println(Num)
	 
	// 関数の呼び出し
	sub()
}

func sub() {
	var a string = "Hello"
	fmt.Println(a)

	// 定義された変数は必ず使わなければならない
	var no string = "Not Use"
	fmt.Println(no)
}