package main

import "fmt"

type T struct {
	/*
	User型をフィールドとして埋め込む
	フィールド名 型名
	型名を省略することができる
	*/
	// User User
	User
}

type User struct {
	Name string
	Age int
}

// 構造体のスライス
type UserSlc []*User

func (u *User) SetName() {
	u.Name = "QWE"
}

// コンストラクタ型関数
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

// 独自型
type MyInt int

func (i MyInt) Print() {
	fmt.Println(i)
}

func main() {
	t := T{User: User{Name: "aaa", Age: 100}}
	fmt.Println(t)

	// Userのフィールドにアクセス
	fmt.Println(t.User)
	// Nameフィールドにアクセス
	fmt.Println(t.User.Name)
	// 型を省略した場合のフィールドアクセス
	fmt.Println(t.Name)

	t.User.SetName()
	fmt.Println(t)

	n := NewUser("taro", 24)
	fmt.Println(n)

	// ポインタを指定し、実態へアクセス
	fmt.Println(*n)

	a := User{Name: "Jiro", Age: 10}
	b := User{Name: "Saburo", Age: 15}
	c := User{Name: "Shiro", Age: 20}
	d := User{Name: "Giro", Age: 25}
	e := User{Name: "Rokuro", Age: 30}
	
	// UserSlcの変数を定義
	slc := UserSlc{}

	// スライスへの格納
	slc = append(slc, 
		&a,
		&b,
		&c,
		&d,
		&e,
	)

	// 各データのアドレスが出力される
	fmt.Println(slc)

	for _, s := range slc {
		fmt.Println(*s)
	}

	// 構造体のスライスを生成
	m := make([]*User, 0)
	m = append(m, &a, &b, &c, &d, &e)
	
	for _, s := range m {
		fmt.Println(*s)
	}
	
	// 構造体をマップで表示
	x := map[int]User {
		1: {Name: "aaa", Age: 10},
		2: {Name: "bbb", Age: 20},
		3: {Name: "ccc", Age: 30},
		4: {Name: "eee", Age: 40},
		5: {Name: "ddd", Age: 50},
	}
	
	fmt.Println(x)
	
	// キーとしてUserを使用
	y := map[User]string {
		{Name: "Taro", Age: 23}: "Tanaka",
		{Name: "Jiro", Age: 34}: "Takahashi",
	}
	
	fmt.Println(y)
	
	// mapをmake関数で生成
	z := make(map[int]User)
	// 値の格納
	z[0] = User{Name: "Kazuko", Age: 23}
	fmt.Println(z)

	var i  MyInt
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	// メソッドの呼び出し
	i.Print()
}