package penjadwalan

import (
	"time"

	"github.com/google/uuid"
)

type Penjadwalan struct {
	ID uuid.UUID
	IDPengurus uuid.UUID
	WaktuPertama time.Time
	AntrianPertama string
	WaktuKedua time.Time
	AntrianKedua string
	WaktuKetiga time.Time
	AntrianKetiga string
}