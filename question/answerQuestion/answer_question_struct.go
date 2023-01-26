package answerquestion

type ResponseData struct {
	Index  int    `json:"index"`
	Title  string `json:"title"`
	Result bool   `json:"result"`
}

type Result struct {
	ResponseData `json:"result"`
}
