package router

import (
	"Jyobi-Project/question"
	answerquestion "Jyobi-Project/question/answerQuestion"
	"Jyobi-Project/question/getQuestion"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	// ここからCorsの設定
	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	r.StaticFS("/static", http.Dir("static"))

	// 問題登録ページ（サンプル）
	q := r.Group("/question")
	{
		q.GET("/create", question.Question)
		// 問題を登録する
		q.POST("/create", question.InsertQuestion)
		// 問題を取得する
		q.POST("/get", getQuestion.GetQuestion)
		q.POST("/answer_question", answerquestion.AnswerQuestion)
	}

	return r
}
