package qurban

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateQurban(qurban Qurban) (Qurban, error)
	UpdateQurban(qurban Qurban) (Qurban, error)
	FindQurbanByID(qurbanID uuid.UUID) (Qurban, error)
	IsImageExists(filename string) (bool, error)
	FindAllByPengurusID(pengurusID uuid.UUID) ([]Qurban, error)
	FindAllByMasjidName(namaMasjid string) ([]Qurban, error)
	DeleteQurbanByID(qurbanID uuid.UUID) error

}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Fungsi untuk create Qurban
func (r *repository) CreateQurban(qurban Qurban) (Qurban, error) {
	err := r.db.Create(&qurban).Error
	if err != nil {
		return qurban, err
	}
	return qurban, nil
}

// Fungsi untuk mengupdate Qurban
func (r *repository) UpdateQurban(qurban Qurban) (Qurban, error) {
	err := r.db.Save(&qurban).Error
	if err != nil {
		return qurban, err
	}
	return qurban, nil
}

// Fungsi untuk mencari Qurban berdasarkan ID
func (r *repository) FindQurbanByID(qurbanID uuid.UUID) (Qurban, error) {
	var qurban Qurban
	err := r.db.Preload("Image").First(&qurban, "id = ?", qurbanID).Error
	if err != nil {
		return qurban, err
	}
	return qurban, nil
}


func (r *repository) IsImageExists(filename string) (bool, error) {
	var count int64
	err := r.db.Model(&Images{}).Where("file_name = ?", filename).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repository) FindAllByPengurusID(pengurusID uuid.UUID) ([]Qurban, error) {
	var qurbans []Qurban
	err := r.db.Where("id_pengurus = ?", pengurusID).Preload("Image").Find(&qurbans).Error
	if err != nil {
		return nil, err
	}
	return qurbans, nil
}

func (r *repository) FindAllByMasjidName(namaMasjid string) ([]Qurban, error) {
	var qurbans []Qurban
	err := r.db.Joins("JOIN pengurus ON qurbans.id_pengurus = pengurus.id_pengurus").
		Where("pengurus.nama_masjid LIKE ?", "%"+namaMasjid+"%").
		Preload("Image").
		Find(&qurbans).Error
	if err != nil {
		return nil, err
	}
	return qurbans, nil
}

func (r *repository) DeleteQurbanByID(qurbanID uuid.UUID) error {
	tx := r.db.Begin()

	var images []Images
	if err := tx.Where("qurban_id = ?", qurbanID).Find(&images).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Hapus file gambar dari file system
	for _, img := range images {
	filePath := filepath.Join("images", img.FileName) // gabungkan folder dengan nama file
	fmt.Println("Menghapus file:", filePath)

	err := os.Remove(filePath)
	if err != nil {
		fmt.Println("Gagal menghapus file:", filePath, "Error:", err)
		if !os.IsNotExist(err) {
			tx.Rollback()
			return fmt.Errorf("failed to delete image file: %v", err)
		}
	}
}

	// Hapus gambar terlebih dahulu
	if err := tx.Where("qurban_id = ?", qurbanID).Delete(&Images{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Hapus qurban
	if err := tx.Where("id = ?", qurbanID).Delete(&Qurban{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
