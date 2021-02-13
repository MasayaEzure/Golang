package main

import "fmt"
import "os"

func main() {
	// 無名関数
	// main関数が処理終了後に呼び出される
	defer func() {
		num := [...]int{1, 2, 3, 4, 5, 6}

		for i := 0; i < len(num); i++ {
			fmt.Println(num)
		}
	}()

	// ラベル付きのfor
Fin:

	for i := 0; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			if j > 1 {
				continue Fin
			}
			fmt.Println(i, j)
		}
		fmt.Println("DID")
	}

	Testdefer()
	Rundeffer()

	// test.txtを生成
	file,err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// ファイルを閉じる（最後に実行）
	defer file.Close()
	// ファイルに書き込む処理
	file.Write([]byte("Goodbye"))
}

func Testdefer() {
	a, b := "Hello", "GOlang"

	// Testdefer()関数が終了したときに呼び出される
	defer fmt.Println(a)
	fmt.Println(b)
}

// 関数の処理終了時、後から登録されたものから順に処理実行となる
func Rundeffer() {
	Num := [...]int{1, 2, 3}

	for j := 0; j < len(Num); j++ {
		defer fmt.Println(Num[0])
		defer fmt.Println(Num[1])
		defer fmt.Println(Num[2])
	}
}
