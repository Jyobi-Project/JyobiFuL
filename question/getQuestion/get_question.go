package getQuestion

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetQuestion(c *gin.Context) {
	questionId, err1 := strconv.Atoi(c.DefaultPostForm("questionID", "0"))

	// fmt.Println(questionId)

	var responseData QuestionData
	var selectQuestion DBGetQuestionData

	if err1 != nil {
		responseData = QuestionData{
			ErrorFlag: true,
		}

		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(200, responseData)

		return
	}

	err := true
	selectQuestion, err = SelectDBQuestion(questionId)

	responseData = QuestionData{
		QuestionId:    questionId,
		QuestionTitle: selectQuestion.QuestionTitle,
		QuestionUser: QuestionUser{
			UserId:   selectQuestion.UserId,
			UserName: selectQuestion.UserName,
			UserMail: selectQuestion.UserMail,
			UserIcon: selectQuestion.UserIcon,
		},
		QuestionCounter: QuestionCounter{
			View:   selectQuestion.QuestionView,
			Answer: selectQuestion.AnswerCount,
		},
		QuestionLang: selectQuestion.LanguageName,
		QuestionDetail: QuestionDetail{
			QuestionText: selectQuestion.QuestionDetail,
		},
		QuestionInputValue: QuestionInputValue{
			Set:   false,
			Value: selectQuestion.InputValue,
		},
		QuestionOutputValue: QuestionOutputValue{
			Set:   false,
			Value: selectQuestion.OutputValue,
		},
		CreateAt:  selectQuestion.CreateAt,
		UpdateAt:  selectQuestion.UpdateAt,
		ErrorFlag: err,
	}

	// fmt.Printf("(%%#v) %#v\n", selectQuestion)
	// fmt.Printf("(%%#v) %#v\n", responseData)

	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(200, responseData)
}
