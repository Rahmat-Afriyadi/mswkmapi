package outlet

import (
	"errors"
	"fmt"
	"wkm/utils"

	"gorm.io/gorm"
)

type OutletRepository interface {
	CreateOutlet(data Outlet) error
	MasterData(search string, limit int, pageParams int) []Outlet
	MasterDataCount(search string) int64
	DetailOutlet(id string) Outlet
	Update(body Outlet) error
}

type mstMtrRepository struct {
	conn *gorm.DB
}

func NewOutletRepository(conn *gorm.DB) OutletRepository {
	return &mstMtrRepository{
		conn: conn,
	}
}

func (lR *mstMtrRepository) DetailOutlet(id string) Outlet {
	mstMtr := Outlet{ID: id}
	lR.conn.Find(&mstMtr)
	return mstMtr
}

func (lR *mstMtrRepository) CreateOutlet(data Outlet) error {
	result := lR.conn.Create(&data)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}

}

func (lR *mstMtrRepository) Update(data Outlet) error {
	record := Outlet{ID: data.ID}
	lR.conn.Find(&record)
	if record.Nama == "" {
		return errors.New("maaf data tidak ada")
	}
	result := lR.conn.Save(&data)
	// lR.Model(&user).Association("Roles").Replace(newRoles)
	if result.Error != nil {
		fmt.Println("ini error ", result.Error)
		return result.Error
	} else {
		return nil
	}
}

func (lR *mstMtrRepository) MasterData(search string, limit int, pageParams int) []Outlet {
	mstMtr := []Outlet{}
	query := lR.conn.Where("kd_mdl like ? or nm_mtr like ? or no_mtr like ? or merk like ? ", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Scopes(utils.Paginate(&utils.PaginateParams{PageParams: pageParams, Limit: limit})).Find(&mstMtr)
	return mstMtr
}

func (lR *mstMtrRepository) MasterDataCount(search string) int64 {
	var mstMtr []Outlet
	query := lR.conn.Where("kd_mdl like ? or nm_mtr like ? or no_mtr like ? or merk like ? ", "%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	query.Select("no_mtr").Find(&mstMtr)
	return int64(len(mstMtr))
}
