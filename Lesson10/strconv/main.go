package main

import (
	"fmt"
	"strconv"
)

// strconv

func main() {

	// 真偽値型を文字列に変換
	bt := true
	fmt.Printf("%T\n", strconv.FormatBool(bt))
	
	// 整数を文字列に変換(第2引数は進数を表している)
	i := strconv.FormatInt(-100, 10)
	fmt.Printf("%v, %T\n", i , i)
	
	// Itoaで簡易的に変換
	j := strconv.Itoa(1234)
	fmt.Printf("%v, %T\n", j, j)
	
	// 浮動小数点に変換
	fmt.Println(strconv.FormatFloat(12.345, 'E', -1, 64))
	
	// 指数表現による文字列化（小数点以下2桁まで）
	fmt.Println(strconv.FormatFloat(12.345, 'e', 2, 64))
	
	// 実数表現による文字列化
	fmt.Println(strconv.FormatFloat(12.345, 'f', -1, 64))
	
	// 実数表現による文字列化（小数点以下2桁まで）
	fmt.Println(strconv.FormatFloat(12.345, 'f', 2, 64))
	
	// 指数部の大きさで変動する表現による文字列化
	fmt.Println(strconv.FormatFloat(12.345, 'g', -1, 64))
	fmt.Println(strconv.FormatFloat(123456789.123, 'f', -1, 64))
	
	// 指数部の大きさで変動する表現による文字列化（仮数部全体が4桁まで）
	fmt.Println(strconv.FormatFloat(123.456, 'g', 4, 64))
	
	// 指数部の大きさで変動する表現による文字列化（仮数部全体が8桁まで）
	fmt.Println(strconv.FormatFloat(123456789.123, 'G', 8, 64))
	
	// 文字列を真偽値に変換
	a, _ := strconv.ParseBool("true")
	fmt.Printf("%v, %T\n", a, a)
	b, _ := strconv.ParseBool("1")
	c, _ := strconv.ParseBool("t")
	d, _ := strconv.ParseBool("T")
	e, _ := strconv.ParseBool("TRUE")
	f, _ := strconv.ParseBool("True")
	fmt.Println(b, c, d, e, f)
	
	// falseに変換できる文字列
	bf, ok := strconv.ParseBool("false")
	// okの戻り値はerror型のため、エラーハンドリングも可能
	if ok != nil {
		fmt.Println("Error…")
	}
	fmt.Printf("%v, %T\n", bf, bf)
	g, ok := strconv.ParseBool("0")
	h, ok := strconv.ParseBool("f")
	i, ok := strconv.ParseBool("F")
	j, ok := strconv.ParseBool("FALSE")
	k, ok := strconv.ParseBool("False")
	fmt.Println(g, h, i, j, k)
	
	// 文字列を整数型に変換
	aa, _ := strconv.ParseInt("12345", 10, 0)
	fmt.Printf("%v, %T\n", aa, aa)
	
	// Atoiで簡易的に変換
	bb, _ := strconv.Atoi("-123")
	fmt.Printf("%v, %T\n", bb, bb)

	// 文字列を浮動小数点型に変換
	fla, _ := strconv.ParseFloat("3.14", 64)
	flb, _ := strconv.ParseFloat(".5", 64)
	flc, _ := strconv.ParseFloat("-.2", 64)
	fld, _ := strconv.ParseFloat("1.2345e8", 64)
	fle, _ := strconv.ParseFloat("1.2345E8", 64)
	fmt.Println(fla, flb, flc, fld, fle)
}
