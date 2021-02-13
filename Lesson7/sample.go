package main

import "fmt"

func any(a interface{}) {
	switch v := a.(type) {
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	case float64:
		fmt.Println(v, "float64")
	}
}

func main() {
	any("Hello")
	any(3.5)

	// 型スイッチ
	// interface型の変数を定義
	var x interface{} = 3

	// int型に復元して、変数iに代入
	i,j := x.(int)
	fmt.Println(i,j)

	// int型に復元したものを別の方に変換することはできない
	// ２つ目の変数は変換に失敗する = falseを返す
	// f, g := x.(float64)
	// fmt.Println(f, g)

	if x == nil {
		fmt.Println("none")
	} else if i,j := x.(int); j {
		fmt.Println(i,j)
	} else if s, t := x.(string); t{
		fmt.Println(s,t)
	} else {
		fmt.Println("Empty")
	}

	// 型によるスイッチ文
	switch x.(type) {
	case int:
		fmt.Println("aaa")
	case string:
		fmt.Println("bbb")
	default:
		fmt.Println("NO")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v,"bool")
	case string:
		fmt.Println(v,"string")
	case int:
		fmt.Println(v,"int")
	}
}