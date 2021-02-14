package main

import "fmt"

// ブランチをtestに変更

// スライス
func main(){
	// 明示的な宣言
	// []内には要素数を指定することができない
	var sl []int = []int{1,2,3}
	fmt.Println(sl)

	// 暗黙的な宣言
	Sl := []int{4,5,6}
	// すべての要素を取り出す
	fmt.Println(Sl[:])

	// 値の代入
	Sl[2] = 7
	fmt.Println(Sl[:])

	// インデックス番号を連続で取り出す
	fmt.Println(Sl[1:3])

	// 指定したインデックス番号の手前まで取り出す
	fmt.Println(Sl[:2])

	// 指定したインデックス番号から最後まで取り出す
	fmt.Println(Sl[0:])

	// 最後の要素を取り出す
	fmt.Println(Sl[1:len(Sl)-1])

	// make関数
	// 第２引数に要素数を指定できる
	SL := make([]int, 10)
	fmt.Println(SL)

	// len：要素数を示す関数
	fmt.Println(len(SL))
	// cap関数
	fmt.Println(cap(SL))

	// append:要素の追加
	sl = append(sl,4,5,6,7,8,9,10)
	fmt.Println(sl)

	// 第２引数が要素数、第３引数が容量
	sL := make([]int,10,20)

	// 要素数を出力
	fmt.Println(len(sL))

	// 容量を出力
	fmt.Println(cap(sL))

	a := []int{1,2,3,4,5}
	b := make([]int,5,20)
	fmt.Println(b)

	// copy関数(第一引数はコピー先、第２引数はコピー元）
	n := copy(b,a)
	fmt.Println(n, b)
}