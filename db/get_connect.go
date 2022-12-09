package connect

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"os"
)

func GetEnv() (err error) {
	return godotenv.Load(".env")
}

// SqlConnect DBのconnectionを取得
// DBの情報は基本的にenvから取り出すこと
func SqlInsertUser() (database *gorm.DB, err error) {
	// envファイルが存在している場合
	if GetEnv() == nil {
		//DB接続に必要な情報を取得
		DBMS := "mysql"
		USER := os.Getenv("INSERT_USER")
		PASS := os.Getenv("INSERT_USER_PASS") //自分の設定したパスワード
		PROTOCOL := os.Getenv("DB_PROTOCOL")
		DBNAME := os.Getenv("DB_NAME")

		//DSNの生成
		CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
		return gorm.Open(DBMS, CONNECT) //DBを開く

		//	envファイルが取得出来なかった場合は、nilを返す
	} else {
		return nil, err
	}
}

func SqlSelectUser() (database *gorm.DB, err error) {
	// envファイルが存在している場合
	if GetEnv() == nil {
		//DB接続に必要な情報を取得
		DBMS := "mysql"
		USER := os.Getenv("SELECT_USER")
		PASS := os.Getenv("SELECT_USER_PASS") //自分の設定したパスワード
		PROTOCOL := os.Getenv("DB_PROTOCOL")
		DBNAME := os.Getenv("DB_NAME")

		//DSNの生成
		CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
		return gorm.Open(DBMS, CONNECT) //DBを開く

		//	envファイルが取得出来なかった場合は、nilを返す
	} else {
		return nil, err
	}
}

func SqlUpdateUser() (database *gorm.DB, err error) {
	// envファイルが存在している場合
	if GetEnv() == nil {
		//DB接続に必要な情報を取得
		DBMS := "mysql"
		USER := os.Getenv("UPDATE_USER")
		PASS := os.Getenv("UPDATE_USER_PASS") //自分の設定したパスワード
		PROTOCOL := os.Getenv("DB_PROTOCOL")
		DBNAME := os.Getenv("DB_NAME")

		//DSNの生成
		CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
		return gorm.Open(DBMS, CONNECT) //DBを開く

		//	envファイルが取得出来なかった場合は、nilを返す
	} else {
		return nil, err
	}
}
