package main

import (
	"flag"
	"fmt"
)

// flag

func main() {

	// コマンドラインのオプション処理
	// go run flag.go -n 20 -m message -x

	// 変数の定義
	var (
		max int
		msg string
		tof bool
	)

	// IntVar 整数のオプション
	flag.IntVar(&max, "n", 32, "最大値")

	// StringVar 文字列のオプション
	flag.StringVar(&msg, "m", "", "メッセージ")

	// BoolVar 真偽型のオプション
	flag.BoolVar(&tof, "tof", false, "拡張オプション")

	// コマンドラインをパース
	flag.Parse()

	fmt.Println("最大値：", max)
	fmt.Println("メッセージ：", msg)
	fmt.Println("オプション：", tof)

}