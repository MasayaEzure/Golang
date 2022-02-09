package main

import (
	"fmt"
	"sync"
	"time"
)

// sync

// ミューテックスによる同期処理
var st struct{A, B, C int}

// ミューテックスを保持するパッケージ変数
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
	// ロック
	mutex.Lock()
	
	st.A = n
	time.Sleep(time.Microsecond)
	st.B = n
	time.Sleep(time.Microsecond)
	st.C = n
	fmt.Println(time.Microsecond)
	fmt.Println(st)

	// アンロック
	mutex.Unlock()
}

func main() {
	mutex = new(sync.Mutex)
	
	for i := 0; i < 5; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				UpdateAndPrint(i)
			}
		}()
	}

	for {

	}

	// goroutine の終了を待ち受ける

	a := new(sync.WaitGroup)

	// 待ち受けるgoroutineの数を指定
	a.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("aaaa")
		}
		// 完了
		a.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("bbb")
		}
		// 完了
		a.Done()
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("cccc")
		}
		// 完了
		a.Done()
	}()

	// goroutineの完了を待ち受ける
	// Doneが3つ完了するまで待つ
	a.Wait()
}