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

func sqlAllSelect() ([]GetSelectUserData, bool) {
	db, err := getConnect.SqlConnect()
	var users []GetSelectUserData
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return nil, false
	} else {
		fmt.Println("DBアクセス成功")
		db.Table("users").Find(&users)
		return users, true
	}
}

func sqlWhereSelect(id int) (GetSelectUserData, bool) {
	db, err := getConnect.SqlConnect()
	var user GetSelectUserData
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return user, false
	} else {
		fmt.Println("DBアクセス成功")
		// sqlインジェクションの対策ができないため、下記はだめ
		//db.Table("users").First(&users, id)
		db.Table("users").First(&user, "id = ?", id)
		return user, true
	}
}
