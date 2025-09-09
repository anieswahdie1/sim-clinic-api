package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sim-clinic-api/internal/model"
	"sim-clinic-api/internal/service"
	"strconv"
)

type MasterDataHandler struct {
	masterDataService service.MasterDataService
}

func NewMasterDataHandler(masterDataService service.MasterDataService) *MasterDataHandler {
	return &MasterDataHandler{masterDataService: masterDataService}
}

// ============ LAYANAN TERAPI HANDLERS ============
func (h *MasterDataHandler) CreateLayananTerapi(c echo.Context) error {
	var request model.LayananTerapiRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	layanan, err := h.masterDataService.CreateLayananTerapi(request)
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusCreated, successResponse(layanan))
}

func (h *MasterDataHandler) GetAllLayananTerapi(c echo.Context) error {
	layanans, err := h.masterDataService.GetAllLayananTerapi()
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(layanans))
}

func (h *MasterDataHandler) GetLayananTerapiByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid ID"))
	}

	layanan, err := h.masterDataService.GetLayananTerapiByID(uint(id))
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(layanan))
}

func (h *MasterDataHandler) UpdateLayananTerapi(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid ID"))
	}

	var request model.LayananTerapiRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	layanan, err := h.masterDataService.UpdateLayananTerapi(uint(id), request)
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(layanan))
}

func (h *MasterDataHandler) DeleteLayananTerapi(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid ID"))
	}

	if err := h.masterDataService.DeleteLayananTerapi(uint(id)); err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(map[string]string{
		"message": "Layanan terapi deleted successfully",
	}))
}

// ============ RIWAYAT PENYAKIT HANDLERS ============
func (h *MasterDataHandler) CreateRiwayatPenyakit(c echo.Context) error {
	var request model.RiwayatPenyakitRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	riwayat, err := h.masterDataService.CreateRiwayatPenyakit(request)
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusCreated, successResponse(riwayat))
}

func (h *MasterDataHandler) GetAllRiwayatPenyakit(c echo.Context) error {
	riwayats, err := h.masterDataService.GetAllRiwayatPenyakit()
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(riwayats))
}

func (h *MasterDataHandler) GetRiwayatPenyakitByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid ID"))
	}

	riwayat, err := h.masterDataService.GetRiwayatPenyakitByID(uint(id))
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(riwayat))
}

func (h *MasterDataHandler) UpdateRiwayatPenyakit(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid ID"))
	}

	var request model.RiwayatPenyakitRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	riwayat, err := h.masterDataService.UpdateRiwayatPenyakit(uint(id), request)
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(riwayat))
}

func (h *MasterDataHandler) DeleteRiwayatPenyakit(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid ID"))
	}

	if err := h.masterDataService.DeleteRiwayatPenyakit(uint(id)); err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(map[string]string{
		"message": "Riwayat penyakit deleted successfully",
	}))
}

// ============ TEKNIK TERAPI HANDLERS ============
func (h *MasterDataHandler) CreateTeknikTerapi(c echo.Context) error {
	var request model.TeknikTerapiRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	teknik, err := h.masterDataService.CreateTeknikTerapi(request)
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusCreated, successResponse(teknik))
}

func (h *MasterDataHandler) GetAllTeknikTerapi(c echo.Context) error {
	teks, err := h.masterDataService.GetAllTeknikTerapi()
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(teks))
}

func (h *MasterDataHandler) GetTeknikTerapiByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid ID"))
	}

	teknik, err := h.masterDataService.GetTeknikTerapiByID(uint(id))
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(teknik))
}

func (h *MasterDataHandler) UpdateTeknikTerapi(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid ID"))
	}

	var request model.TeknikTerapiRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	teknik, err := h.masterDataService.UpdateTeknikTerapi(uint(id), request)
	if err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(teknik))
}

func (h *MasterDataHandler) DeleteTeknikTerapi(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid ID"))
	}

	if err := h.masterDataService.DeleteTeknikTerapi(uint(id)); err != nil {
		return handleServiceError(c, err)
	}

	return c.JSON(http.StatusOK, successResponse(map[string]string{
		"message": "Teknik terapi deleted successfully",
	}))
}
