package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/enqilo-store/domain/product/interfaces"
	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/shared/common"
)

type handlerProduct struct {
	usecase    interfaces.Usecase
	repository interfaces.Repository
	Json       common.JSON
}

func ProductHandler(productRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, json common.JSON) {
	handler := &handlerProduct{
		usecase:    usecase,
		repository: repository,
		Json:       json,
	}

	productRoute.POST("/product", handler.CreateProduct)

}

func (h *handlerProduct) CreateProduct(c echo.Context) error {
	var req request.CreateProduct

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	id, createdAt, err := h.usecase.CreateProduct(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"id": id, "createdAt": createdAt})
}

func (h *handlerProduct) GetProduct(c echo.Context) error {

	var req request.GetProducts

	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &req); err != nil {
		return err
	}

	fmt.Println(req)

	return c.JSON(http.StatusOK, map[string]string{"message": "Get product"})
}
