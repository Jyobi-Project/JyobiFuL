package createQuestion

import (
  "Jyobi-Project/db/connect"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

// InsertDBQuestion 問題をDBに登録する
func InsertDBQuestion(questionData QuestionData) bool {
  db, err := getConnect.SqlConnect()
  if err != nil {
    fmt.Println("error")
    fmt.Println(err)
    return false
  } else {
    result := db.Exec(
      "INSERT INTO questions(question_user_id, question_title, question_detail, input_value, output_value, question_lang, example_answer) VALUES (?,?,?,?,?,(SELECT language_id FROM languages WHERE language_name = ?),?)",
      questionData.UserId,
      questionData.QuestionTitle,
      questionData.QuestionDetail,
      questionData.InputValue,
      questionData.OutputValue,
      questionData.QuestionLang,
      questionData.ExampleAnswer,
    )

    if result.Error != nil {
      return false
    } else {
      return true
    }
  }
}
