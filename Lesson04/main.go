package main

import (
	"fmt"
	"strconv"
)

func Sum(a, b int) int {
	result := a * b
	return result
}

func main() {

	s := Sum(10, 10)

	var score int = 100
	var message string = ""

	if s >= score {
		message = "Success!"
	} else if s >= 80 {
		message = "Almost!"
	} else {
		message = "False…"
	}

	fmt.Println(message)

	var number string = "aaa"
	i, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%T\n", i)

	i := 0
	for {
		i++
		if i == 15 {
			break
		} else if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// 条件付きfor文
	for i < 10 {
		fmt.Println(i)
		i++
	}

	arr := [...]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		if i == 3 {
			continue
		}
		fmt.Println(arr[i])
	}

	for i, v := range arr {
		if i == 2 {
			continue
		}
		fmt.Println(i, v)
	}

	m := map[string]int{
		"aaa": 10,
		"bbb": 20,
		"ccc": 30,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
