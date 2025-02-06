package request

type SigninRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	AutoLogin string `json:"auto_login"`
}
