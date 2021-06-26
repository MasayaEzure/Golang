package main

import (
	"fmt"
	// "time"
)

// チャンネル
// 複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造

func main() {
	// チャンネルの宣言
	// var a chan int

	// 受信専用のチャンネル（明示的な宣言）
	// var b <-chan int

	// 送信専用
	// var c chan<- int

	// チャンネルの初期化
	// a = make(chan int)
	// b := make(chan int)

	// 容量を調べる
	// fmt.Println(cap(a))
	// fmt.Println(cap(b))

	// 第２引数にバッファサイズを指定することができる
	// c := make(chan int, 10)
	// バッファサイズを調べる
	// fmt.Println(cap(c))

	// データを送信する
	// c <- 10
	// fmt.Println(len(c))

	// c <- 20
	// c <- 100
	// fmt.Println(len(c))

	//  データを受信する
	// i := <-c
	// fmt.Println(i)

	// j := <-c
	// fmt.Println(j)

	// k := <-c
	// fmt.Println(k)

	sample := make(chan int,3)
	// Sample := make(chan int)
	sample <- 1
	sample <- 2
	sample <- 3
	// for文を使う場合、チャンネルはクローズ処理をする
	close(sample)
	
	for i := range sample {
		fmt.Println(i)
	}

	// go receive(sample)
	// go receive(Sample)

	// クローズ処理
	// 閉じたチャンネルにデータの送信はできない
	// データの受信は可能
	// close(sample)
	// fmt.Println(<-sample)

	// 第２引数はチャンネルの開閉状態を表す
	/* i,j := <-sample
	fmt.Println(i,j) */

	// クローズ状態のチャンネルで処理が実行される
	/*k,l := <-sample
	fmt.Println(k,l)*/

	// d := 0
	// for d < 100 {
	// 	sample <- d
	// 	Sample <- d
	// 	time.Sleep(50 * time.Millisecond)
	// 	d++
	// }

	// 並行処理
	go sub_receive("Taro",sample)
	go sub_receive("Jiro",sample)

	/*i := 0
	for i < 100 {
		sample <- i
		i++
	}
	close(sample)
	// 処理終了後、三秒待つ
	time.Sleep(3 * time.Second)
	*/
}

func receive(c chan int) {
	for {
		// チャンネルから値を受信
		i := <-c
		fmt.Println(i)
	}
}

func sub_receive(name string, ch chan int) {
	for {
		a,b := <-ch
		// クローズされていないかをチェック
		if !b {
			break
		}
		fmt.Println(name,a)
	}

	fmt.Println(name + " " + "end")
}