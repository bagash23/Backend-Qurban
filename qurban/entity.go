package qurban

import "github.com/google/uuid"

type Qurban struct {
	ID uuid.UUID
	IDPengurus uuid.UUID
	NamaPemberi string
	KategoriHewan string
	JumlahHewan string
	Status string
	TanggalPendaftaran string
	TanggalPenyembelihan string
	Image []Images
}

type Images struct {
	ID       uuid.UUID
	QurbanID uuid.UUID
	FileName string    
	FileURL  string    
}