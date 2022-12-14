package home

import "github.com/gin-gonic/gin"

func IndexPage(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=UTF-8")
	ctx.String(200, "<h1>HOMEの画面です</h1>")
}
