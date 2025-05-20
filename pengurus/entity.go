package pengurus

import "github.com/google/uuid"

type Pengurus struct {
	ID uuid.UUID
	IDPengurus uuid.UUID
	NamaMasjid string
	NomorPengurus string
	AlamatMasjid string
	KotaMasjid string
	KodePos string
	ProvinsiMasjid string
}