package penjadwalan

import "time"

type PenjadwalanFormatter struct {
	WaktuPertama time.Time `json:"waktu_pertama"`
	AntrianPertama string `json:"antrian_pertama"`
	WaktuKedua time.Time `json:"waktu_kedua"`
	AntrianKedua string `json:"antrian_kedua"`
	WaktuKetiga time.Time `json:"waktu_ketiga"`
	AntrianKetiga string `json:"antrian_ketiga"`
}

func FormatPenjadwalan(penjadwalan Penjadwalan) PenjadwalanFormatter {
	formatter := PenjadwalanFormatter {
		WaktuPertama: penjadwalan.WaktuPertama,
		AntrianPertama: penjadwalan.AntrianPertama,
		WaktuKedua: penjadwalan.WaktuKedua,
		AntrianKedua: penjadwalan.AntrianKedua,
		WaktuKetiga: penjadwalan.WaktuKetiga,
		AntrianKetiga: penjadwalan.AntrianKetiga,
	}
	return formatter
}
func FormatPenjadwalanList(penjadwalanList []Penjadwalan) []PenjadwalanFormatter {
	formattedList := []PenjadwalanFormatter{}

	for _, penjadwalan := range penjadwalanList {
		formattedList = append(formattedList, FormatPenjadwalan(penjadwalan))
	}

	return formattedList
}