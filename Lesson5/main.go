package main

import "fmt"


func main() {
	// 算術演算子
	fmt.Println(10 + 2)
	fmt.Println(10 - 2)
	fmt.Println(10 * 2)
	fmt.Println(10 / 2)
	fmt.Println(10 % 2)

	n := 0
	n += 5
	fmt.Println(n)

	n -= 3
	fmt.Println(n)

	n ++
	fmt.Println(n)

	n --
	fmt.Println(n)

	// 比較演算子
	a := 10
	fmt.Println(a == 10)
	fmt.Println(a != 10)

	b := 20
	fmt.Println(b != 10)
	fmt.Println(b == 10)
	fmt.Println(b >= 20)
	fmt.Println(b > 20)

	// 論理演算子
	fmt.Println(true && false == true)
	fmt.Println(true || false == true)
	fmt.Println(true != true)
	fmt.Println(false != true)
}