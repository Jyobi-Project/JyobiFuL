package createQuestion

import (
  "github.com/gin-gonic/gin"
)

// Question 問題を登録するページ（サンプル）
func Question(context *gin.Context) {
  context.Redirect(302, "/static/create_question.html")
}

// InsertQuestion 問題を登録する
func InsertQuestion(context *gin.Context) {
  // structにデータを入れる
  insertData := DataQuestion{
    QuestionUserId: context.DefaultPostForm("userId", ""),   // ユーザID
    QuestionDetail: context.DefaultPostForm("title", ""),    // 問題名
    QuestionTitle:  context.DefaultPostForm("detail", ""),   // 問題詳細
    InputValue:     context.DefaultPostForm("input", ""),    // 入力値
    OutputValue:    context.DefaultPostForm("output", ""),   // 出力値
    QuestionLang:   context.DefaultPostForm("language", ""), // 解答言語
    ExampleAnswer:  context.DefaultPostForm("answer", ""),   // 模範解答
  }

  // レスポンスのstruct
  type response struct {
    Result bool  `json:"result"`
    Error  error `json:"error,omitempty"`
  }

  // バリデーションチェック
  if invalid := ValidateInsertQuestion(insertData); invalid != nil {
    // response structに結果を入れる
    resultJson := &response{
      Result: false,
      Error:  invalid,
    }

    context.Header("Content-Type", "application/json; charset=utf-8")
    context.JSON(200, resultJson)
    return
  }

  // 問題を登録する
  result := InsertDBQuestion(insertData)

  // response structに結果を入れる
  resultJson := &response{
    Result: result,
  }

  // レスポンスを返す
  context.Header("Content-Type", "application/json; charset=utf-8")
  context.JSON(200, resultJson)
}
