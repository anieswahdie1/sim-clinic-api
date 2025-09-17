package handler

import (
	"net/http"
	"sim-clinic-api/internal/model"
	"sim-clinic-api/internal/service"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var tagCustomerHandler = "internal.handler.customer_handler."

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

func (h *CustomerHandler) GetCustomers(ctx echo.Context) error {
	var (
		tag     = tagCustomerHandler + "GetCustomers."
		request model.RequestPagination
	)

	if err := ctx.Bind(&request); err != nil {
		logrus.Error(map[string]interface{}{
			"tag":     tag + "01",
			"payload": request,
			"error":   err,
		})

		return ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	cust, err := h.customerService.GetCustomer(request)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, successResponse(cust))

}
