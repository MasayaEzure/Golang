// パッケージの宣言は一つまで
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println(time.Now())

	// 明示的な変数定義
	var i int = 100
	fmt.Println(i)

	var s, t, r string = "aaa", "bbb", "ccc"
	fmt.Println(s, t, r) 

	var (
		x int = 123
		y string = "hello"
		z bool = true
	)
	fmt.Println(x, y, z)

	x = 45
	y = "Good"
	z = false
	fmt.Println(x,y,z)

	// 暗黙的な変数定義
	str := "Good Night"
	number := 678
	fmt.Println(str, number)

	str = "Good Morning"
	number = 9
	fmt.Println(str, number)

	fmt.Println(aaa)
	sample()
}

func sample() {
	var ex string = "abcdef"
	fmt.Println(ex)

}

var aaa int = 111