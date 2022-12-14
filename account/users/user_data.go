package users

// ユーザ関連のストラクト

type UserData struct {
	UserName     string `json:"user_name,omitempty"`
	UserMail     string `json:"user_mail,omitempty"`
	UserPassword string `json:"user_password,omitempty"`
	UserIcon     string `json:"user_icon,omitempty"`
}
