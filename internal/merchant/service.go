package merchant

type MerchantService interface {
	CreateMerchant(data Merchant) error
	MasterData(search string, limit int, pageParams int) []Merchant
	MasterDataCount(search string) int64
	MasterDataAll() []Merchant
	DetailMerchant(id string) Merchant
	Update(body Merchant) error
}

type mstMtrService struct {
	trR MerchantRepository
}

func NewMerchantService(tR MerchantRepository) MerchantService {
	return &mstMtrService{
		trR: tR,
	}
}

func (s *mstMtrService) Update(body Merchant) error {
	return s.trR.Update(body)
}
func (s *mstMtrService) DetailMerchant(id string) Merchant {
	return s.trR.DetailMerchant(id)
}
func (s *mstMtrService) MasterData(search string, limit int, pageParams int) []Merchant {
	return s.trR.MasterData(search, limit, pageParams)
}
func (s *mstMtrService) MasterDataCount(search string) int64 {
	return s.trR.MasterDataCount(search)
}

func (s *mstMtrService) MasterDataAll() []Merchant {
	return s.trR.MasterDataAll()
}

func (s *mstMtrService) CreateMerchant(data Merchant) error {
	return s.trR.CreateMerchant(data)
}
