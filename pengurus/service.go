package pengurus

import (
	"github.com/google/uuid"
)

type Service interface {
	DaftarMasjid(id uuid.UUID, input RegisterMasjid)(Pengurus, error)
	GetPengurusByUserID(userID string) ([]Pengurus, error)
	FindMasjidByInput(input string) ([]Pengurus, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DaftarMasjid(id uuid.UUID, input RegisterMasjid)(Pengurus, error) {
	pengurus := Pengurus{}

	pengurus.ID = uuid.New()
	pengurus.IDPengurus = id
	pengurus.NamaMasjid = input.NamaMasjid
	pengurus.NomorPengurus = input.NomorPengurus
	pengurus.AlamatMasjid = input.AlamatMasjid
	pengurus.KodePos = input.KodePos
	pengurus.KotaMasjid = input.KotaMasjid
	pengurus.ProvinsiMasjid = input.ProvinsiMasjid

	savedPengurus, err := s.repository.SavePengurus(pengurus)
	if err != nil {
		return pengurus, err
	}

	return savedPengurus, nil
}

func (s *service) GetPengurusByUserID(userID string) ([]Pengurus, error) {
	pengurusList, err := s.repository.GetPengurusByUserID(userID)
	if err != nil {
		return nil, err
	}
	return pengurusList, nil
}

func (s *service) FindMasjidByInput(input string) ([]Pengurus, error) {
	pengurusList, err := s.repository.FindMasjidByInput(input)
	if err != nil {
		return nil, err
	}

	return pengurusList, nil
}