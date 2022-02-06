package main

import (
	// "fmt"
	"log"
	"os"
)

// log

func main(){
	// ログの出力先を変更
	log.SetOutput(os.Stdout)
	
	log.Print("aaaa\n")
	log.Println("bbb")
	log.Printf("ccc%d\n", 3)
	
	// Fatal系
	log.Fatal("aaa\n")
	log.Fatalln("bbb")
	log.Fatalf("ccc%d\n", 3)
	
	// Panic系
	log.Panic("aaa\n")
	log.Panicln("bbb")
	log.Panicf("ccc%d\n", 3)
	
	// 任意のファイルを作成し、出力先に指定
	f, err := os.Create("test.log")
	if err != nil {
		return
	}
	
	// 作成したファイルを出力先に指定
	log.SetOutput(f)
	log.Println("ファイルに書き込む")
	
	// ログのフォーマットを指定
	log.SetFlags(log.LstdFlags)
	log.Println("aaa")
	
	// マイクロ秒を追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("bbb")
	
	// 時刻とファイルの行番号
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("cccc")
	
	log.SetFlags(log.LstdFlags)
	// ログのプリフィックスを設定
	log.SetPrefix("[log]")
	log.Println("dddd")

	// ロガーの生成
	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime | log.Lshortfile)
	logger.Println("message")
	log.Println("message")

	_, err := os.Open("xxxaaaa")
	if err != nil {
		logger.Fatalln("Exit", err)
	}

}