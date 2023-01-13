package question

import (
  "Jyobi-Project/db/connect"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

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
