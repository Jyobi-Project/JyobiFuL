package home

import "github.com/gin-gonic/gin"

// 登録ページを表示
func Home(ctx *gin.Context) {
	ctx.Redirect(302, "/static/web-app3.html")
}
