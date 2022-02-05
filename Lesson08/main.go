package main

import (
	// 任意のパッケージ名を付与
	"Lesson8/sample"
	"fmt"
	f "fmt"
)

func appName() string {
	const AppName = "aaaaa"
	var Ver string = "1.0"
	return AppName + " " + Ver
}

func Sample(a string) (b string) {
	b = a
	return b
}

func main() {
	f.Println(sample.Max)
	f.Println(sample.ReturnMin())

	fmt.Println(appName())
	fmt.Println(Sample("xxxx"))
}