package question

import (
  "github.com/gin-gonic/gin"
  "strconv"
)

// Question 問題を登録するページ（サンプル）
func Question(ctx *gin.Context) {
  ctx.Redirect(302, "/static/create_question.html")
}

// InsertQuestion 問題を登録する
func InsertQuestion(context *gin.Context) {
  userId, _ := strconv.Atoi(context.DefaultPostForm("userId", "0"))     // ユーザID
  title := context.DefaultPostForm("title", "")                         // 問題名
  detail := context.DefaultPostForm("detail", "")                       // 問題詳細
  input := context.DefaultPostForm("input", "")                         // 入力値
  output := context.DefaultPostForm("output", "")                       // 出力値
  language, _ := strconv.Atoi(context.DefaultPostForm("language", "0")) // 解答言語
  answer := context.DefaultPostForm("answer", "")                       // 模範解答

  // question data structにデータを入れる
  insertData := DataQuestion{
    QuestionUserId: userId,
    QuestionTitle:  title,
    QuestionDetail: detail,
    InputValue:     input,
    OutputValue:    output,
    QuestionLang:   language,
    ExampleAnswer:  answer,
  }

  // 問題を登録する
  result := InsertDBQuestion(insertData)

  // レスポンスのstruct
  type response struct {
    Result bool `json:"result"`
  }

  // response structに結果を入れる
  resultJson := &response{
    Result: result,
  }

  // レスポンスを返す
  context.Header("Content-Type", "application/json; charset=utf-8")
  context.JSON(200, resultJson)
}
