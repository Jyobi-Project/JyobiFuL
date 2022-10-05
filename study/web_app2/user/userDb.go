package user

import (
	"Jyobi-Project/study/web_app2/db/getConnect"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func SqlInsert(userData UserData) bool {
	db, err := getConnect.SqlConnect()
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return false
	} else {
		fmt.Println("DBアクセス成功")
		fmt.Println(userData)
		result := db.Table("users").Select("name", "age", "address").Create(&userData)
		print(result)
		if result.Error != nil {
			return false
		} else {
			return true
		}
	}
}
