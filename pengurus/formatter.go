package pengurus

type PengurusFormatter struct {
	NamaMasjid string `json:"nama_masjid"`
	NomorPengurus string `json:"nomor_pengurus"`
	AlamatMasjid string `json:"alamat_masjid"`
	KotaMasjid string `json:"kota_masjid"`
	KodePos string `json:"kode_pos"`
	ProvinsiMasjid string `json:"provinsi_masjid"`
}

func FomatPengurus(pengurus Pengurus) PengurusFormatter {
	formatter := PengurusFormatter {
		NamaMasjid: pengurus.NamaMasjid,
		NomorPengurus: pengurus.NomorPengurus,
		AlamatMasjid: pengurus.AlamatMasjid,
		KotaMasjid: pengurus.KotaMasjid,
		KodePos: pengurus.KodePos,
		ProvinsiMasjid: pengurus.ProvinsiMasjid,
	}
	return formatter
}

func FormatPengurusList(pengurusList []Pengurus) []PengurusFormatter {
	formattedList := []PengurusFormatter{}

	for _, pengurus := range pengurusList {
		formattedList = append(formattedList, FomatPengurus(pengurus))
	}

	return formattedList
}