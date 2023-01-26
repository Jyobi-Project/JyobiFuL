package createQuestion

import (
  validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ValidateInsertQuestion 登録する問題のバリデーションチェック
func ValidateInsertQuestion(insertData QuestionData) error {
  return validation.ValidateStruct(
    &insertData,
    validation.Field(
      &insertData.UserId,
      validation.Required.Error("ユーザIDがありません"),
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
    ),
  )
}
