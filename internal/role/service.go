package role

import "wkm/entity"

type RoleService interface {
	CreateRole(data entity.Role) error
	MasterData(search string, limit int, pageParams int) []entity.Role
	MasterDataCount(search string) int64
	DetailRole(id string) entity.Role
	Update(body entity.Role) error
	MasterDataAll() []entity.Role
}

type roleService struct {
	trR RoleRepository
}

func NewRoleService(tR RoleRepository) RoleService {
	return &roleService{
		trR: tR,
	}
}

func (s *roleService) MasterDataAll() []entity.Role {
	return s.trR.MasterDataAll()
}
func (s *roleService) Update(body entity.Role) error {
	return s.trR.Update(body)
}
func (s *roleService) DetailRole(id string) entity.Role {
	return s.trR.DetailRole(id)
}
func (s *roleService) MasterData(search string, limit int, pageParams int) []entity.Role {
	return s.trR.MasterData(search, limit, pageParams)
}
func (s *roleService) MasterDataCount(search string) int64 {
	return s.trR.MasterDataCount(search)
}

func (s *roleService) CreateRole(data entity.Role) error {
	return s.trR.CreateRole(data)
}
