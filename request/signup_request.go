package request

type SignupRequest struct {
	NoHp      string `json:"no_hp"`
	Fullname  string `json:"name"`
	Password  string `json:"password"`
	Password1 string `json:"password_confirmation"`
}

type OtpCheck struct {
	NoHp string `json:"no_hp"`
	Otp  int    `json:"otp"`
}
