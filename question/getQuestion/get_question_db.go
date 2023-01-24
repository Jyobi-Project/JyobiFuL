package getQuestion

import (
	getConnect "Jyobi-Project/db/connect"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// InsertDBQuestion 問題をDBに登録する
func SelectDBQuestion(questionId int) (DBGetQuestionData, bool) {
	// 変数の初期化
	var selectQuestion DBGetQuestionData
	// コネクションの取得
	db, err := getConnect.SqlConnect()
	// 取得に失敗した場合はエラーフラグをturuにし、return
	if err != nil {
		fmt.Println(err)
		return selectQuestion, true
	}

	// dbから問題情報を取得
	result := db.Table("questions").Select(
		"question_id, question_title, u.user_id, user_name, user_mail, user_icon, language_name, question_detail, input_value, output_value, example_answer, question_view, answer_count, create_at, update_at").Joins(
		"INNER JOIN users as u on u.user_id = questions.question_user_id").Joins(
		"INNER JOIN languages as l on l.language_id = questions.question_lang").Where(
		"question_id = ?", questionId).First(&selectQuestion)

	// 実行に失敗した場合は、errorにする
	if result.Error != nil {
		return selectQuestion, true
	} else {
		return selectQuestion, false
	}

}
