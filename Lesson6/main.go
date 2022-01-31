package main

import "fmt"

// 構造体
type User struct {
	// フィールド
	Name string
	Age int
}

func UpdateUser(user User) {
	user.Name = "AAA"
	user.Age = 100
}

// ポインタ型で引数を取る
func OtherUpdateUser(user *User) {
	user.Name = "AAA"
	user.Age = 100
}

func (u User) sayName() {
	fmt.Println(u.Name, "from sayName")
}

func (u *User) SetName(name string) {
	u.Name = name
}

func main() {
	// 明示的な変数定義
	var user1 &User
	// フィールドの更新
	user1.Age = 25

	fmt.Println(user1)

	// 暗黙的な変数定義
	user2 := User{}
	user2.Name = "Jiro"
	user2.Age = 30

	fmt.Println(user2)

	// 初期値を予め設定した場合の宣言
	user3 := User{"Saburo", 50}

	// メソッドの呼び出し
	user3.SetName("aaa") 
	user3.sayName()

	user := &User{Name: "waooo"}
	user.SetName("yeaaa")
	user.sayName()

	fmt.Println(user3)

	// 一つの初期値を設定した場合の宣言
	user4 := User{Name: "Siro"}
	fmt.Println(user4)

	// 構造体の変数宣言時に関数を使う場合
	user5 := new(User)
	fmt.Println(user5)

	// アドレス演算子をつけた場合の宣言
	user6 := &User{}
	fmt.Println(user6)

	UpdateUser(user1)
	OtherUpdateUser(user6)

	fmt.Println(user1)
	fmt.Println(user6)
}