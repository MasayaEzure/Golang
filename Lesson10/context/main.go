package main

import (
	"context"
	"fmt"
	"time"
)

// context

// 処理に時間がかかる関数を定義(2秒間)
func longProcess(c context.Context, ch chan string) {
	fmt.Println("start")
	time.Sleep(2 * time.Second)
	fmt.Println("end")
	ch <- "result"
}

func main() {
	ch := make(chan string)
	// コンテキストの作成
	c := context.Background()
	c, cancel := context.WithTimeout(c, 1*time.Second)
	defer cancel()
	go longProcess(c, ch)

L:
	for {
		select {
		case <- c.Done():
			fmt.Println("error")
			fmt.Println(c.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}
	}

	fmt.Println("break")
}