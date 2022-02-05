// パッケージの宣言は一つまで
package main

import (
	"fmt"
	"time"
)

// メイン関数
func main() {
	fmt.Println("Hello world")
	// 現在の時刻を出力
	fmt.Println(time.Now())

	/*
	変数
	明示的な定義
	var 変数名 型 = 値
	メイン関数外でも定義できる
	*/
	var i int = 100
	fmt.Println(i)

	var s, t, r string = "aaa", "bbb", "ccc"
	fmt.Println(s, t, r) 

	// 異なる型で定義
	var (
		x int = 123
		y string = "hello"
		z bool = true
	)
	fmt.Println(x, y, z)

	// 値の更新
	x = 45
	y = "Good"
	z = false
	fmt.Println(x,y,z)

	/*
	暗黙的な定義
	変数名 := 値
	型指定が不要
	*/
	str := "Good Night"
	number := 678
	fmt.Println(str, number)

	// 値の更新
	str = "Good Morning"
	number = 9
	fmt.Println(str, number)

	// 変数の出力
	fmt.Println(aaa)

	// 関数の呼び出し
	sample()
}

func sample() {
	var ex string = "abcdef"
	fmt.Println(ex)

}

var aaa int = 111