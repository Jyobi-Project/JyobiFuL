package users

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserData struct {
	Name     string `json:"name,omitempty"`
	Mail     string `json:"mail,omitempty"`
	Age      int    `json:"age,omitempty"`
	Password string `json:"password,omitempty"`
}

func CreateUser(ctx *gin.Context) {
	name := ctx.DefaultPostForm("name", "名無し")
	mail := ctx.DefaultPostForm("mail", "sample@gmail.com")
	age, _ := strconv.Atoi(ctx.DefaultPostForm("age", "0"))
	pw := ctx.DefaultPostForm("pw", "pw")

	userData := UserData{
		Name:     name,
		Mail:     mail,
		Age:      age,
		Password: pw,
	}

	ctx.Header("Content-Type", "application/json; charset=utf-8")
	ctx.JSON(200, userData)

	if SqlInsert(userData) {
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>登録成功</h1>")
	} else {
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>登録失敗</h1>")
	}
}

func SelectAllUser(ctx *gin.Context) {
	result, flag := SqlAllSelect()
	if flag {
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.JSON(200, result)
	} else {
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>DB Errorです</h1>")
	}
}

func SelectWhereUser(ctx *gin.Context) {
	mail := ctx.DefaultPostForm("mail", "sample@gmail.com")

	result, flag := SqlWhereSelect(mail)

	if flag {
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.JSON(200, result)
	} else {
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>DB Errorです</h1>")
	}
}

func UpdateUser(ctx *gin.Context) {
	mail := ctx.DefaultPostForm("mail", "sample@gmail.com")

	result, flag := SqlWhereSelect(mail)

	if flag {
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.JSON(200, result)
	} else {
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>DB Errorです</h1>")
	}
}
