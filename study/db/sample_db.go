package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"

	"Jyobi-Project/study/db/getConnect"
)

func main() {
	_, err := getConnect.SqlConnect()
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	} else {
		fmt.Println("DBアクセス成功")
	}

}
