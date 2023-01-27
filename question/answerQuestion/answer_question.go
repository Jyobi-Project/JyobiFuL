package answerquestion

import (
	"Jyobi-Project/question/getQuestion"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
)

func AnswerQuestion(c *gin.Context) {
	// ---------------------------postを受け取る------------------------
	questionId, err1 := strconv.Atoi(c.DefaultPostForm("questionID", "0")) // 問題ID
	userId, err2 := strconv.Atoi(c.DefaultPostForm("userID", "0"))         // 回答者ID
	token := c.DefaultPostForm("token", "")                                // 回答者token
	code := c.DefaultPostForm("code", "")                                  // 回答文

	// ------------------------------------------------------------------

	// ---------------------変数の初期化---------------------------------

	var selectQuestion getQuestion.DBGetQuestionData
	var sqlErr bool
	responseData := Result{
		Result: []ResponseData{{
			Index:  0,
			Title:  "",
			Result: false,
		}},
	}

	// -----------------------------------------------------------------

	if err1 != nil || err2 != nil {

		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, responseData)

		return
	}

	fmt.Println(userId)
	fmt.Println(token)

	// 問題情報の取得
	selectQuestion, sqlErr = getQuestion.SelectDBQuestion(questionId)

	// エラーならその時点でreturn
	if sqlErr {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, responseData)

		return
	}

	// --------------------問題情報から欲しい情報を抜き取る------------------
	language := selectQuestion.LanguageName
	languageExtension := GetMine(language)
	inputValue := selectQuestion.InputValue
	outputValue := selectQuestion.OutputValue

	// ------------------------------------------------------------------

	// --------------------ファイルデータの初期化 -------------------------
	fileWriteData := []byte(code)            //回答文をファイルに書き込むための変数
	fileWriteInputData := []byte(inputValue) // inputValueをファイルに書き込むための変数

	dt := strconv.FormatInt(time.Now().Unix(), 10) // ファイル名に使うunix time
	fileName := dt + languageExtension             // 回答ファイルの拡張子を取得
	inputFileName := dt + ".txt"                   // inputvalueのファイルの名前

	// ------------------------------------------------------------------

	// ------------------------ファイル作成-------------------------------
	fp, err := os.Create(fileName)        // 回答文が書き込まれるファイル
	fp2, err2 := os.Create(inputFileName) // inputValueが書き込まれるファイル

	defer fp.Close()  // 何があってもファイルポインタを閉じる
	defer fp2.Close() // 何があってもファイルポインタを閉じる

	// -------------------------------------------------------------------

	// -----------------------ファイル作成がエラーの時----------------------
	if err != nil && err2 != nil {

		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, responseData)

		return
	}

	// --------------------------------------------------------------------

	// -----------------------------ファイル書込み--------------------------
	// 戻り値の一つ目は書き込んだ文字数なので使わない
	_, err = fp.Write(fileWriteData)
	_, err2 = fp2.Write(fileWriteInputData)

	if err != nil && err2 != nil {

		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, responseData)

		return
	}

	// -------------------------------------------------------------------

	//  -------------------------dockerでプログラムを実行-------------------
	command := exec.Command("sh", "docker.sh", fileName, language, dt, inputFileName)

	commandErr := command.Run()

	if commandErr != nil {
		fmt.Println(commandErr)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, responseData)
		return
	}

	// --------------------------------------------------------------------

	// -----------------------プログラムの実行結果を取得----------------------
	resultFileName := dt + "_result.txt"

	resultFile, err4 := os.Open(resultFileName)
	if err4 != nil {
		// fmt.Println(err4)
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, responseData)
	}
	defer resultFile.Close()

	buf := make([]byte, 64)
	for {
		n, err5 := resultFile.Read(buf)
		if n == 0 {
			break
		}
		if err5 != nil {
			fmt.Println(err5)
			c.Header("Content-Type", "application/json; charset=utf-8")
			c.JSON(200, responseData)
			return
		}
	}

	answerResult := string(buf)

	// fmt.Println(answerResult)
	// fmt.Println(outputValue)

	result := chop(answerResult)

	fmt.Println(result)

	// ---------------------------------------------------------------------

	// outputValueのなかにanswerResultが含まれているか
	if strings.Contains(answerResult, outputValue) {
		responseData.Result[0].Result = true
	}

	// 作成したファイルを削除
	os.Remove(fileName)
	os.Remove(inputFileName)
	os.Remove(resultFileName)

	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(200, responseData)

}

func chop(s string) string {
	if strings.HasSuffix(s, "\r\n") {
		fmt.Println("rあるよ")
	}
	s = strings.TrimRight(s, "\n")
	if strings.HasSuffix(s, "\r") {
		s = strings.TrimRight(s, "\r")
	}

	return s
}
