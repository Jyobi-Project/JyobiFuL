package users

import (
	"github.com/gin-gonic/gin"
)

// ログイン処理

func Login(ctx *gin.Context) {
	mail := ctx.PostForm("mail")
	pw := ctx.PostForm("pw")

	// mailを元にhash化されたpwをもってくる
	flag, hashPw := SqlSelectUser(mail)

	if flag {
		if EqualHashToPw(hashPw, pw) {
			ctx.Header("Content-Type", "text/html; charset=UTF-8")
			ctx.String(200, "<h1>ログイン成功だよー</h1>")
		} else {
			ctx.Header("Content-Type", "text/html; charset=UTF-8")
			ctx.String(200, "<h1>パスワードが違うよー</h1>")
		}
	} else {
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>SQL ERROR</h1>")
	}
}
