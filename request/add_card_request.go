package request

type AddMemberCard struct {
	Kode   string `json:"kode"`
	Nama   string `json:"nama"`
	UserID string `json:"user_id"`
}
