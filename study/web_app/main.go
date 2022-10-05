package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type response struct {
	Name string `json:"name"`
	Age  int    `json:"page"`
	Food string `json:"food"`
}

func main() {
	router := gin.Default()

	// 自動的にファイルを返すよう設定 --- (*1)
	router.StaticFS("/static", http.Dir("static"))

	// ルートなら /static/index.html にリダイレクト --- (*2)
	router.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(302, "/static/index.html")
	})

	// フォームの内容を受け取って挨拶する --- (*3)
	router.GET("/get", func(ctx *gin.Context) {
		name := ctx.Query("name")
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>Hello, "+name+"</h1>")
	})

	router.POST("/post", func(ctx *gin.Context) {
		name := ctx.DefaultPostForm("name", "何もない")
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>dbに以下を登録するよ: "+name+"</h1>")
	})

	router.GET("/path/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		message := name + " is " + action
		ctx.Header("Content-Type", "text/html; charset=UTF-8")
		ctx.String(200, "<h1>"+message+"</h1>")
	})

	router.POST("/get_json", func(ctx *gin.Context) {
		name := ctx.DefaultPostForm("name", "何もない")
		age, _ := strconv.Atoi(ctx.DefaultPostForm("age", "0"))
		food := ctx.DefaultPostForm("food", "何もない")
		sampleJson := response{
			Name: name,
			Age:  age,
			Food: food,
		}
		//res, _ := json.Marshal(sampleJson)
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.JSON(200, sampleJson)
	})

	router.GET("/get_json_test", func(ctx *gin.Context) {
		sampleJson := response{
			Name: "name",
			Age:  21}
		fmt.Println(sampleJson)
		res3B, _ := json.Marshal(sampleJson)
		fmt.Println(res3B)
		fmt.Println(string(res3B))
		ctx.Header("Content-Type", "application/json; charset=utf-8")
		ctx.JSON(200, string(res3B))
	})

	// サーバーを起動
	err := router.Run("127.0.0.1:8888")
	if err != nil {
		log.Fatal("サーバー起動に失敗", err)
	}
}
