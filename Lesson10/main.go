package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// メイン関数の処理が終わるタイミングで実行される
	os.Exit(1)
	fmt.Println("aaaa")

	// deferの処理も実行されない
	defer func() {
		fmt.Println("defer")
	}()
	// 実行されない
	os.Exit(0)

	_, err := os.Open("text.txt")
	if err != nil {
		// エラーを出力しながら処理を終了
		log.Fatalln(err)
	}

	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	// os.Argsの要素数を表示
	fmt.Printf("length=%d\n", len(os.Args))
	// os.Argsの中身をすべて表示
	for k, v := range os.Args {
		fmt.Println(k, v)
	}

	// ファイル操作
	f, err := os.Open("abcd.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	// ファイル操作（Create）
	f, _ := os.Create("aaa.txt")

	f.Write([]byte("Hello\n"))
	f.WriteAt([]byte("Golang\n"), 6)
	f.Seek(0, os.SEEK_END)
	f.WriteString("Yeah")

	// ファイル操作（Read）
	f, err := os.Open("aaa.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)

	n, err := f.Read(bs)
	// 読み込んだバイト数を出力
	fmt.Println(n)
	// スライスに読み込んだ内容を文字列に変換して出力
	fmt.Println(string(bs))

	bsOther := make([]byte, 128)

	m, err := f.ReadAt(bsOther, 10)
	fmt.Println(m)
	fmt.Println(string(m))

	// ファイル操作（OpenFile）
	// O_RDONLY：読み込み専用
	// O_WRONLY：書き込み専用
	// O_RDWR：読み書き可能
	// O_APPEND：ファイル末尾に追記
	// O_CREATE：ファイル作成
	// O_TRUNK：可能であればファイルの内容をオープン時に空にする
	f, err := os.OpenFile("aaa.txt", os.O_RDWR | os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	xyz := make([]byte, 128)
	n, err := f.Read(xyz)
	fmt.Println(n)
	fmt.Println(string(xyz))

}