package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"sim-clinic-api/internal/model"
	"sim-clinic-api/internal/service"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerService service.CustomerService) *CustomerHandler {
	return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) CreateCustomer(c echo.Context) error {
	var request model.AddCustomerRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse("Invalid request body"))
	}

	if err := request.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	payloadCustomer := model.Customer{
		Id:                 uuid.New().String(),
		CustomerName:       request.CustomerName,
		CustomerAddress:    request.CustomerAddress,
		City:               request.City,
		Gender:             request.Gender,
		PhoneNumber:        request.PhoneNumber,
		InformedConsent:    request.InformedConsent,
		CodeRegister:       "Code",
		SourceTerapistInfo: request.SourceTerapistInfo,
	}
	customer, err := h.customerService.CreateCustomer(payloadCustomer)
	if err != nil {
		return handleServiceError(c, err)
	}
	return c.JSON(http.StatusCreated, successResponse(customer))
}
