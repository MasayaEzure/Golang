package main

import "fmt"

// 例外処理
// 並行処理

func main() {
	// main関数の処理終了時に走る関数
	defer func() {
		// パニック状態であればrecover()で修復される
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	// 処理を終了させる
	panic("error")
	fmt.Println("start")

	
}