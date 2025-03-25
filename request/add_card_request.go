package request

type AddMemberCard struct {
	Kode   string `json:"kode"`
	UserID string `json:"user_id"`
}

type CreateNewMember struct {
	NoMsn      string `json:"no_msn"`
	NmCustomer string `json:"nm_customer"`
	NoKartu    string `json:"no_kartu"`
	TglExpired string `json:"tgl_expired"`
}
