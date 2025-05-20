package handler

import (
	"masjid/helper"
	"masjid/pengurus"
	"masjid/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type pengurusHandler struct {
	pengurusService pengurus.Service
}

func NewPengurusHandler(pengurusService pengurus.Service) *pengurusHandler {
	return &pengurusHandler{pengurusService}
}

func (h *pengurusHandler) RegisterPengurus(c *gin.Context) {
	var input pengurus.RegisterMasjid

	// Ambil data user yang sedang login dari middleware
	currentUser, exists := c.Get("currentUser")
	if !exists {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	userData := currentUser.(user.User) // Konversi interface{} ke struct user.User
	if userData.Role != "Pengurus" {
		response := helper.APIResponse("Forbidden: Only Pengurus can access", http.StatusForbidden, "error", nil)
		c.JSON(http.StatusForbidden, response)
		return
	}

	// Binding request JSON ke struct input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Register pengurus failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}


	newPengurus, err := h.pengurusService.DaftarMasjid(userData.IDUser, input)
	if err != nil {
		response := helper.APIResponse("Register pengurus failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := pengurus.FomatPengurus(newPengurus)
	response := helper.APIResponse("Pengurus has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *pengurusHandler) GetPengurusByUserID(c *gin.Context) {
	currentUser, exists := c.Get("currentUser")
	if !exists {
		response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
		c.JSON(http.StatusUnauthorized, response)
		return
	}

	userData := currentUser.(user.User)

	if userData.Role != "Pengurus" {
		response := helper.APIResponse("Forbidden: Only Pengurus can access", http.StatusForbidden, "error", nil)
		c.JSON(http.StatusForbidden, response)
		return
	}

	pengurusList, err := h.pengurusService.GetPengurusByUserID(userData.IDUser.String())
	if err != nil {
		response := helper.APIResponse("Failed to get pengurus data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	
	formattedPengurusList := pengurus.FormatPengurusList(pengurusList)
	response := helper.APIResponse("Success retrieving pengurus data", http.StatusOK, "success", formattedPengurusList)
	c.JSON(http.StatusOK, response)
}

func (h *pengurusHandler) SearchMasjid(c *gin.Context) {
	nama := c.Query("nama")

	if nama == "" {
		response := helper.APIResponse("Nama masjid tidak boleh kosong", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	masjids, err := h.pengurusService.FindMasjidByInput(nama)
	if err != nil {
		response := helper.APIResponse("Gagal mengambil data masjid", http.StatusInternalServerError, "error", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	formatted := pengurus.FormatPengurusList(masjids)
	response := helper.APIResponse("Data masjid ditemukan", http.StatusOK, "success", formatted)
	c.JSON(http.StatusOK, response)
}
