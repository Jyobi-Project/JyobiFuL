package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.StaticFS("/static", http.Dir("static"))

	return r
}
