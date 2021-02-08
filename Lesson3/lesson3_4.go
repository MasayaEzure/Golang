package main

import "fmt"
// 型変換を実行するパッケージ
import "strconv"

// 型変換

func main() {
	var i int = 1
	j := float64(i)

	fmt.Println(j)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", j)

	// floatからintへの型変換時、小数点は表示されない
	k := 10.5
	l := int(k)
	fmt.Printf("%T\n", l)
	fmt.Println(l)


	// 文字列から数値への型変換
	var a string = "100"
	// _：２つ目の変数を破棄することができる
	b, _ := strconv.Atoi(a)

	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// 数値から文字列への型変換
	var c int = 200
	d := strconv.Itoa(c)

	fmt.Println(d)
	fmt.Printf("%T\n", d)

	// 文字列からバイトへの型変換
	var e string = "Hello"
	f := []byte(e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)

	// バイトから文字列への変換
	g := string(f)
	fmt.Println(g)
	fmt.Printf("%T\n", g)
}