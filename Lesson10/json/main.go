package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// json

// 構造体からJSONテキストへの変換
type A struct{}

type User struct {
	ID int `json:"id"`
	// IDを非表示にする
	// ID int `json:"-"`
	Name string `json:"name"`
	// 値を隠す
	// Name string `json:"name,omitempty"`
	Email string `json:"email"`
	Created time.Time `json:"created"`
	A A `json:"A"`
	// 構造体を隠す
	// A *A `json:"A,omitempty"`
}

// Marshalをカスタムする関数
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + u.Name,
	})
	return v, err
}

// Unmarshalをカスタムする関数
func (x *User) UnmarshalJSON(d []byte) error {
	type OtherUser struct {
		Name string
	}
	var c OtherUser
	err := json.Unmarshal(d, &c)
	if err != nil {
		fmt.Println(err)
	}
	x.Name = c.Name + "#!"
	return err
}

func main() {
	a := new(User)
	a.ID = 4
	a.Name = "Shiro"
	a.Email = "taro@example.com"
	a.Created = time.Now()

	// Marshal：JSONに変換
	b, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

	// JSONから構造体への変換
	fmt.Printf("%T\n", b)

	c := new(User)

	// unmarshal：JSONをデータ型に変換
	if err := json.Unmarshal(b, &c); err != nil {
		fmt.Println(err)
	}
	fmt.Println(c)
	
}