package main

import (
	"fmt"
	"os"
	"time"
)

func sample(a interface{}) {
	var message string = ""

	// 型のswitch
	switch v := a.(type) {
	case int:
		fmt.Println(v)
		message = "int!"
	case float64:
		fmt.Println(v)
		message = "float64#"
	case string:
		fmt.Println(v)
		message = "string$"
	default:
		fmt.Println("???")
		message = "none…"
	}

	fmt.Println(message)
}

func Test() {
	// Test関数が終了したときに実行される
	defer fmt.Println("-- END --")
	fmt.Println("-- START --")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}

func Sub() {
	for {
		fmt.Println("Sub loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	// switch文
	switch n := 2; n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("none")
	}

	sample("aaaa")
	sample(100)
	sample(3.1)
	sample(true)

	// ラベル付きのFor文
	Loop:
	for {
		for {
			for {
				fmt.Println("-- START --")
				break Loop
			}
			fmt.Println("aaaaa")
		}
		fmt.Println("aaaaa")
	}
	fmt.Println("-- END --")

	// main関数が終了したときに実行される無名関数
	defer func() {
		fmt.Println("This")
		fmt.Println("is")
		fmt.Println("defer")
	}()

	Test()
	RunDefer()

	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	f.Write([]byte("Hello"))

	// 例外処理（あまり使用されない）
	defer func() {
		// recover：panicによって発生したエラーから復帰するための処理
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	// panic(処理の強制終了)
	panic("Runtime Error")
	fmt.Println("aaa")

	// メイン関数のループ処理と並行で実行される （goroutine）
	go Sub()

	for {
		fmt.Println("Main loop")
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Println("This is main")

}

/*
init
メイン関数より先に実行される特別な関数
複数定義も可能だが、あまり望ましくない
*/
func init() {
	fmt.Println("This is init")
}
