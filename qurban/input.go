package qurban

type RegisterImageInput struct {
	FileName string `json:"fileName" binding:"required"`
	FileURL  string `json:"fileURL" binding:"required"`
}

type RegisterQurban struct {
	NamaPemberi string `json:"namaPemberi" binding:"required"`
	KategoriHewan string `json:"kategoriHewan" binding:"required"`
	JumlahHewan string `json:"jumlahHewan" binding:"required"`
	Status string `json:"status" binding:"required"`
	TanggalPendaftaran string `json:"tanggalPendaftaran" binding:"required"`
	TanggalPenyembelihan string `json:"tanggalPenyembelihan" binding:"required"`
	Image []RegisterImageInput `json:"images" binding:"required,dive"`
}