package users

import (
	"github.com/gin-gonic/gin"
)

// アカウント作成のストラクト
type UserData struct {
	User_name     string `json:"user_name,omitempty"`
	User_mail     string `json:"user_mail,omitempty"`
	User_password string `json:"user_password,omitempty"`
	Salt          string `json:"salt,omitempty"`
	User_icon     string `json:"user_icon,omitempty"`
}

// アカウント作成関数
func CreateAccount(ctx *gin.Context) {
	name := ctx.DefaultPostForm("name", "名無しさん")
	mail := ctx.DefaultPostForm("mail", "")
	pw := ctx.DefaultPostForm("pw", "")
	//rePw := ctx.DefaultPostForm("rePw", "")

	// TODO: 後でソルトを作成し記述する -> テストの為決め打ちなう
	//salt := function.CreateSalt()
	salt := "afoajofijsioppfj"
	// TODO: デフォルト決め打ち -> AWS S3のパスをぶちこむ
	Icon := "Default"

	userData := UserData{
		User_name:     name,
		User_mail:     mail,
		User_password: pw,
		Salt:          salt,
		User_icon:     Icon,
	}

	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(200, userData)

	if SqlCreateAccount(userData) {
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>登録成功</h1>")
	} else {
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>登録失敗</h1>")
	}
}
