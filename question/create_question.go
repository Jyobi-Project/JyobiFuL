package question

import (
  "github.com/gin-contrib/sessions"
  "github.com/gin-gonic/gin"
)

// DataQuestion 問題のstruct
type DataQuestion struct {
  QuestionId     string `json:"questionId,omitempty"`   // 問題ID
  QuestionUserId string `json:"userId,omitempty"`       // 作成者ID
  QuestionTitle  string `json:"title,omitempty"`        // 問題名
  QuestionDetail string `json:"detail,omitempty"`       // 問題詳細
  InputValue     string `json:"input,omitempty"`        // 入力値
  OutputValue    string `json:"output,omitempty"`       // 出力値
  QuestionLang   string `json:"language,omitempty"`     // 解答言語
  ExampleAnswer  string `json:"answer,omitempty"`       // 模範解答
  QuestionView   string `json:"questionView,omitempty"` // 閲覧数
  CreateAt       string `json:"createAt,omitempty"`     // 作成日
  UpdateAt       string `json:"updateAt,omitempty"`     // 更新日
}

// Question 問題を登録するページ（サンプル）
func Question(context *gin.Context) {
  context.Redirect(302, "/static/create_question.html")
}

// InsertQuestion 問題を登録する
func InsertQuestion(context *gin.Context) {
  session := sessions.Default(context)
  userId := session.Get("UserId")
  // structにデータを入れる
  insertData := DataQuestion{
    QuestionUserId: userId.(string),                         // ユーザID
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
