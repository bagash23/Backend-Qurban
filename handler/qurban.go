package handler

import (
	"masjid/helper"
	"masjid/qurban"
	"masjid/user"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type qurbanHandler struct {
	qurbanService qurban.Service
}

func NewQurbanHandler(qurbanService qurban.Service) *qurbanHandler {
	return &qurbanHandler{qurbanService}
}

func (h *qurbanHandler) saveUploadedFiles(c *gin.Context, files []*multipart.FileHeader) ([]qurban.RegisterImageInput, error) {
	var images []qurban.RegisterImageInput
	uniqueFileNames := make(map[string]bool)

	for _, file := range files {
		if file.Size == 0 {
			continue
		}

		// Skip kalau filename sudah diproses di loop
		if _, exists := uniqueFileNames[file.Filename]; exists {
			continue
		}
		uniqueFileNames[file.Filename] = true

		// Cek ke DB apakah filename sudah ada
		exists, err := h.qurbanService.IsImageExists(file.Filename)
		if err != nil {
			return nil, err
		}
		if exists {
			// Lewati file yang sudah ada di DB
			continue
		}

		fileName := file.Filename
		savePath := "images/" + fileName

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			return nil, err
		}

		images = append(images, qurban.RegisterImageInput{
			FileName: fileName,
			FileURL:  "/" + savePath,
		})
	}

	return images, nil
}



func (h *qurbanHandler) RegisterQurban(c *gin.Context) {
	currentUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil))
		return
	}

	userData := currentUser.(user.User)
	if userData.Role != "Pengurus" {
		c.JSON(http.StatusForbidden, helper.APIResponse("Forbidden: Only Pengurus can access", http.StatusForbidden, "error", nil))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Failed to read form", http.StatusBadRequest, "error", nil))
		return
	}

	files := form.File["images"]
	images, err := h.saveUploadedFiles(c, files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to save file", http.StatusInternalServerError, "error", nil))
		return
	}

	input := qurban.RegisterQurban{
		NamaPemberi:          c.PostForm("nama_pemberi"),
		KategoriHewan:        c.PostForm("kategori_hewan"),
		JumlahHewan:          c.PostForm("jumlah_hewan"),
		Status:               c.PostForm("status"),
		TanggalPendaftaran:   c.PostForm("tanggal_pendaftaran"),
		TanggalPenyembelihan: c.PostForm("tanggal_penyembelihan"),
		Image:                images,
	}

	qurbanData, err := h.qurbanService.CreateQurban(userData.IDUser, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to create Qurban", http.StatusInternalServerError, "error", nil))
		return
	}

	c.JSON(http.StatusOK, helper.APIResponse("Qurban successfully created", http.StatusOK, "success", qurban.FormatQurban(qurbanData)))
}

func (h *qurbanHandler) UpdateQurban(c *gin.Context) {
	currentUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil))
		return
	}

	userData := currentUser.(user.User)
	if userData.Role != "Pengurus" {
		c.JSON(http.StatusForbidden, helper.APIResponse("Forbidden: Only Pengurus can access", http.StatusForbidden, "error", nil))
		return
	}

	qurbanID := c.Param("qurbanID")
	if qurbanID == "" {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Qurban ID is required", http.StatusBadRequest, "error", nil))
		return
	}

	parsedID, err := uuid.Parse(qurbanID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Invalid Qurban ID", http.StatusBadRequest, "error", nil))
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Failed to read form data", http.StatusBadRequest, "error", nil))
		return
	}

	files := form.File["images"]

	var images []qurban.RegisterImageInput
	if len(files) > 0 {
		images, err = h.saveUploadedFiles(c, files)
		if err != nil {
			c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to save file", http.StatusInternalServerError, "error", nil))
			return
		}
	}

	input := qurban.RegisterQurban{
		NamaPemberi:          c.PostForm("nama_pemberi"),
		KategoriHewan:        c.PostForm("kategori_hewan"),
		JumlahHewan:          c.PostForm("jumlah_hewan"),
		Status:               c.PostForm("status"),
		TanggalPendaftaran:   c.PostForm("tanggal_pendaftaran"),
		TanggalPenyembelihan: c.PostForm("tanggal_penyembelihan"),
	}

	// Hanya update gambar jika ada file baru
	if len(images) > 0 {
		input.Image = images
	}

	updatedQurban, err := h.qurbanService.UpdateQurban(parsedID, userData.IDUser, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to update Qurban", http.StatusInternalServerError, "error", nil))
		return
	}

	c.JSON(http.StatusOK, helper.APIResponse("Qurban successfully updated", http.StatusOK, "success", qurban.FormatQurban(updatedQurban)))
}

func (h *qurbanHandler) GetQurbanByPengurus(c *gin.Context) {
	currentUser, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil))
		return
	}

	userData := currentUser.(user.User)
	if userData.Role != "Pengurus" {
		c.JSON(http.StatusForbidden, helper.APIResponse("Forbidden: Only Pengurus can access", http.StatusForbidden, "error", nil))
		return
	}

	qurbans, err := h.qurbanService.FindAllByPengurusID(userData.IDUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to fetch Qurban data", http.StatusInternalServerError, "error", nil))
		return
	}

	var formattedQurbans []qurban.QurbanFormatter
	for _, q := range qurbans {
		formattedQurbans = append(formattedQurbans, qurban.FormatQurban(q))
	}

	c.JSON(http.StatusOK, helper.APIResponse("List of Qurban data", http.StatusOK, "success", formattedQurbans))
}

func (h *qurbanHandler) GetQurbanByMasjidName(c *gin.Context) {
	namaMasjid := c.Query("nama_masjid")
	if namaMasjid == "" {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Nama masjid is required", http.StatusBadRequest, "error", nil))
		return
	}

	qurbans, err := h.qurbanService.FindAllByMasjidName(namaMasjid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to fetch Qurban data", http.StatusInternalServerError, "error", nil))
		return
	}

	var formattedQurbans []qurban.QurbanFormatter
	for _, q := range qurbans {
		formattedQurbans = append(formattedQurbans, qurban.FormatQurban(q))
	}

	c.JSON(http.StatusOK, helper.APIResponse("List of Qurban data by masjid name", http.StatusOK, "success", formattedQurbans))
}

func (h *qurbanHandler) DeleteQurbanByID(c *gin.Context) {
	idParam := c.Param("id")

	// Konversi string ke UUID
	qurbanID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Qurban ID"})
		return
	}

	// Hapus Qurban
	err = h.qurbanService.DeleteQurbanByID(qurbanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Qurban deleted successfully"})
}
