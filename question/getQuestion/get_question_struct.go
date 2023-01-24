package getQuestion

// 閲覧数、回答数を管理するstruct
type QuestionCounter struct {
	View   int `json:"view"`
	Answer int `json:"answer"`
}

// 標準入力値を管理するstruct
type QuestionInputValue struct {
	Set   bool   `json:"set,omitempty"`
	Value string `json:"value,omitempty"`
}

// 標準出力値を管理するstruct
type QuestionOutputValue struct {
	Set   bool   `json:"set,omitempty"`
	Value string `json:"value,omitempty"`
}

// ?? なにこれ
type QAValue struct {
	Id int `json:"id,omitempty"`
}

// 問題作成者を管理するstruct
type QuestionUser struct {
	UserId   int    `json:"id,omitempty"`
	UserName string `json:"name,omitempty"`
	UserMail string `json:"mail,omitempty"`
	UserIcon string `json:"icon"`
}

// 問題文を管理するstruct
type QuestionDetail struct {
	QuestionText string `json:"text,omitempty"`
}

// response dataとして使用するstruct
type QuestionData struct {
	QuestionId          int                            `json:"id,omitempty"`    // 問題ID
	QuestionTitle       string                         `json:"title,omitempty"` // 問題名
	QuestionUser        `json:"user,omitempty"`        // 作成者情報
	QuestionCounter     `json:"counter,omitempty"`     // 閲覧数
	QuestionLang        string                         `json:"language,omitempty"` // 解答言語
	QuestionDetail      `json:"body,omitempty"`        // 問題詳細
	QAValue             `json:"qa"`                    // 模範解答
	QuestionInputValue  `json:"modelInput,omitempty"`  // 入力値
	QuestionOutputValue `json:"modelOutput,omitempty"` // 出力値
	CreateAt            string                         `json:"createAt,omitempty"` // 作成日
	UpdateAt            string                         `json:"updateAt"`           // 更新日
	ErrorFlag           bool                           `json:"error_flag"`
}

// dbからの値を格納するstruct
type DBGetQuestionData struct {
	QuestionId     int    `json:"QuestionId,omitempty"`     // 問題ID
	QuestionTitle  string `json:"QuestionTitle,omitempty"`  // 問題title
	UserId         int    `json:"UserId,omitempty"`         // 作成者id
	UserName       string `json:"UserName,omitempty"`       // 作成者の名前
	UserMail       string `json:"UserMail,omitempty"`       // 作成者のメール
	UserIcon       string `json:"UserIcon,omitempty"`       // 作成者のアイコン
	LanguageName   string `json:"LanguageName,omitempty"`   // 問題の言語id
	QuestionDetail string `json:"QuestionDetail,omitempty"` // 問題文
	InputValue     string `json:"InputValue,omitempty"`     // 標準入力値
	OutputValue    string `json:"OutputValue,omitempty"`    // 標準出力値
	ExampleAnswer  string `json:"ExampleAnswer"`            // 模範解答
	QuestionView   int    `json:"QuestionView,omitempty"`   // 閲覧数
	AnswerCount    int    `json:"AnswerCount"`              // 回答数
	CreateAt       string `json:"CreateAt,omitempty"`       // 作成日
	UpdateAt       string `json:"UpdateAt"`                 // 更新日
}
