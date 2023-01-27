package questionList

import (
  "Jyobi-Project/db/connect"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

// SelectAllDBQuestion 問題をDBに登録する
func SelectAllDBQuestion() []DBQuestionList {
  db, err := getConnect.SqlConnect()
  if err != nil {
    fmt.Println("error")
    fmt.Println(err)
    return nil
  } else {
    var questions []DBQuestionList
    result := db.Raw("SELECT q.question_id, " +
      "question_title, " +
      "user_name, " +
      "l.language_name, " +
      "question_view, " +
      "answer_count, " +
      "qa.user_id IS NOT NULL as is_answered, " +
      "qb.user_id IS NOT NULL as is_bookmark " +
      "FROM questions as q " +
      "INNER JOIN languages as l ON l.language_id = q.question_lang " +
      "INNER JOIN users as u ON u.user_id = q.question_user_id " +
      "LEFT OUTER JOIN questions_bookmarks as qb ON qb.question_id = q.question_id AND qb.user_id = q.question_user_id " +
      "LEFT OUTER JOIN question_answers as qa ON qa.question_id = q.question_id AND qa.user_id = q.question_user_id " +
      "ORDER BY q.create_at DESC",
    ).Scan(&questions)

    if result.Error != nil {
      return nil
    } else {
      return questions
    }
  }
}
