package getConnect

package getConnect

import (
_ "github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
"github.com/joho/godotenv"
"os"
)

// SqlConnect DBのconnectionを取得
// DBの情報は基本的にenvから取り出すこと
func SqlConnect() (database *gorm.DB, err error) {

	err = godotenv.Load(".env")

	// envファイルが存在している場合
	if err == nil {
		//DB接続に必要な情報を取得
		DBMS := "mysql"
		USER := os.Getenv("DB_USER")
		PASS := os.Getenv("DB_PASS") //自分の設定したパスワード
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
