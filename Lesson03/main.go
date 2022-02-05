package main

import (
	"fmt"
)

/*
定数
基本は関数外に記述する
頭文字を大文字にすると、他のパッケージからも呼び出し可能
*/
const  (
	URL = "https://xxx.co.jp"
	Name = "aaaaaa"
	Version = 1.2
)

const (
	A = 1
	B
	C = "aaa"
	// itoa：連続する整数の連番を生成
	D = iota
	E
	F
)

// 返り値が一つの関数
func Sum (x, y int) int {
	result := x + y
	return result
}

// 返り値が複数の関数
func Div (x,y int) (int, int, int) {
	a := x * y
	b := x / y
	c := x % y
	return a, b, c
}

// 返り値に変数を定義する関数
func Result (price int) (sum int) {
	sum = price * 2
	// 変数 sum を返す
	return
}

// 返り値がない関数
func Noreturn () {
	fmt.Println("aaa")
	return
}

// 関数を返す関数
func ReturnFunc () func() {
	return func() {
		fmt.Println("ReturnFunc")
	}
}

// 関数を引数に取る関数
func CallFunc (f func()) {
	f()
}

// クロージャー
func Closure() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// ジェネレーター
func Gene() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}


func main () {
	fmt.Println(URL, Name, Version)
	// 値の上書き
	fmt.Println(A,B,C)
	fmt.Println(D,E,F)

	// 関数の呼び出し
	res := Sum(10, 100)
	fmt.Println(res)

	res2, res3, res4 := Div(10, 3)
	fmt.Println(res2, res3, res4)

	// _ で関数の返り値を破棄することができる
	res4, res5, _ := Div(20, 6)
	fmt.Println(res4, res5)

	fmt.Println(Result(300))

	Noreturn()

	// 無名関数
	f := func (a, b, c int) int {
		result := a + b + c
		return result
	} (4, 5, 6)

	fmt.Println(f)

	// 関数の呼び出し
	r := ReturnFunc()
	r()

	// 関数の実行
	CallFunc(func() {
		fmt.Println("CallFunc")
	})

	// クロージャーの実行
	c := Closure()
	fmt.Println(c("aaaa"))
	fmt.Println(c("bbb"))
	fmt.Println(c("cccc"))

	g := Gene()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	
}