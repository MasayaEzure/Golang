package main

import (
	// "fmt"
	"log"
	"net/http"
	"text/template"
)

/*
type MyHandler struct {}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
*/

func top(w http.ResponseWriter, r * http.Request) {
	// templateの構造体を生成
	t, err := template.ParseFiles("tmpl.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, "AAAAaaaa")
}

func main() {
	// 第一引数：URL、第２引数：関数
	http.HandleFunc("/top", top)

	// サーバーを作成
	// 第一引数：ネットワークアドレス、第二引数：Handlerを指定
	http.ListenAndServe(":8080", nil)
}