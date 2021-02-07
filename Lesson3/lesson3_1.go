package main

import "fmt"

func main() {
	var i int = 1
	var j int = 10

	// 変数の型を調べる
	fmt.Printf("%T\n", i)

	// 型変換
	fmt.Printf("%T\n", int32(j))

	// 浮動小数点型
	var fl float64 = 2.4
	// float32は、基本的には使用しない
	var a float32 = 3.1

	fmt.Println(fl)

	// 暗黙的な定義の場合、float64になる
	Fl := 2.6
	fmt.Println(fl + Fl)

	// 型を調べる
	fmt.Printf("%T, %T\n", fl, Fl)
	fmt.Printf("%T\n", a)

	// 型変換
	fmt.Printf("%T\n", float64(a))

	zero := 0.0

	// +Inf
	pinf := 1.0 / zero
	fmt.Println(pinf)

	// -Inf
	ninf := -1.0 / zero
	fmt.Println(ninf)

	// NaN
	nan := zero / zero
	fmt.Println(nan)
}
