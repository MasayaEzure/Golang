package main

import "fmt"

// インターフェース
type Stringpy interface {
	ToString() string
}

type User struct {
	Name string
	Age int
}

func (u *User) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", u.Name, u.Age) 
}

type Car struct {
	Number int
	Model string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

// カスタムエラーのinterface
type ExceptionError struct {
	// フィールド
	Message string
	ErrCode int
}

func (e *ExceptionError) Error() string {
	return e.Message
}

// エラーを返す関数
func RaiseError() error {
	return &ExceptionError{
		Message: "カスタムエラー",
		ErrCode: 1234,
	}
}

type Point struct {
	a int
	b string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.a, p.b)
}

func main() {
	/*
	i := []Stringpy {
		&User{Name: "Taro", Age: 25},
		&Car{Number: 1234, Model: "ABC"},
	}
	
	for _, v := range i {
		fmt.Println(v.ToString())
	}
	
	err := RaiseError()
	fmt.Println(err.Error())
	
	e, ok := err.(*ExceptionError)
	if ok {
		fmt.Println(e.ErrCode)
	}
	*/

	p := &Point{10, "ABC"}
	fmt.Println(p)

}