package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type Person struct {
	Name string
	Age int
}

func main() {
	db, _ := sql.Open("sqlite3", "./example.sql")
	
	defer db.Close()
	
	// テーブル作成
	cmd := `CREATE TABLE IF NOT EXISTS persons(
				name STRING,
				age INT)`

	_, err := db.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// データの追加
	insert := `INSERT INTO persons (name, age) VALUES (?, ?)`
	_, err := db.Exec(insert, "Jiro", 30)
	if err != nil {
		log.Fatalln(err)
	}

	// データの更新
	update := `UPDATE persons SET age = ? WHERE name = ?`
	_, err := db.Exec(update, 26, "Taro")
	if err != nil {
		log.Fatalln(err)
	}

	// データの取得
	s := "SELECT * FROM persons WHERE age = ?"
	// QueryRowでレコードを取得
	r := db.QueryRow(s, 30)
	var p Person
	err := r.Scan(&p.Name, &p.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalln("No Row")
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	// 複数データの取得
	m := "SELECT * FROM persons"
	// Queryで条件に合致するすべての情報を取得
	rows, _ := db.Query(m)
	defer rows.Close()

	// スライスを生成
	var mp []Person
	for rows.Next() {
		var p2 Person
		err := rows.Scan(&p2.Name, &p2.Age)
		if err != nil {
			log.Println(err)
		}
		mp = append(mp, p2)

	}

	for _, pe := range mp {
		fmt.Println(pe.Name, pe.Age)
	}

	// データの削除
	delete := "DELETE FROM persons WHERE name = ?"
	_, err := db.Exec(delete, "Jiro")
	if err != nil {
		log.Fatalln(err)
	}
}
