package main

import (
	"Jyobi-Project/study/db/getConnect"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func sqlInsert(name string, age int, address string) bool {
	db, err := getConnect.SqlConnect()
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return false
	} else {
		fmt.Println("DBアクセス成功")
		sql, err := db.Exec("INSERT INTO users VALUES(NULL, ?, ?, ?)", name, age, address)
		if err != nil {
			return false
		} else {

		}
		return true
	}
}
