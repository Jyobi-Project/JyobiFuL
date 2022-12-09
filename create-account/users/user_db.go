package users

import (
	"Jyobi-Project/db"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func SqlCreateAccount(userData UserData) bool {
	db, err := connect.SqlInsertUser()
	if err != nil {
		fmt.Println("DBアクセス失敗")
		fmt.Println("送信されたデータの内容 -> ", userData)
		fmt.Println(err)
		return false
	} else {
		fmt.Println("DBアクセス成功")
		fmt.Println(userData)
		result := db.Table("users").Select("user_name", "user_mail", "user_password", "user_icon").Create(&userData)
		print(result)
		if result.Error != nil {
			return false
		} else {
			return true
		}
	}
}
