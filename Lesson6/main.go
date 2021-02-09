package main

import "fmt"

// 関数
func Avg(i, j, k int)(sum int) {
	sum = i + j + k
	return
}

// 関数を返す関数
func RuturnFunc() func(){
	return func() {
		fmt.Println("関数を返してるお")
	}
}

// 関数を引数に取る関数（実際に呼び出してるのはこっち）
func CallFunc(f func()) {
	f()
}

// クロージャーの実装
// 「”引数に文字列を受け取って文字列を返す”という関数」を返す関数
func Later() func(string) string {
	// 変数を定義（空文字）
	var store string

	// 無名関数内で"next"の引数を受け取る
	return func(next string) string {

	// storeをsという変数にコピーする
		s := store

	// sをnext（引数）に代入
		store = next
		
		return s
	}
}

// ジェネレーターの実装
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	// インスタンス化して出力
	sample := Avg(50, 60, 70)
	fmt.Println(sample/3)

	// 無名関数
	b := func(x, y int) int {
		return x + y
	}(5, 7)

	fmt.Println(b)

	c := RuturnFunc()
	c()

	CallFunc(func() {
		fmt.Println("引数の関数だお")
	})

	z := Later()
	fmt.Println(z("Hello"))  // この段階では空文字が出力される（storeの初期値は空文字のため）
	fmt.Println(z("My"))	// ここから"Hello"が出力される
	fmt.Println(z("World"))
	fmt.Println(z("Golang"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
}