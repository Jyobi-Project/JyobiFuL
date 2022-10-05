package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserData struct {
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
}

// 登録ページを表示
func Home(ctx *gin.Context) {
	ctx.Redirect(302, "/static/user_home.html")
}

func Update(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=UTF-8")
	ctx.String(200, "<h1>UserUpdateの画面です</h1>")
}

func UserDelete(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html; charset=UTF-8")
	ctx.String(200, "<h1>UserDeleteの画面です</h1>")
}

func Insert(ctx *gin.Context) {

	name := ctx.DefaultPostForm("name", "何もない")
	age, _ := strconv.Atoi(ctx.DefaultPostForm("age", "0"))
	address := ctx.DefaultPostForm("address", "何もない")
	userData := UserData{
		Name:    name,
		Age:     age,
		Address: address,
	}

	if SqlInsert(userData) {
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>登録成功</h1>")
	} else {
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>登録失敗</h1>")
	}
}
