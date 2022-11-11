package router

import (
	"Jyobi-Project/study/web-app3/home"
	"Jyobi-Project/study/web-app3/users"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/static", http.Dir("static"))

	r.GET("/", home.Home)
	//r.POST("/create_user", users.CreateUser)
	r.POST("/create_user", users.CreateUser)
	r.POST("/select_mail", users.SelectWhereUser)
	r.GET("/SelectAllUser", users.SelectAllUser)

	return r
}
