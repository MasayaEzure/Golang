package main

import "fmt"

func main() {
	sl := []string{
		"aaa",
		"bbb",
		"cccc",
	}
	fmt.Println(sl)

	// 第一引数：インデックス番号、第２引数：要素の値
	for i,v := range sl {
		fmt.Println(i,v)
	}

	// 古典的なfor
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

	fmt.Println(Sum(1,3,5))
	// 引数をいくつでも渡せる
	fmt.Println(Sum(1,1,1,1,1,1,1,1,1,1,1,1,1,1))

	a := []int{1,2,3}
	// 引数にスライスを渡すことも可能（...を必ずつける）
	fmt.Println(Sum(a...))

	// マップ
	// 明示的な宣言
	var m = map[string]int{
		"A":100,
		"B":200,
	}

	fmt.Println(m)
	fmt.Println(m["A"])

	// 暗黙的な宣言
	v := map[string]int {
		"a": 111,
		"b": 222,
	}

	fmt.Println(v)
	
	// 空のマップを生成
	y := make(map[int]string)
	fmt.Println(y)

	// 要素を追加
	y[0] = "aaa"
	y[1] = "bbb"
	y[2] = "ccc"
	y[3] = "ddd"

	fmt.Println(y)
	// 要素数を出力
	fmt.Println(len(y))

	// delete関数
	// 第一引数にmap、第二引数にキーを入力
	delete(y, 2)
	fmt.Println(y)

	// 空文字を出力
	fmt.Println(y[4])

	_,z:= y[4]
	// エラーハンドリング
	if !z {
		fmt.Println("Error")
	}
	fmt.Println(z)

	test := map[int]string{1:"aaa",2:"bbb",3:"ccc"}
	fmt.Println(test)

	// mapのfor文
	for k,v := range test {
		  fmt.Println(k,v)
	}
}

// 可変長引数
func Sum(s ...int) int {
	n := 0
	for _,v := range s {
		n += v
	}
	return n
}