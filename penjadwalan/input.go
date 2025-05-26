package penjadwalan

import "time"

type RegisterPenjadwalan struct {
	WaktuPertama time.Time `json:"waktu_pertama" binding:"required"`
	AntrianPertama string `json:"antrian_pertama" binding:"required"`
	WaktuKedua time.Time `json:"waktu_kedua" binding:"required"`
	AntrianKedua string `json:"antrian_kedua" binding:"required"`
	WaktuKetiga time.Time `json:"waktu_ketiga" binding:"required"`
	AntrianKetiga string `json:"antrian_ketiga" binding:"required"`
}