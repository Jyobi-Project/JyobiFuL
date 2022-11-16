package router

import (
	"Jyobi-Project/create-account/home"
	"Jyobi-Project/create-account/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/static", http.Dir("static"))

	r.GET("/", home.Home)
	r.POST("/create_account", users.CreateAccount)

	return r
}
