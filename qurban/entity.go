package qurban

import "github.com/google/uuid"

type Qurban struct {
	ID                   uuid.UUID `gorm:"type:char(36);primaryKey"`
	IDPengurus           uuid.UUID `gorm:"type:char(36)"`
	NamaPemberi          string
	KategoriHewan        string
	JumlahHewan          string
	Status               string
	TanggalPendaftaran   string
	TanggalPenyembelihan string
	Image                []Images  `gorm:"foreignKey:QurbanID"`
}

type Images struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey"`
	QurbanID uuid.UUID `gorm:"type:char(36);index"`
	FileName string    
	FileURL  string    
}
