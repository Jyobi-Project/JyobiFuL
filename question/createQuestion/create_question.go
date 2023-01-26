package createQuestion

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

// Question 問題を登録するページ（サンプル）
func Question(context *gin.Context) {
  context.Redirect(302, "/static/create_question.html")
}

// InsertQuestion 問題を登録する
func InsertQuestion(context *gin.Context) {
  // requestデータを取り出す
  var request QuestionData
  if err := context.BindJSON(&request); err != nil {
    context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
    context.JSON(http.StatusBadRequest, resultJson)
    return
  }

  // 問題を登録する
  if result := InsertDBQuestion(request); result != true {
    context.Header("Content-Type", "application/json; charset=utf-8")
    context.JSON(http.StatusBadRequest, result)
  }

  // マイページにリダイレクト（今はページがない）
  context.Redirect(http.StatusOK, "http://localhost:3000/question/create")
}
