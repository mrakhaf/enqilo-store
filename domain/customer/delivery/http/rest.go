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

func CustomerHandler(productRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, json common.JSON) {
	handler := &handlerProduct{
		usecase:    usecase,
		repository: repository,
		Json:       json,
	}

	productRoute.POST("/customer/register", handler.Register)

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

	return h.Json.FormatJson(c, http.StatusOK, "User registered successfully", data)
}
