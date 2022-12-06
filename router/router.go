package router

import (
  "Jyobi-Project/question"
  "github.com/gin-contrib/sessions"
  "github.com/gin-contrib/sessions/cookie"
  "github.com/gin-gonic/gin"
  "net/http"
)

func GetRouter() *gin.Engine {
  r := gin.Default()

  // セッションの宣言
  store := cookie.NewStore([]byte("secret"))
  r.Use(sessions.Sessions("jyobifulSession", store))

  r.GET("/logout", func(context *gin.Context) {
    // セッションからデータを破棄する
    session := sessions.Default(context)
    session.Clear()
    session.Save()

    context.Redirect(302, "/static/index.html")
  })

  r.GET("/login", func(context *gin.Context) {
    context.Redirect(302, "/static/index.html")
  })

  r.StaticFS("/static", http.Dir("static"))
  // 問題登録ページ（サンプル）

  q := r.Group("/question")
  q.Use(checkSession())
  {
    q.GET("/create-form", question.Question)
    // 問題を登録する
    q.POST("/create", question.InsertQuestion)
  }

  return r
}

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
