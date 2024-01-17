package main

import (
	"fmt"
	"time"
)

func main() {
	// map（明示的な宣言）

	var m = map[string]int{
		"aaa": 1,
		"bbb": 2,
		"ccc": 3,
	}

	fmt.Println(m)
	fmt.Println(m["aaa"])
	fmt.Println(m["bbb"])
	fmt.Println(m["ccc"])

	n := make(map[int]string)
	n[0] = "naaa"
	n[1] = "nbbb"
	n[2] = "nccc"

	fmt.Println(n)
	fmt.Println(n[0])
	fmt.Println(n[1])
	fmt.Println(n[2])

	i, ok := n[3]
	if !ok {
		fmt.Println("Not found")
	}
	fmt.Println(i, ok)

	delete(n, 0)  // 削除対象のmap、キー
	fmt.Println(n)

	fmt.Println(len(n))

	for k, v := range m {
		fmt.Println(k, v)
	}

	/*
		チャンネル
		複数の goroutine 間でのデータの受け渡しするために設計された構造
		宣言
		var チャンネル名 chan 型
		受信専用のチャンネル宣言
		var チャンネル名 <-chan 型
		送信専用のチャンネル宣言
		var チャンネル名 chan<- 型
	*/
	c := make(chan int)
	fmt.Println(cap(c))

	// 第2引数に容量を指定
	d := make(chan int, 10)
	fmt.Println(cap(d))

	// データを送信
	d <- 1
	d <- 2
	d <- 3
	fmt.Println("len", len(d))

	// データを受信
	fmt.Println("len", len(d))
	fmt.Println(<-d)
	fmt.Println("len", len(d))

	fmt.Println("len", len(d))
	fmt.Println(<-d)
	fmt.Println("len", len(d))

	fmt.Println("len", len(d))
	fmt.Println(<-d)
	fmt.Println("len", len(d))

	// channel & goroutine
	aaa := make(chan int)
	bbb := make(chan int)

	go Reciever(aaa)
	go Reciever(bbb)

	u := 0
	for u < 100 {
		aaa <- u
		bbb <- u
		time.Sleep(50 * time.Millisecond)
		u++
	}

	// close
	cl := make(chan int, 3)
	cl <- 1
	close(cl)

	fmt.Println(<-cl)

	go OhterReciever("1aaa", cl)
	go OhterReciever("2bbb", cl)

	for i := 0; i < 100; i++ {
		cl <- i
	}
	close(cl)
	time.Sleep(3 * time.Second)

	cl <- 1
	cl <- 2
	cl <- 3

	close(cl)

	for a := range cl {
		fmt.Println(a)
	}

	cstr := make(chan string, 3)
	cstr <- "Hello"
	cstr <- "yeah"
	cstr <- "Uhh"

	/*
		select
		チャンネルに対する処理に対してのみ有効
		どのケース式が実行されるかは分からない
	*/
	select {
	case v1 := <-cl:
		fmt.Println(v1 + 100)
	case v2 := <-cstr:
		fmt.Println(v2 + " from cstr")
	default:
		fmt.Println("Not Found...")
	}

	cha := make(chan int)
	chb := make(chan int)
	chc := make(chan int)

	// Reciever
	go func() {
		for {
			a := <-cha
			chb <- a + 1
		}
	}()

	go func() {
		for {
			b := <-chb
			chc <- b * 2
		}
	}()

	f := 0
	for {
		select {
		case cha <- f:
			f++
		case c := <-chc:
			fmt.Println("recieved ", c)
		}
		break
	}
}

// チャンネルにデータが送信されたら実行される関数
func Reciever(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func OhterReciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println("END")
}
