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
  r.GET("/question", question.Question)
  // 問題を登録する
  r.POST("/question/create", question.InsertQuestion)

  return r
}
