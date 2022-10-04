# Jyobi-Project

## db接続手順書

### 事前準備

- [mysqlをubuntuにインストール](https://learn.microsoft.com/ja-jp/windows/wsl/tutorials/wsl-databasem, "公式ドキュメント")
- [パスワード設定でエラー起きたときの処理](https://exerror.com/failed-error-set-password-has-no-significance-for-user-rootlocalhost-as-the-authentication-method-used-doesnt-store-authentication-data-in-the-mysql-server/, "エラー時の参考")

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

## webアプリの実行

### goの設定

```go
git mod init フォルダ名
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
- 実行するフォルダ直下に.envファイルを生成する
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
