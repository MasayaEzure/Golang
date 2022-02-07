package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ioutil

func main() {
	// 入力全体を読み込む
	f, _ := os.Open("aaa.txt")
	g, _ := ioutil.ReadAll(f)
	fmt.Println(string(g))

	// ファイルの書き込み
	if err := ioutil.WriteFile("bbb.txt", g, 0666); err != nil {
		log.Fatalln(err)
	}
}