package pengurus

type RegisterMasjid struct {
    NamaMasjid string `json:"namaMasjid" binding:"required"`
    NomorPengurus string `json:"nomorPengurus" binding:"required"`
    AlamatMasjid string `json:"alamatMasjid" binding:"required"`
    KotaMasjid string `json:"kotaMasjid" binding:"required"`
    KodePos string `json:"kodePos" binding:"required"`
    ProvinsiMasjid string `json:"provinsiMasjid" binding:"required"`
}