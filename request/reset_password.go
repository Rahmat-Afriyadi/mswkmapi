package request

type ResetPassword struct {
	IdUser             string
	PasswordLama       string `form:"password_lama" json:"password_lama"`
	Password           string `form:"password" json:"password"`
	PasswordKonfirmasi string `form:"password_konfirmasi" json:"password_konfirmasi"`
}

type ResetPasswordOtp struct {
	Token    string `form:"token" json:"token"`
	Password string `form:"password" json:"password"`
}
