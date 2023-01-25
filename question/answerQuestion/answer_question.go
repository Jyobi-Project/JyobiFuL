package answerquestion

import (
	"fmt"
	"os"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
)

func AnswerQuestion(c *gin.Context) {
	// questionIDの取得
	questionId, err1 := strconv.Atoi(c.DefaultPostForm("questionID", "0"))
	userId, err2 := strconv.Atoi(c.DefaultPostForm("userID", "0"))
	token := c.DefaultPostForm("token", "")
	code := c.DefaultPostForm("code", "")
	fmt.Println(questionId)
	fmt.Println(userId)
	fmt.Println(token)
	fmt.Println(code)

	responseData := ResponseData{
		Index:  0,
		Title:  "サンプルケース1",
		Result: false,
	}

	if err1 != nil || err2 != nil {

		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, responseData)

		return
	}

	fileWriteData := []byte(code)

	dt := strconv.FormatInt(time.Now().Unix(), 10)

	fp, err := os.Create(dt + ".txt")

	defer fp.Close()

	if err != nil {
		fmt.Println(err)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, responseData)

		return
	}

	_, err = fp.Write(fileWriteData)

	if err != nil {
		fmt.Println(err)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, responseData)

		return
	}

	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(200, responseData)

}
