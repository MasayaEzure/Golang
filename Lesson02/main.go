package main

import (
	"fmt"
	"strconv"
)

func  main()  {
	var i int8 = 10
	var j int64 = 20
	// fmt.Println(i + j)

	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", j)

	fmt.Printf("%T\n", int64(i))
	fmt.Println(int64(i) + j)

	var f float64 = 1.2
	var g float32 = 3.4

	// 暗黙的に定義した場合、float64 になる
	aaa := 1.23

	fmt.Println(f)
	fmt.Println(g)
	fmt.Printf("%T\n", aaa)

	var s string = "abcd"
	fmt.Println(s, `"`)

	fmt.Println(string(s[0]))
	fmt.Println(string(s[1]))
	fmt.Println(string(s[2]))
	fmt.Println(string(s[3]))

	fmt.Println(`aaa
	aaaa
		aaaa`)

	b := []byte{72,73}
	fmt.Println(b)

	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)


	// 配列
	// 明示的な定義
	var sample [5]int = [5]int {1,2,3,4,5}
	fmt.Println(sample)
	fmt.Printf("%T\n", sample)
	 
	// 暗黙的な定義
	sample2 := [3]string {"aaa","bbb","ccc"}
	fmt.Println(sample2)
	fmt.Printf("%T\n", sample2)
	 
	sample3 := [...]int {10,20,30}
	fmt.Println(sample3)
	fmt.Printf("%T\n", sample3)

	sample[0] = 10
	sample2[1] = "zzz"
	sample3[2] = 100

	fmt.Println(sample[0])
	fmt.Println(sample2[1])
	fmt.Println(sample3[2])

	var sample4 [5]int
	sample4 = sample
	fmt.Println(sample4)

	fmt.Println(len(sample2))

	// インターフェース
	var x interface{}
	x = 100
	fmt.Println(x)
	
	x = "xxx"
	fmt.Println(x)

	x = [3]int{1,2,3}
	fmt.Println(x)

	var number int = 123
	numberFlt := float64(number)

	fmt.Printf("%T\n", number)
	fmt.Printf("%T\n", numberFlt)

	var numberStr string = "100"
	numInt, err := strconv.Atoi(numberStr)  // string → int

	// エラーハンドリング
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(numInt)
	fmt.Printf("%T\n", numInt)

	var num int = 200
	numStr := strconv.Itoa(num)  // int → string
	fmt.Println(numStr)
	fmt.Printf("%T\n", numStr)

	var greet string = "Hello!"
	greetByte := []byte(greet)  // string → byte
	fmt.Println(greetByte)
	fmt.Printf("%T\n", greetByte)

	greetStr := string(greetByte)  // byte → string
	fmt.Println(greetStr)
	fmt.Printf("%T\n", greetStr)

}