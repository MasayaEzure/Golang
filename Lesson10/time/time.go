package main

import (
	"fmt"
	"time"
)

func main() {
	// 現在時刻
	t := time.Now()
	fmt.Println(t)

	// 指定の時刻を生成
	c := time.Date(2022, 02, 05, 16, 25, 0, 0, time.Local)
	fmt.Println(c)

	// 年
	fmt.Println(c.Year())
	// 年間通算日
	fmt.Println(c.YearDay())
	// 月
	fmt.Println(c.Month())
	// 曜日
	fmt.Println(c.Weekday())
	// 日
	fmt.Println(c.Day())
	// 時間
	fmt.Println(c.Hour())
	// 分
	fmt.Println(c.Minute())
	// 秒
	fmt.Println(c.Second())
	// ナノ秒
	fmt.Println(c.Nanosecond())
	// 場所
	fmt.Println(c.Zone())


	// 時間の感覚を表現
	// 1h
	fmt.Println(time.Hour)
	// 1hの型を示している
	fmt.Printf("%T\n", time.Hour)
	// 1m
	fmt.Println(time.Minute)
	// 1s
	fmt.Println(time.Second)
	// 1ms
	fmt.Println(time.Millisecond)
	// 1µs
	fmt.Println(time.Microsecond)
	// 1ns
	fmt.Println(time.Nanosecond)

	d, _ := time.ParseDuration("2.5h")
	// 2h30m0s
	fmt.Println(d)

	t := time.Now()
	// 現在時刻に2分15秒時間を追加
	t = t.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t)

	tt := time.Date(2021, 07, 24, 0, 0, 0, 0, time.Local)
	tn := time.Now()
	fmt.Println(tn)
	
	// 時刻の差分を取得
	diff := tt.Sub(tn)
	fmt.Println(diff)
	
	// 時間の比較
	fmt.Println(tn.Before(tt))
	fmt.Println(tn.After(tt))
	fmt.Println(tt.Before(tn))
	fmt.Println(tt.After(tn))

	// ５秒間処理を停止させる
	time.Sleep(5 * time.Second)
	fmt.Println("stop")


}