package router

import (
	"Jyobi-Project/study/web_app2/home"
	"Jyobi-Project/study/web_app2/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/static", http.Dir("static"))

	r.GET("/", home.IndexPage)
	r.GET("/user", user.Home)

	return r
}
