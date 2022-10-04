package user

import (
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=UTF-8")
	ctx.String(200, "<h1>UserHomeの画面です</h1>")
}

func Update(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=UTF-8")
	ctx.String(200, "<h1>UserUpdateの画面です</h1>")
}

func UserDelete(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=UTF-8")
	ctx.String(200, "<h1>UserDeleteの画面です</h1>")
}
