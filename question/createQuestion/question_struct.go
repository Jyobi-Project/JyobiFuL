package createQuestion

// DataQuestion 問題のstruct
type DataQuestion struct {
  QuestionId     string `json:"questionId,omitempty"`   // 問題ID
  QuestionUserId string `json:"userId,omitempty"`       // 作成者ID
  QuestionTitle  string `json:"title,omitempty"`        // 問題名
  QuestionDetail string `json:"detail,omitempty"`       // 問題詳細
  InputValue     string `json:"input,omitempty"`        // 入力値
  OutputValue    string `json:"output,omitempty"`       // 出力値
  QuestionLang   string `json:"language,omitempty"`     // 解答言語
  ExampleAnswer  string `json:"answer,omitempty"`       // 模範解答
  QuestionView   string `json:"questionView,omitempty"` // 閲覧数
  CreateAt       string `json:"createAt,omitempty"`     // 作成日
  UpdateAt       string `json:"updateAt,omitempty"`     // 更新日
}
