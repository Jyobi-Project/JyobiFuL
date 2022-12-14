package router

import (
	"Jyobi-Project/account/home"
	"Jyobi-Project/account/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/static", http.Dir("static"))

	r.GET("/", home.Home)
	r.POST("/CreateAccount", users.CreateAccount)
	r.POST("/login", users.Login)
	return r
}
