package penjadwalan

import "github.com/google/uuid"

type Service interface {
	CreatePenjadwalan(id uuid.UUID, input RegisterPenjadwalan) (Penjadwalan, error)
	GetPenjadwalanByUserID(userID string) ([]Penjadwalan, error)
	FindAllByMasjidName(input string) ([]Penjadwalan, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreatePenjadwalan(id uuid.UUID, input RegisterPenjadwalan) (Penjadwalan, error) {
	penjadwalan := Penjadwalan{
		ID: uuid.New(),
		IDPengurus: id,
		WaktuPertama: input.WaktuPertama,
		AntrianPertama: input.AntrianPertama,
		WaktuKedua: input.WaktuKedua,
		AntrianKedua: input.AntrianKedua,
		WaktuKetiga: input.WaktuKetiga,
		AntrianKetiga: input.AntrianKetiga,
	}
	newPenjadwalan, err := s.repository.CreatePenjadwalan(penjadwalan)
	if err != nil {
		return newPenjadwalan, err
	}
	return penjadwalan, nil
}

func (s *service) GetPenjadwalanByUserID(userID string) ([]Penjadwalan, error) {
	penjadwalanList, err := s.repository.GetPenjadwalanByUserID(userID)
	if err != nil {
		return nil, err
	}
	return penjadwalanList, nil
}

func (s * service) FindAllByMasjidName(input string) ([]Penjadwalan, error) {
	return s.repository.FindAllByMasjidName(input)
}