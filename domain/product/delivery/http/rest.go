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
	productRoute.GET("/product", handler.GetProducts)
	productRoute.PUT("/product/:id", handler.UpdateProduct)
	productRoute.DELETE("/product/:id", handler.DeleteProduct)

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

	return h.Json.FormatJson(c, http.StatusCreated, "success", map[string]string{"id": id, "createdAt": createdAt})
}

func (h *handlerProduct) GetProducts(c echo.Context) error {

	var req request.GetProducts

	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	fmt.Println(req)

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	products, err := h.usecase.SearchProducts(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusOK, "success", products)
}

func (h *handlerProduct) UpdateProduct(c echo.Context) error {

	productID := c.Param("id")

	var req request.CreateProduct

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	_, err := h.repository.GetDataProductById(productID)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}

	err = h.usecase.UpdateProduct(c.Request().Context(), productID, req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "update success"})

}

func (h *handlerProduct) DeleteProduct(c echo.Context) error {
	productID := c.Param("id")

	_, err := h.repository.GetDataProductById(productID)

	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	}

	err = h.usecase.DeleteProduct(c.Request().Context(), productID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "delete success"})
}
