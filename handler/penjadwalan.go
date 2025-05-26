package handler

import (
	"masjid/helper"
	"masjid/penjadwalan"
	"masjid/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type penjadwalanHandler struct {
	penjadwalanService penjadwalan.Service
}

func NewPenjadwalanHandler(penjadwalanService penjadwalan.Service) *penjadwalanHandler {
	return &penjadwalanHandler{penjadwalanService}
}

func (h *penjadwalanHandler) RegisterPenjadwlan(c *gin.Context) {
	var input penjadwalan.RegisterPenjadwalan

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

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("Gagal pembuatan penjadwalan", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newPenjadwalan, err := h.penjadwalanService.CreatePenjadwalan(userData.IDUser, input)
	if err != nil {
		response := helper.APIResponse("Gagal pembuatan penjadwalan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := penjadwalan.FormatPenjadwalan(newPenjadwalan)
	response := helper.APIResponse("Penjadwalan has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *penjadwalanHandler) GetPenjadwalanByUserID(c *gin.Context) {
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

	penjadwalanList, err := h.penjadwalanService.GetPenjadwalanByUserID(userData.IDUser.String())
	if err != nil {
		response := helper.APIResponse("Failed to get pengurus data", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formattedPengurusList := penjadwalan.FormatPenjadwalanList(penjadwalanList)
	response := helper.APIResponse("Success retrieving pengurus data", http.StatusOK, "success", formattedPengurusList)
	c.JSON(http.StatusOK, response)
}

func (h *penjadwalanHandler) GetPenjadwalanByMasjidName(c *gin.Context) {
	namaMasjid := c.Query("nama_masjid")
	if namaMasjid == "" {
		c.JSON(http.StatusBadRequest, helper.APIResponse("Nama masjid is required", http.StatusBadRequest, "error", nil))
		return
	}
	penjadwalans, err := h.penjadwalanService.FindAllByMasjidName(namaMasjid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, helper.APIResponse("Failed to fetch Penjadwalan data", http.StatusInternalServerError, "error", nil))
		return
	}
	var formattedPenjadwalan []penjadwalan.PenjadwalanFormatter
	for _, q := range penjadwalans {
		formattedPenjadwalan = append(formattedPenjadwalan, penjadwalan.FormatPenjadwalan(q))
	}

	c.JSON(http.StatusOK, helper.APIResponse("List of Penjadwalan data by masjid name", http.StatusOK, "success", formattedPenjadwalan))
}