package outlet

type OutletService interface {
	CreateOutlet(data Outlet) error
	MasterData(search string, limit int, pageParams int) []Outlet
	MasterDataCount(search string) int64
	DetailOutlet(id string) Outlet
	Update(body Outlet) error
	Delete(id string, name string) error
}

type mstMtrService struct {
	trR OutletRepository
}

func NewOutletService(tR OutletRepository) OutletService {
	return &mstMtrService{
		trR: tR,
	}
}

func (s *mstMtrService) Update(body Outlet) error {
	return s.trR.Update(body)
}
func (s *mstMtrService) Delete(id string, name string) error {
	return s.trR.Delete(id, name)
}
func (s *mstMtrService) DetailOutlet(id string) Outlet {
	return s.trR.DetailOutlet(id)
}
func (s *mstMtrService) MasterData(search string, limit int, pageParams int) []Outlet {
	return s.trR.MasterData(search, limit, pageParams)
}
func (s *mstMtrService) MasterDataCount(search string) int64 {
	return s.trR.MasterDataCount(search)
}

func (s *mstMtrService) CreateOutlet(data Outlet) error {
	return s.trR.CreateOutlet(data)
}
