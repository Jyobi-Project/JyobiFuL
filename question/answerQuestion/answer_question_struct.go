package answerquestion

type ResponseData struct {
	Index  int    `json:"index,omitempty"`
	Title  string `json:"title,omitempty"`
	Result bool   `json:"result,omitempty"`
}
