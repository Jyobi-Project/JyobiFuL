# Jyobi-Project

## db接続手順書

### 事前準備

- [mysqlをubuntuにインストール](https://learn.microsoft.com/ja-jp/windows/wsl/tutorials/wsl-database "公式ドキュメント")
- [パスワード設定でエラー起きたときの処理](https://exerror.com/failed-error-set-password-has-no-significance-for-user-rootlocalhost-as-the-authentication-method-used-doesnt-store-authentication-data-in-the-mysql-server/ "エラー時の参考")

### goの設定

- githubのコードを落としておく

```go
go get "github.com/jinzhu/gorm"
go get github.com/go-sql-driver/mysql
```

### ddlを実行

- study/db階層にあるddl.sqlを実行

### ????

goを実行するときにはsudoをつけるとワーニングがでない
面倒くさいので誰か原因究明よろしく

### goからmysqlへの接続
- 今回は[gorm](https://gorm.io/docs/index.html "公式ドキュメント")を用いたsql操作を採用している<br>
mysql-driverのみを用いた接続でもいいが、structの仕様やテストなどの安全面も考慮しこちらを採用<br>
より詳細なsqlの書き方については、上記の公式ドキュメントに乗っています
- セキュアな接続方法や、接続の設定方法などは公式ドキュメントに乗っています<br>
私は心が折れました
```go
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

```

### insertの実行
- [ユーザ登録ページ](http://127.0.0.1:8888/static/user_home.html "ユーザ登録ページ")を用いたsql操作を採用している<br>
- structを用いたinsert処理
- structの定義<br>
jsonで定義しているが、gormを使った定義も可能<br>
用途によって使い分けするべき
```go
type UserData struct {
  Name    string `json:"name,omitempty"`
  Age     int    `json:"age,omitempty"`
  Address string `json:"address,omitempty"`
}

userData := UserData{
  Name:    "名前",
  Age:     21,
  Address: "住所を入れる",
}
```

- 上記で定義したuserDataを使ってinsertを実行
```go
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
```
- コードの説明<br>
connectionについては省略
- db.Tableでinsertを実行するテーブルを設定
- Selectでinsertするカラムを指定
- Createで挿入するデータ(struct)を指定<br>
structには&(ポインタ)指定が必要？
```go
result := db.Table("users").Select("name", "age", "address").Create(&userData)
```
- 結果的に生成され実行されたsql
```sql
INSERT INTO users(name, age, address) VALUES ("名前", 20, "住所");
```

### 注意点
- 知識不足により、DB階層とAP階層に分けることができなかった。(pointerって難しい)<br>
APとDBは同じ階層に設置し、connectだけdb階層に置くとうまくいく

### selectの実行

- Firstの場合はポインタは配列にする必要はない
- Findの場合はポインタの配列にすること
```go
// FindとFirstの違い
db.First: 1件のみの検索 基本的に主キーでorder byがされる
db.Find: 検索条件にマッチするものすべてを持ってくる
```

#### 全件検索
- [全件検索resultページ](http://127.0.0.1:8888/user/AllUser "ユーザ登録ページ")
```go
db, err := getConnect.SqlConnect()
var users []UserData
if err != nil {
  fmt.Println("error")
  fmt.Println(err)
  return nil, false
} else {
  fmt.Println("DBアクセス成功")
  db.Table("users").Find(&users)
  return users, true
}
```
- コードの説明
- 実行結果を取得するストラクトを定義<br>
[]UserDataの部分でカラムを指定可能(サンプルコードでは、name,age,addressのみ)
```go
var users []UserData
```
- 検索結果の格納<br>
上記で定義したusersにFindを使ってポインタに結果を格納
```go
db.Table("users").Find(&users)
```

#### where句・sqlインジェクション
- where句を使用したselect
- bind処理の例<br>
- bindの例
- 主キーが数値の時はみたいな事がドキュメントにあるが、よくわからないのでこれで統一して
- ?にbindされた値が入る。pythonと一緒
```go
db.Table("users").First(&user, "id = ?", id)
```

#### その他のよく使うsqlの書き方

- 時短のため、goファイルへのサンプルコードは書いていないものもある

##### GROUP BY & HAVING
- 集計関数の書き方<br>
Selectの中に書くだけ
```go
db.Table("users").Select("address, count(*) as count").Group("address").Having("count(*) > ?", 1).Find(&result)
```
- 生成されたsql
- 住所ごとの人数を数え、2人以上のデータを調べるsql
```sql
SELECT address, count(*) as count FROM users GROUP BY address HAVING count(*) > 1;
```

#### LIMIT・OFFSET
OFFSETとは<br>
飛ばす行数
- LIMITの書き方
```go
db.Table("users").Limit(10).Find(&result)
```
- 生成されるsql
```sql
SELECT * FROM users limit 10;
```

#### DISTINCT
```go
db.Distinct("name", "age").Order("name, age desc").Find(&results)
```

#### JOIN
```go
db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Find(&results)
db.Table("users").Select("users.name, emails.email").Joins("inner join emails on emails.user_id = users.id").Find(&results)
```


## webアプリの実行

### 仕様ライブラリ
- 今回はginを仕様
- [gin公式ドキュメント](https://pkg.go.dev/github.com/gin-gonic/gin)

### goの設定

```go
git mod init プロジェクト名
を実行しgo.modファイルを生成する(既に実行済みの場合は不要
```

- githubのコードを落としておく

```GO
go get github.com/gin-gonic/gin@v1.7.4
```

### GETデータの取得

- 基本的には文字列型で送られてくるため、用途に合わせたキャストをすること

```go
name := ctx.Query("name") // keyがnameのデータを取得 データが存在していなくてもエラーにならない
name := ctx.DefaultQuery("limit", "10") // 第2引数にデフォルト値を設定できる
```

### POSTデータの取得

- GET同様の仕様

```go
name := ctx.PostForm("name") // keyがnameのデータを取得 データが存在していなくてもエラーにならない
name := ctx.DefaultPostForm("limit", "10") // 第2引数にデフォルト値を設定できる
```

### URLパラメータの値を取得

- URLの値を取得できる
- react routerのparameter的なのを使いたいときに使えるかも？？
- *のところが可変の値になる
- 実行例URL: http://localhost:8888/path/hoge/fuga/piyo/ >>> hoge is /fuga/piyo/
- 実行例URL: http://localhost:8888/path/he/tsushima >>> he is tsushima

```go
r.GET("/path/:name/*action", func (c *gin.Context) {
  name := c.Param("name")
  action := c.Param("action")
  message := name + " is " + action
  c.String(http.StatusOK, message)
})
```

### JSONデータの取得

- 使うかわからんから一旦パス

### jsonデータを返す

- 構造体を定義し,情報を格納する

```go
type response struct {
  Name string `json:"name"`
  Age  int    `json:"page"`
  Food string `json:"food"`
}
```

```go
sampleJson := &response{
  Name: name,
  Age:  age,
  Food: food,
}
```

- jsonとして返却する
```go
ctx.Header("Content-Type", "application/json; charset=utf-8")
ctx.JSON(200, sampleJson)
```

### 注意点
- 構造体の文字はじめは必ず大文字にする
- 以下を実行してもいいが、jsでjsonにエンコードする必要が出るため、要件に合わせてやるように
```go
res, _ := json.Marshal(sampleJson)
ctx.Header("Content-Type", "application/json; charset=utf-8")
ctx.JSON(200, string(res))
```

### mysqlの起動
```bash
sudo /etc/init.d/mysql start
```

## ENVファイルの利用

### 必要なパッケージのインストール
- 以下をインストール
```go
go get github.com/joho/godotenv
```

### envファイルの準備
- プロジェクト直下に.envファイルを生成する
```bash
DB_PASS="自分のpassword"
DB_USER="利用するuser"
DB_PROTOCOL="tcp(localhost:3306)"
DB_NAME="go_example"
```

### goの記述
```go
err = godotenv.Load(".env")

if err == nil {
  DBMS := "mysql"
  USER := os.Getenv("DB_USER")
  PASS := os.Getenv("DB_PASS") //自分の設定したパスワード 
  PROTOCOL := os.Getenv("DB_PROTOCOL")
  DBNAME := os.Getenv("DB_NAME")
}else{
  fmt.Println("envファイルがないです")
}
```

- サーバーの停止はctrl+cでやること
- staticフォルダはgopathで設定したプロジェクト直下に置くこと
