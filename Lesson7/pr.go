package main

import "fmt"
import "time"

func sub() {
	for {
		fmt.Println("Sub Loop")
		// 一定時間空ける処理
		time.Sleep(100 * time.Millisecond)
	}
}

// 例外処理

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

	// sub関数とmain関数を並行処理
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}

}

// init関数：main関数よりも先に処理が走る
// 初期化処理に便利
func init() {
	fmt.Println("init")
}