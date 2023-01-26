package createQuestion

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

// InsertQuestion 問題を登録する
func InsertQuestion(context *gin.Context) {
  // requestデータを取り出す
  var request QuestionData
  if err := context.BindJSON(&request); err != nil {
    context.JSON(http.StatusOK, gin.H{"error": err.Error()})
    return
  }

  // レスポンスのstruct
  type response struct {
    Result bool  `json:"result"`
    Error  error `json:"error,omitempty"`
  }

  // バリデーションチェック
  if invalid := ValidateInsertQuestion(request); invalid != nil {
    // response 結果を入れる
    resultJson := &response{
      Result: false,
      Error:  invalid,
    }

    context.Header("Content-Type", "application/json; charset=utf-8")
    context.JSON(http.StatusOK, resultJson)
    return
  }

  // 問題を登録する
  result := InsertDBQuestion(request)
  context.Header("Content-Type", "application/json; charset=utf-8")
  context.JSON(http.StatusOK, result)

  // マイページにリダイレクト（今はページがない）
  // context.Redirect(http.StatusMovedPermanently, "http://localhost:3000/question/create")
}
