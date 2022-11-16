package users

import (
	"Jyobi-Project/study/db/getConnect"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func SqlCreateAccount(userData UserData) bool {
	db, err := getConnect.SqlConnect()
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return false
	} else {
		fmt.Println("DBアクセス成功")
		fmt.Println(userData)
		result := db.Table("users").Select("user_name", "user_mail", "user_password", "salt", "user_icon").Create(&userData)
		print(result)
		if result.Error != nil {
			return false
		} else {
			return true
		}
	}
}
