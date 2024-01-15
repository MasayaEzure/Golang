package main

import (
	"fmt"
)

// 定数
// 頭文字を大文字にすると、他のパッケージからも呼び出し可能
const  (
	URL = "https://xxx.co.jp"
	Name = "aaaaaa"
	Version = 1.2
)

const (
	A = 1
	B
	C = "aaa"
	// 連続する整数の連番を生成
	D = iota
	E
	F
)

func Sum (x, y int) int {
	result := x + y
	return result
}

func Div (x,y int) (int, int, int) {
	a := x * y
	b := x / y
	c := x % y
	return a, b, c
}

func Result (price int) (sum int) {
	sum = price * 2
	return
}

func Noreturn () {
	fmt.Println("aaa")
	return
}

func ReturnFunc () func() {
	return func() {
		fmt.Println("ReturnFunc")
	}
}

func CallFunc (f func()) {
	f()
}

func Closure() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func Gene() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main () {
	fmt.Println(URL, Name, Version)

	fmt.Println(A,B,C)
	fmt.Println(D,E,F)

	res := Sum(10, 100)
	fmt.Println(res)

	res2, res3, res4 := Div(10, 3)
	fmt.Println(res2, res3, res4)

	res4, res5, _ := Div(20, 6)  // _ で関数の返り値を破棄する
	fmt.Println(res4, res5)

	fmt.Println(Result(300))

	Noreturn()

	f := func (a, b, c int) int {
		result := a + b + c
		return result
	} (4, 5, 6)

	fmt.Println(f)

	r := ReturnFunc()
	r()

	CallFunc(func() {
		fmt.Println("CallFunc")
	})

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