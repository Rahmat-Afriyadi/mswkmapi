package merchant

type MerchantService interface {
	CreateMerchant(data Merchant) error
	MasterData(search string, kategori string, lokasi string, limit int, pageParams int) []Merchant
	MasterDataCount(search string, kategori string, lokasi string) int64
	MasterDataSearch(search string) []Merchant
	MasterDataAll() []Merchant
	DetailMerchant(id string, lokasi string) Merchant
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
func (s *mstMtrService) DetailMerchant(id string, lokasi string) Merchant {
	return s.trR.DetailMerchant(id, lokasi)
}
func (s *mstMtrService) MasterData(search string, kategori string, lokasi string, limit int, pageParams int) []Merchant {
	return s.trR.MasterData(search, kategori, lokasi, limit, pageParams)
}

func (s *mstMtrService) MasterDataSearch(search string) []Merchant {
	return s.trR.MasterDataSearch(search)
}
func (s *mstMtrService) MasterDataCount(search string, kategori string, lokasi string) int64 {
	return s.trR.MasterDataCount(search, kategori, lokasi)
}

func (s *mstMtrService) MasterDataAll() []Merchant {
	return s.trR.MasterDataAll()
}

func (s *mstMtrService) CreateMerchant(data Merchant) error {
	return s.trR.CreateMerchant(data)
}
