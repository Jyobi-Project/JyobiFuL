# Jyobi-Project

## バリデーションチェック
バリデーションの種類はここに載っている！！！

👇👇👇👇👇👇👇👇👇👇👇👇👇👇👇👇

参考サイト：[ozzo-validation](https://github.com/go-ozzo/ozzo-validation)

👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆👆

### インストール
```go
  go get github.com/go-ozzo/ozzo-validation
```

### エラー
インストールしたとき`missing go.sum entry for module providing package`エラーでた
```go
  go mod tidy
```
で直った。
### コマンド詳細
go mod tidyコマンドでは、モジュール内の全てのパッケージのビルドタグの組み合わせを確認するので、安全に不要なモジュールを削除することができます。
一方で、gobuildやgotestでは、不要なパッケージの削除はされません。ビルド時に単一のパッケージのロードしか行われず、使用されていないパッケージを知ることができないためです。

らしい。

## セッション

### インストール
```go
go get github.com/gin-contrib/sessions
```

### インポート
```go
import "github.com/gin-contrib/sessions"
```

### 使い方
#### セッションの宣言
```go
store := cookie.NewStore([]byte("secret"))
r.Use(sessions.Sessions("jyobifulSession", store))
```

#### set
```go
session := sessions.Default(c)
session.Set("UserId")
session.Save()
```

#### get
```go
session := sessions.Default(c)
userId := session.Get("UserId")
```

#### clear
```go
func Logout(c *gin.Context) {
    session := sessions.Default(c)
    session.Clear()
    session.Save()
}
```

### 使うぜ
これがセッションチェックの関数だぜ
```go
func checkSession() gin.HandlerFunc {
  return func(c *gin.Context) {
    session := sessions.Default(c)
    userId := session.Get("UserId")

    // セッションがない場合、ログインフォームをだす
    if userId == nil {
      c.Redirect(http.StatusMovedPermanently, "/login")
      c.Abort()
    } else {
      c.Next()
    }
  }
}
```
セッションチェックを行いたいページをグループ化だぜ！  
それに`Use`を使ってセッションチェック関数を呼び出すぜ！
```go
q := r.Group("/question")
  q.Use(checkSession())
  {
    q.GET("/create-form", question.Question)
    // 問題を登録する
    q.POST("/create", question.InsertQuestion)
  }
```
以上だぜ！