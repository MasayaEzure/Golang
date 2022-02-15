package main

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"net/url"
)

// net/http
// クライアント

func main() {
	/*
	// GETメソッド
	res, _ := http.Get("https://example.com")
	fmt.Println(res.StatusCode)
	fmt.Println(res.Proto)

	fmt.Println(res.Header["Date"])
	fmt.Println(res.Header["Context-Type"])

	fmt.Println(res.Request.Method)
	fmt.Println(res.Request.URL)

	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(b))
	*/

	// POSTメソッド
	v := url.Values{}

	v.Add("id", "1")
	v.Add("message", "あああああ")
	fmt.Println(v.Encode())

	res, err := http.PostForm("http://example.com/", v)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
	
}