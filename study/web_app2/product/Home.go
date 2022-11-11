package product

import "github.com/gin-gonic/gin"

func Home(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=UTF-8")
	ctx.String(200, "<h1>Product Homeの画面です</h1>")
}
