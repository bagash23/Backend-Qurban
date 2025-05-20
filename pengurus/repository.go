package pengurus

import "gorm.io/gorm"

type Repository interface {
	SavePengurus(pengurus Pengurus)(Pengurus, error)
	GetPengurusByUserID(userID string) ([]Pengurus, error)
	FindMasjidByInput(input string) ([]Pengurus, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) SavePengurus(pengurus Pengurus)(Pengurus, error) {
	err := r.db.Create(&pengurus).Error
	if err != nil {
		return pengurus, err
	}
	return pengurus, nil
}

func (r *repository) GetPengurusByUserID(userID string) ([]Pengurus, error) {
	var pengurusList []Pengurus
	err := r.db.Where("id_pengurus = ?", userID).Find(&pengurusList).Error
	if err != nil {
		return nil, err
	}
	return pengurusList, nil
}

func (r *repository) FindMasjidByInput(input string) ([]Pengurus, error) {
	var results []Pengurus
	err := r.db.Where("nama_masjid LIKE ?", "%"+input+"%").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}