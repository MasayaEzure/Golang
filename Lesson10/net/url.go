package main

import (
	"fmt"
	"net/url"
)

// net/url

func main() {
	// URLを更新
	u, _ := url.Parse("https://example.com/search?a=1&b=2#top")
	fmt.Println(u.Scheme)
	fmt.Println(u.Host)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)
	fmt.Println(u.Fragment)

	fmt.Println(u.Query())

	// URLの生成
	url := &url.URL{}
	url.Scheme = "https:"
	url.Host = "aaaa.com"
	q := url.Query()
	q.Set("q", "Golang")
	// 次でエスケープ処理される
	// q.Set("aaa", "GO言語")

	// エスケープ処理を実行
	url.RawQuery = q.Encode()

	fmt.Println(url)
}