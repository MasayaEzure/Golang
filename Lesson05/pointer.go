package main

import "fmt"

func Double(i int) {
	i *= 2
}

func OtherDouble(i *int) {
	*i *= 2
}

func AntherDouble(s []int) {
	for i, v := range s {
		s[i] = v * 2
	}
}

func main() {
	var n int = 20
	fmt.Println(n)

	// メモリ上のアドレスを表示
	fmt.Println(&n)

	// ポインタ
	var p *int = &n  // &n は変数名
	// アドレス
	fmt.Println(p)
	// 値
	fmt.Println(*p)

	*p = 30
	fmt.Println(n)

	n = 50
	fmt.Println(*p)

	OtherDouble(&n)
	fmt.Println(n)

	OtherDouble(p)
	fmt.Println(*p)

	var sl []int = []int{1, 2, 3}

	AntherDouble(sl)
	fmt.Println(sl)
}
