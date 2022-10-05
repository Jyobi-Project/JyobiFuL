package User

import (
	"Jyobi-Project/study/web_app2/db/getConnect"
	"Jyobi-Project/study/web_app2/user"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func SqlInsert(userData *user.UserData) bool {
	db, err := getConnect.SqlConnect()
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
		return false
	} else {
		fmt.Println("DBアクセス成功")
		fmt.Println(db.Select("name", "age", "address").Create(userData))
		if err != nil {
			return false
		} else {

		}
		return true
	}
}
