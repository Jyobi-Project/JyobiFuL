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
		return false
	} else {
		fmt.Println("DBアクセス成功")
		result := db.Table("users").Select("user_name", "user_mail", "user_password", "user_icon").Create(&userData)
		if result.Error != nil {
			return false
		} else {
			return true
		}
	}
}

func SqlSelectUser(mail string) (bool, string) {
	db, err := connect.SqlSelectUser()
	var hashPw UserData
	if err != nil {
		fmt.Println("DBアクセス失敗")
		return false, ""
	} else {
		fmt.Println("DBアクセス成功")
		result := db.Table("users").Select("user_password").Find(&hashPw).Where("user_mail = ?", mail).Limit(1)
		if result.Error != nil {
			return false, ""
		} else {
			return true, hashPw.UserPassword
		}
	}
}
