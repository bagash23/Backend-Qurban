package penjadwalan

import "gorm.io/gorm"

type Repository interface {
	CreatePenjadwalan(penjadwalan Penjadwalan) (Penjadwalan, error)
	GetPenjadwalanByUserID(userID string) ([]Penjadwalan, error)
	FindAllByMasjidName(input string) ([]Penjadwalan, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePenjadwalan(penjadwalan Penjadwalan) (Penjadwalan, error) {
	err := r.db.Create(&penjadwalan).Error
	if err != nil {
		return penjadwalan, err
	}
	return penjadwalan, nil
}

func (r *repository) GetPenjadwalanByUserID(userID string) ([]Penjadwalan, error) {
	var penjadwalanList []Penjadwalan
	err := r.db.Where("id_pengurus = ?", userID).Find(&penjadwalanList).Error
	if err != nil {
		return nil, err
	}
	return penjadwalanList, nil
}

func (r *repository) FindAllByMasjidName(input string) ([]Penjadwalan, error) {
	var penjadwalan []Penjadwalan
	err := r.db.Joins("JOIN pengurus ON penjadwalans.id_pengurus = pengurus.id_pengurus").Where("pengurus.nama_masjid LIKE ?", "%"+input+"%").Find(&penjadwalan).Error
	if err != nil {
		return nil, err
	}
	return penjadwalan, nil
}