package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Account struct {
	ID string
	PW string
}

func main() {
	// ヘッダー、クエリの追加
	// Parse：正しいURLなのかを確認する
	base, _ := url.Parse("http://example.com/")

	// クエリの確認
	ref, _ := url.Parse("index/lists?id=1")

	// ResolveReference：クエリを追加したURLの生成
	// 相対パスから絶対パスに変換
	endpoint := base.ResolveReference(ref).String()
	fmt.Println(endpoint)

	// Get
	// リクエストの生成（リクエストはしておらず、structを作っただけの状態）
	req, _ := http.NewRequest("GET", endpoint, nil)

	// reqにヘッダーを付与
	req.Header.Add("Content-Type", `application/json`)

	// URLのクエリを確認
	q := req.URL.Query()

	// クエリの追加
	q.Add("name", "test")
	// 特殊文字に対応するため、encordingする
	fmt.Println(q.Encode())

	// クライアントを作成
	var client *http.Client = &http.Client{}

	// 実際にアクセスする
	resp, _ := client.Do(req)

	// 読み込み
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)

	// Post
	base2, _ := url.Parse("https://example.com/")
	ref2, _ := url.Parse("index/post?id=2")

	// POST時のデータ
	AccountData := &Account{ID :"123", PW: "password123"}
	d, _ := json.Marshal(AccountData)

	// ResolveReference
	ep := base2.ResolveReference(ref2).String()
	fmt.Println(ep)

	// bytes.NewBuffer([]byte("password"))：リクエストの領域を作成
	req2, _ := http.NewRequest("POST", ep, bytes.NewBuffer(d))

	// req2にヘッダーを付与
	req2.Header.Add("Content-Type", `application/json`)

	// URLのクエリを確認
	q2 := req2.URL.Query()

	// クエリを追加
	q2.Add("name", "aaaa")
	fmt.Println(q2.Encode())

	// URLに戻す（日本語などを変換する）
	req2.URL.RawQuery = q2.Encode()

	// クライアントを作成
	var client2 *http.Client = &http.Client{}

	// 実際にアクセスする
	resp2, _ := client2.Do(req2)

	// 読み込み
	body2, _ := ioutil.ReadAll(resp2.Body)
	fmt.Println(body2)
}