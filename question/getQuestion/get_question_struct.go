package question

type QuestionCounter struct {
	View   int `json:"view"`
	Answer int `json:"answer"`
}

type QuestionInputValue struct {
	Set   bool   `json:"set,omitempty"`
	Value string `json:"value,omitempty"`
}

type QuestionOutputValue struct {
	Set   bool   `json:"set,omitempty"`
	Value string `json:"value,omitempty"`
}

type QAValue struct {
	Id int `json:"id,omitempty"`
}

type QuestionUser struct {
	UserId   int    `json:"id,omitempty"`
	UserName string `json:"name,omitempty"`
	UserMail string `json:"mail,omitempty"`
	UserIcon string `json:"icon,omitempty"`
}

type QuestionDetail struct {
	QuestionText string `json:"text,omitempty"`
}

// DataQuestion 問題のstruct
type DataQuestion struct {
	QuestionId          int    `json:"id,omitempty"`       // 問題ID
	QuestionTitle       string `json:"title,omitempty"`    // 問題名
	QuestionUser        `json:"user,omitempty"`            // 作成者情報
	QuestionCounter     `json:"questionView,omitempty"`    // 閲覧数
	QuestionLang        string `json:"language,omitempty"` // 解答言語
	QuestionDetail      `json:"body,omitempty"`            // 問題詳細
	QAValue             `json:"qa"`                        // 模範解答
	QuestionInputValue  `json:"modelInput,omitempty"`      // 入力値
	QuestionOutputValue `json:"modelOutput,omitempty"`     // 出力値
	CreateAt            string `json:"createAt,omitempty"` // 作成日
	UpdateAt            string `json:"updateAt"`           // 更新日
	ErrorFlag           bool   `json:"error_flag"`
}
