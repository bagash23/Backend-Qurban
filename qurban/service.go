package qurban

import (
	"fmt"

	"github.com/google/uuid"
)

type Service interface {
	CreateQurban(id uuid.UUID, input RegisterQurban) (Qurban, error)
	UpdateQurban(qurbanID uuid.UUID, idPengurus uuid.UUID, input RegisterQurban) (Qurban, error)
	GetQurbanByID(qurbanID uuid.UUID) (Qurban, error) 
	IsImageExists(filename string) (bool, error)
	FindAllByPengurusID(pengurusID uuid.UUID) ([]Qurban, error)
	FindAllByMasjidName(namaMasjid string) ([]Qurban, error)
	DeleteQurbanByID(qurbanID uuid.UUID) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateQurban(id uuid.UUID, input RegisterQurban)(Qurban, error) {
	qurban := Qurban{
		ID:                   uuid.New(),
		IDPengurus:           id,
		NamaPemberi:          input.NamaPemberi,
		KategoriHewan:        input.KategoriHewan,
		JumlahHewan:          input.JumlahHewan,
		Status:               input.Status,
		TanggalPendaftaran:   input.TanggalPendaftaran,
		TanggalPenyembelihan: input.TanggalPenyembelihan,
	}

	for _, imgInput := range input.Image {
		image := Images{
			ID:       uuid.New(),
			QurbanID: qurban.ID,
			FileName: imgInput.FileName,
			FileURL:  imgInput.FileURL,
		}
		qurban.Image = append(qurban.Image, image)
	}

	return s.repository.CreateQurban(qurban)
}


func (s *service) UpdateQurban(qurbanID uuid.UUID, idPengurus uuid.UUID, input RegisterQurban) (Qurban, error) {
	qurban, err := s.repository.FindQurbanByID(qurbanID)
	if err != nil {
		return qurban, err
	}

	if qurban.IDPengurus != idPengurus {
		return qurban, fmt.Errorf("Pengurus tidak memiliki hak untuk mengupdate qurban ini")
	}

	// Update data qurban
	qurban.NamaPemberi = input.NamaPemberi
	qurban.KategoriHewan = input.KategoriHewan
	qurban.JumlahHewan = input.JumlahHewan
	qurban.Status = input.Status
	qurban.TanggalPendaftaran = input.TanggalPendaftaran
	qurban.TanggalPenyembelihan = input.TanggalPenyembelihan

	// Update gambar jika ada
	for _, imgInput := range input.Image {
		image := Images{
			ID:       uuid.New(),
			QurbanID: qurban.ID,
			FileName: imgInput.FileName,
			FileURL:  imgInput.FileURL,
		}
		qurban.Image = append(qurban.Image, image)
	}

	return s.repository.UpdateQurban(qurban)
}

func (s *service) GetQurbanByID(qurbanID uuid.UUID) (Qurban, error) {
	return s.repository.FindQurbanByID(qurbanID)
}

func (s *service) IsImageExists(filename string) (bool, error) {
	return s.repository.IsImageExists(filename)
}

func (s *service) FindAllByPengurusID(pengurusID uuid.UUID) ([]Qurban, error) {
	qurbans, err := s.repository.FindAllByPengurusID(pengurusID)
	if err != nil {
		return nil, err
	}
	return qurbans, nil

}

func (s *service) FindAllByMasjidName(namaMasjid string) ([]Qurban, error) {
	return s.repository.FindAllByMasjidName(namaMasjid)
}

func (s *service) DeleteQurbanByID(qurbanID uuid.UUID) error {
	_, err := s.repository.FindQurbanByID(qurbanID)
	if err != nil {
		return err
	}

	// Hapus data
	err = s.repository.DeleteQurbanByID(qurbanID)
	if err != nil {
		return err
	}

	return nil
}
