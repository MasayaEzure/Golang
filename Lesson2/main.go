package main

import (
	"fmt"
	"strconv"
)

func  main()  {
	// int型
	var i int8 = 10
	var j int64 = 20

	// int型がそれぞれ異なるため、ミスマッチが起こる
	// fmt.Println(i + j)

	// 型を表示(%Tは型を表示する書式)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", j)

	// キャスト
	fmt.Printf("%T\n", int64(i))
	fmt.Println(int64(i) + j)

	// float型
	var f float64 = 1.2
	var g float32 = 3.4

	// 暗黙的に定義した場合、自動的に float64 になる
	aaa := 1.23

	fmt.Println(f)
	fmt.Println(g)
	fmt.Printf("%T\n", aaa)

	// String型
	var s string = "abcd"
	fmt.Println(s, `"`)
	// 文字列の要素を取得
	fmt.Println(string(s[0]))
	fmt.Println(string(s[1]))
	fmt.Println(string(s[2]))
	fmt.Println(string(s[3]))

	// このままの書式で出力する(バッククオーテーションを使う)
	fmt.Println(`aaa
	aaaa
		aaaa`)

	// byte型
	b := []byte{72,73}
	fmt.Println(b)

	// キャスト
	fmt.Println(string(b))
	c := []byte("HI")
	fmt.Println(c)

	/*
	配列型
	明示的な定義
	var 配列名 [要素数]型 {値}
	要素のサイズを変更することはできない
	*/
	var sample [5]int = [5]int {1,2,3,4,5}
	fmt.Println(sample)
	fmt.Printf("%T\n", sample)
	 
	/*
	暗黙的な定義
	配列名 := [要素数]型 {値}
	[...]で要素数を省略することができる(要素数を自動でカウント)
	*/
	sample2 := [3]string {"aaa","bbb","ccc"}
	fmt.Println(sample2)
	fmt.Printf("%T\n", sample2)
	 
	sample3 := [...]int {10,20,30}
	fmt.Println(sample3)
	fmt.Printf("%T\n", sample3)

	// 値の更新
	sample[0] = 10
	sample2[1] = "zzz"
	sample3[2] = 100

	// 値の取り出し
	fmt.Println(sample[0])
	fmt.Println(sample2[1])
	fmt.Println(sample3[2])

	// 配列の代入
	var sample4 [5]int
	sample4 = sample
	fmt.Println(sample4)

	// 配列の要素数を取得
	fmt.Println(len(sample2))

	/*
	インターフェース
	すべての型との互換性を持つ
	初期値は「nil」
	*/
	var x interface{}
	x = 100
	fmt.Println(x)
	
	x = "xxx"
	fmt.Println(x)

	x = [3]int{1,2,3}
	fmt.Println(x)

	// 型変換（キャスト）
	var number int = 123
	numberFlt := float64(number)

	fmt.Printf("%T\n", number)
	fmt.Printf("%T\n", numberFlt)

	var aiu string = "100"
	// string から int への型変換
	aiuInt, err := strconv.Atoi(aiu)

	// エラーハンドリング
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(aiuInt)
	fmt.Printf("%T\n", aiuInt)

	// int から string への型変換
	var num int = 200
	numStr := strconv.Itoa(num)
	fmt.Println(numStr)
	fmt.Printf("%T\n", numStr)

	// string から byte への型変換
	var greet string = "Hello!"
	greetByte := []byte(greet)
	fmt.Println(greetByte)
	fmt.Printf("%T\n", greetByte)

	// byte から string への型変換
	greetStr := string(greetByte)
	fmt.Println(greetStr)
	fmt.Printf("%T\n", greetStr)

}