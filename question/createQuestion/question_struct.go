package createQuestion

type QuestionData struct {
  UserId         int    `json:"userId,omitempty"`
  QuestionTitle  string `json:"title,omitempty"`
  QuestionDetail string `json:"detail,omitempty"`
  InputValue     string `json:"input,omitempty"`
  OutputValue    string `json:"output,omitempty"`
  QuestionLang   string `json:"language,omitempty"`
  // Tests []Test `json:"tests,omitempty"`
  ExampleAnswer string `json:"answer,omitempty"`
}

// テスト項目は追加機能
// type Test struct {
//   TestInput string `json:"input"`
//   TestOutput string `json:"output"`
// }
