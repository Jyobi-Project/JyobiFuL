package router

import (
  "Jyobi-Project/question"
  "github.com/gin-gonic/gin"
  "net/http"
)

func GetRouter() *gin.Engine {
  r := gin.Default()

  r.StaticFS("/static", http.Dir("static"))

  // 問題登録ページ（サンプル）
  q := r.Group("/question")
  {
    q.GET("/create", question.Question)
    // 問題を登録する
    q.POST("/create", question.InsertQuestion)
  }

  return r
}
