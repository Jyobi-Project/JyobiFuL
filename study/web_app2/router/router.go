package router

import (
	"Jyobi-Project/study/web_app2/home"
	"Jyobi-Project/study/web_app2/product"
	"Jyobi-Project/study/web_app2/user"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", home.IndexPage)
	r.GET("/user", user.Home)
	r.GET("/product", product.Home)
	r.GET("/user/update", user.Update)
	r.GET("/user/delete", user.UserDelete)
	return r
}
