package users

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"regexp"
)

// 名前のバリデーションチェックをする関数
func ValidationName(name string) error {
	return validation.Validate(
		// 空白を削除
		ReplacementSpace(name),
		// 空きを許容しない
		validation.Required.Error("名前が入力されていません"),
		// 1文字以上64文字以下である
		validation.Length(1, 64).Error("名前は1~64文字で入力してください"),
	)
}

// メールのバリデーションチェックをする関数
func ValidationMail(mail string) error {
	return validation.Validate(
		// 空白を削除
		ReplacementSpace(mail),
		// 空きを許容しない
		validation.Required.Error("メールが入力されていません"),
		// 5文字以上64文字以下である
		validation.Length(5, 64).Error("メールは5~64文字で入力してください"),
		// Email形式のみ
		is.Email.Error("メールの形式が違います"),
	)
}

// パスワードのバリデーションチェックをする関数
func ValidationPW(pw string) error {
	return validation.Validate(
		// 空白を削除
		ReplacementSpace(pw),
		// 空きを許容しない
		validation.Required.Error("パスワードが入力されていません"),
		// 5文字以上64文字以下である
		validation.Length(8, 64).Error("パスワードは8~64文字で入力してください"),
		// パスワードの形式のみ
		validation.Match(regexp.MustCompile("^[a-zA-Z\\d.?/-@]{8,64}$")).Error("パスワードの形式が違います"),
	)
}
