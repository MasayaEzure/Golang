package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	// 文字列の結合
	a := strings.Join([]string{"a","b","c"}, ",")
	b := strings.Join([]string{"a","b","c"}, "")
	fmt.Println(a, b)
	
	// 文字列に含まれる部分文字列の検索
	c := strings.Index("cdjkfhakfj", "a")
	fmt.Println(c)
	
	d := strings.LastIndex("cdjkfhakfj", "fj")
	fmt.Println(d)
	
	e := strings.IndexAny("cdjkfhakfj", "cd")
	f := strings.LastIndexAny("cdjkfhakfj", "cdjkfhakfj")
	fmt.Println(e, f)
	
	g := strings.HasPrefix("ABC", "A")
	h := strings.HasSuffix("ABC", "B")
	fmt.Println(g, h)
	
	i := strings.Contains("afn;sjdhfka", "a")
	j := strings.ContainsAny("afn;sjdhfka", "jh")
	fmt.Println(i, j)
	
	k := strings.Count("abcdef", "d")
	// 文字列の長さ + 1 の値が返ってくる
	l := strings.Count("abcdef", "")
	fmt.Println(k, l)
	
	// 文字列を繰り返し結合
	m := strings.Repeat("abcd", 8)
	// 出力されない
	n := strings.Repeat("abcd", 0)
	fmt.Println(m)
	fmt.Println(n)
	
	// 文字列の置換
	o := strings.Replace("xyzxyzxyz", "x", "a", 1)
	// 該当箇所をすべて変換
	p := strings.Replace("xyzxyzxyz", "x", "a", -1)
	fmt.Println(o, p)
	
	// 文字列の分割
	q := strings.Split("a,b,c,d,e", ",")
	r := strings.SplitAfter("a,b,c,d,e", ",")
	s := strings.SplitN("a,b,c,d,e", ",", 2)
	t := strings.SplitAfterN("a,b,c,d,e", ",", 5)
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
	fmt.Println(t)
	
	// 大文字の変換
	u := strings.ToLower("AAA")
	// 小文字の変換
	v := strings.ToUpper("bbb")
	fmt.Println(u, v)
	
	// 文字列から空白を排除
	w := strings.TrimSpace("     - aaa -    ")
	// 全角
	x := strings.TrimSpace("　　　　bbb　　")
	fmt.Println(w, x)
	*/

	// 文字列からスペースで区切られれたフィールドの取り出し
	y := strings.Fields("a b c d")
	fmt.Println(y)
	fmt.Println(y[3])
	
}