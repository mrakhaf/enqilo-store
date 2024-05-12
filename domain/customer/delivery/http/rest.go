package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/enqilo-store/domain/customer/interfaces"
	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/shared/common"
)

type handlerProduct struct {
	usecase    interfaces.Usecase
	repository interfaces.Repository
	Json       common.JSON
}

func CustomerHandler(customerRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, json common.JSON) {
	handler := &handlerProduct{
		usecase:    usecase,
		repository: repository,
		Json:       json,
	}

	customerRoute.POST("/customer/register", handler.Register)
	customerRoute.GET("/customer", handler.GetCustomers)

}

func (h *handlerProduct) Register(c echo.Context) error {
	var req request.RegisterCustomer

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	data, err := h.usecase.Register(c.Request().Context(), req)

	if err != nil {
		if err.Error() == "phone number already exist" {
			return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusCreated, "User registered successfully", data)
}

func (h *handlerProduct) GetCustomers(c echo.Context) error {
	var req request.GetAllCustomerParam

	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	data, err := h.usecase.GetCustomers(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusOK, "Success get data", data)
}
