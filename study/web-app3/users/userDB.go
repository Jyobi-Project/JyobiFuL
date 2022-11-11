package users

import (
	"Jyobi-Project/study/db/getConnect"
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
		result := db.Table("prac_users").Select("name", "mail", "age", "password").Create(&userData)
		print(result)
		if result.Error != nil {
			return false
		} else {
			return true
		}
	}
}

func SqlAllSelect() ([]UserData, bool) {
	db, err := getConnect.SqlConnect()
	var users []UserData
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return nil, false
	} else {
		fmt.Println("DBアクセス成功")
		db.Table("prac_users").Find(&users)
		return users, true
	}
}

func SqlWhereSelect(mail string) ([]UserData, bool) {
	db, err := getConnect.SqlConnect()
	var users []UserData
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return nil, false
	} else {
		fmt.Println("DBアクセス成功")
		db.Table("prac_users").Where("mail like ?", "%"+mail+"%").Find(&users)
		return users, true
	}
}

func SqlUpdateUser(userData UserData) bool {
	db, err := getConnect.SqlConnect()
	var users []UserData
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return false
	} else {
		fmt.Println("DBアクセス成功")
		db.Table("prac_users").Where("mail like ?", "%"+mail+"%").Update("name", "hello").Find(&users)
		return true
	}
}
