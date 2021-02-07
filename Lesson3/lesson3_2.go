package main

import "fmt"

func main() {
	var s string = "Hello Go"
	fmt.Println(s)
	fmt.Println(string(s[0]))
	fmt.Println(string(s[7]))

	// 型を調べる
	fmt.Printf("%T\n", s)

	var a string = "500"
	fmt.Println(a)

	// 型を調べる
	fmt.Printf("%T\n", a)

	// 複数行の文字列を表示させる（バッククォーテーションを使う）
	fmt.Println(`s
	s
		 s
				   s`)

	// ダブルクォーテーションを文字列として使用する
	fmt.Println("\"", "\"")
	fmt.Println(`"`)

	// byte（uint8）型
	byteA := []byte{72, 73}
	fmt.Println(byteA)
	// 文字列へ変換
	fmt.Println(string(byteA))

	// バイトの配列に変換して表示
	c := []byte("Hello")
	fmt.Println(c)
	// 文字列へ変換
	fmt.Println(string(c))
}