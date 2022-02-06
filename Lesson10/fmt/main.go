package main

import (
	"fmt"
	"os"
)

// fmt

func main() {
	// 標準出力
	fmt.Print("標準")
	// 改行
	fmt.Print("標準\n")
	// 改行
	fmt.Println("改行")

	// Println系
	fmt.Println("aaaa", "bbb", "cccc")

	// Printf系
	fmt.Printf("%s\n", "aaaaa")
	// 値の文字型をそのまま出力
	fmt.Printf("%#v\n", "aaaaa")
	
	// Sprint系（フォーマットした結果を文字列で返す）
	s := fmt.Sprint("sss")
	t := fmt.Sprintf("%v\n", "ttt")
	u := fmt.Sprintln("uuu")
	
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println(u)

	// Fprint系（書き込み先を指定）
	fmt.Fprint(os.Stdout, "aaaaa")
	fmt.Fprintf(os.Stdout, "bbb")
	fmt.Fprintln(os.Stdout, "cccccc")

	f, _ := os.Create("test.txt")
	defer f.Close()

	fmt.Fprintln(f, "Fprint aaaa")
	
}