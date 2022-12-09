package users

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// 空白を削除する関数
func ReplacementSpace(str string) string {
	return strings.TrimSpace(str)
}

// PWをハッシュ化する関数
func ConvertPwToHash(pw string) string {
	hashPw, _ := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hashPw)
}

// PWとハッシュ値を比較する関数
func EqualHashToPw(hash string, pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}
