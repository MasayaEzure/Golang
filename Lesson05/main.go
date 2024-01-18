package main

import "fmt"

func Sum(s ... int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	// スライス
	// 明示的な宣言
	var sl [] int = []int {1, 2, 3, 4}

	fmt.Println(sl)
	fmt.Println(sl[0])
	fmt.Println(sl[1])
	fmt.Println(sl[2])

	// 暗黙的な宣言
	sl2 := []int {10, 20, 30, 40}
	fmt.Println(sl2)

	// make関数（スライスを生成）　
	aaa := make([]string, 5)
	aaa[0] = "hello"
	aaa[1] = "good morning"
	aaa[2] = "good evening"
	aaa[3] = "good bye"
	aaa[4] = "see you"

	// すべての要素を出力
	fmt.Println(aaa[:])
	
	aaa[4] = "good night"
	
	// インデックス番号の範囲を指定して出力（3の手前まで）
	fmt.Println(aaa[:3])
	// 配列の一番最後の要素を出力
	fmt.Println(aaa[len(aaa) - 1])

	/*
	append
	要素を追加する
	第2引数に要素を複数渡すことができる
	*/
	sl = append(sl, 5, 6, 7)
	fmt.Println(sl)

	fmt.Println(len(sl))

	sample := make([]int, 3, 10) // 要素数、容量
	// cap (容量を出力)
	fmt.Println(cap(sample))

	slc := make([]int, 4, 10)
	fmt.Println(slc)

	// copy関数
	n := copy(slc, sl) // コピー先、コピー元
	fmt.Println(n, slc)

	str := []string{
		"aaa",
		"bbb",
		"ccc",
	}

	fmt.Println(str)
	
	for i, v := range str {
		fmt.Println(i, v)
	}

	fmt.Println(Sum(1,2,3,4,5,6))
	// 可変長引数でスライスを出力
	fmt.Println(Sum(sl...))
}