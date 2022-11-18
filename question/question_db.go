package question

import (
  "Jyobi-Project/db/connect"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

// DataQuestion 問題のstruct
type DataQuestion struct {
  QuestionId     int    `json:"questionId,omitempty"`   // 問題ID
  QuestionUserId int    `json:"userId,omitempty"`       // 作成者ID
  QuestionTitle  string `json:"title,omitempty"`        // 問題名
  QuestionDetail string `json:"detail,omitempty"`       // 問題詳細
  InputValue     string `json:"input,omitempty"`        // 入力値
  OutputValue    string `json:"output,omitempty"`       // 出力値
  QuestionLang   int    `json:"language,omitempty"`     // 解答言語
  ExampleAnswer  string `json:"answer,omitempty"`       // 模範解答
  QuestionView   int    `json:"questionView,omitempty"` // 閲覧数
  CreateAt       string `json:"createAt,omitempty"`     // 作成日
  UpdateAt       string `json:"updateAt,omitempty"`     // 更新日
}

// InsertDBQuestion 問題をDBに登録する
func InsertDBQuestion(questionData DataQuestion) bool {
  db, err := getConnect.SqlConnect()
  if err != nil {
    fmt.Println("error")
    fmt.Println(err)
    return false
  } else {
    result := db.Table("questions").Select(
      "question_user_id",
      "question_title",
      "question_detail",
      "input_value",
      "output_value",
      "question_lang",
      "example_answer",
    ).Create(&questionData)

    if result.Error != nil {
      return false
    } else {
      return true
    }
  }
}
