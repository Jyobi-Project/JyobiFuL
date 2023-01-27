package questionList

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func SelectQuestion(context *gin.Context) {
  var questions []DBQuestionList
  questions = SelectAllDBQuestion()
  if questions == nil {
    context.JSON(http.StatusOK, gin.H{"error": "selectエラー"})
    return
  }

  context.Header("Content-Type", "application/json; charset=utf-8")
  context.JSON(http.StatusOK, questions)
}
