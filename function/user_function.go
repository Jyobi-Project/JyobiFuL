package function

// Saltを作成する関数
func CreateSalt() string {
	return salt
}

// PWをハッシュ化する関数
func HashPw(salt string, pw string) string {
	return hashPw
}

// 空白置換をする関数
func SpaceReplacement(str string) string {
	return str
}

// 文字数チェックする関数
func LengthValidation(str string, max int, min int) bool {
	// TODO: 文字列の長さを取得して変数strにぶち込んでmaxとminと比較しろ
	return true
}

// メールのバリデーションチェックをする関数
func MailValidation(mail string) bool {
	return true
}

// パスワードのバリデーションチェックをする関数
func PwValidation(pw string) bool {
	return true
}
