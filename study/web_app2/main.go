package main

import (
	"Jyobi-Project/study/web_app2/router"
	"log"
)

func main() {
	r := router.GetRouter()
	// サーバーを起動
	err := r.Run("127.0.0.1:8888")
	if err != nil {
		log.Fatal("サーバー起動に失敗", err)
	}
}
