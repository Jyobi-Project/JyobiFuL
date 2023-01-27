package questionList

type SearchData struct {
  SearchValue string `json:"searchValue"`
  Tag         `json:"tag"`
  Sort        bool `json:"sort"`
}

type Tag struct {
  daytime  int
  bookmark int
  view     int
  answer   int
}

type DBQuestionList struct {
  QuestionId    int    `json:"questionId"`
  QuestionTitle string `json:"title"`
  UserName      string `json:"contributor"`
  LanguageName  string `json:"language"`
  QuestionView  int    `json:"viewsCount"`
  AnswerCount   int    `json:"responseCount"`
  IsAnswered    bool   `json:"isAnswered"`
  IsBookmark    bool   `json:"isBookmark"`
}
