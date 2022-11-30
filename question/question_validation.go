package question

import (
  validation "github.com/go-ozzo/ozzo-validation/v4"
  "github.com/go-ozzo/ozzo-validation/v4/is"
)

// ValidateInsertQuestion 登録する問題のバリデーションチェック
func ValidateInsertQuestion(insertData DataQuestion) error {
  return validation.ValidateStruct(
    &insertData,
    validation.Field(
      &insertData.QuestionUserId,
      validation.Required.Error("セッションエラー"),
      is.Digit.Error("セッションエラー"),
    ),
    validation.Field(
      &insertData.QuestionTitle,
      validation.Required.Error("問題名が入力されていません。"),
      validation.Length(1, 64).Error("問題名は1文字以上64文字以内です"),
    ),
    validation.Field(
      &insertData.QuestionDetail,
      validation.Required.Error("問題詳細が入力されていません。"),
    ),
    validation.Field(
      &insertData.OutputValue,
      validation.Required.Error("出力値が入力されていません。"),
    ),
    validation.Field(
      &insertData.QuestionLang,
      validation.Required.Error("解答言語が選択されていません。"),
      is.Digit.Error("解答言語が選択されていません。"),
    ),
  )
}
