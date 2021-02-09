package main

import "fmt"

// 定数（関数外から書く）
const Pi = 3.14

// 複数定義
const (
	URL = "https://aaa.co.jp"
	name = "test"
)

// bとcにはaの値、eとfにはdの値がそれぞれ代入される
const (
	a = 1
	b
	c
	d = "aaa"
	e
	f
)

// iota：連続する整数の連番を生成する
const (
	u = iota
	v
	w
	x
	y
	z
)

// 変数の最大値であっても定数は出力可
const i = 9223372036854775807 + 1

func main() {
	fmt.Println(Pi)

	// 一度定義した定数は、値の変更が不可（cannot assign to Pi）
	// Pi = 3.15
	// fmt.Println(Pi) 

	fmt.Println(URL, name)
	fmt.Println(a,b,c,d,e,f)
	fmt.Println(i - 1)
	fmt.Println(u,v,w,x,y,z)
}