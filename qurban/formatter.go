package qurban

type QurbanFormatter struct {
	NamaPemberi string `json:"nama_pemberi"`
	KategoriHewan string `json:"kategori_hewan"`
	JumlahHewan string `json:"jumlah_hewan"`
	Status string `json:"status"`
	TanggalPendaftaran string `json:"tanggal_pendaftaran"`
	TanggalPenyembelihan string `json:"tanggal_penyembelihan"`
	Image []Images `json:"images"`
}

func FormatQurban(qurban Qurban) QurbanFormatter {
	formatter := QurbanFormatter {
		NamaPemberi: qurban.NamaPemberi,
		KategoriHewan: qurban.KategoriHewan,
		JumlahHewan: qurban.JumlahHewan,
		Status: qurban.Status,
		TanggalPendaftaran: qurban.TanggalPendaftaran,
		TanggalPenyembelihan: qurban.TanggalPenyembelihan,
		Image: qurban.Image,
	}
	return formatter
}