package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 標準入力を行単位で読み込む

	// 標準入力をソースにしたスキャナの生成
	scanner := bufio.NewScanner(os.Stdin)

	// 入力のスキャンが成功する限り繰り返す
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// スキャンでエラーが発生した場合の処理
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みエラー：", err)
	}
	
}