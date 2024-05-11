package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/enqilo-store/domain/product/interfaces"
	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/shared/common"
	"github.com/mrakhaf/enqilo-store/shared/common/jwt"
	"github.com/mrakhaf/enqilo-store/shared/utils"
)

type handlerProduct struct {
	usecase    interfaces.Usecase
	repository interfaces.Repository
	Json       common.JSON
	JwtAccess  *jwt.JWT
}

func ProductHandler(productRoute *echo.Group, publicRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, json common.JSON, jwtAccess *jwt.JWT) {
	handler := &handlerProduct{
		usecase:    usecase,
		repository: repository,
		Json:       json,
		JwtAccess:  jwtAccess,
	}

	productRoute.POST("/product", handler.CreateProduct)
	productRoute.POST("/product/checkout", handler.Checkout)
	productRoute.GET("/product/checkout/history", handler.GetCheckoutHistory)
	publicRoute.GET("/product/customer", handler.SearchSku)
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

	isImage := utils.CheckImageType(req.ImageUrl)

	if !isImage {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "image type not supported"})
	}

	id, createdAt, err := h.usecase.CreateProduct(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusCreated, "success", map[string]string{"id": id, "createdAt": createdAt})
}

func (h *handlerProduct) Checkout(c echo.Context) error {
	var req request.Checkout

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	id, createdAt, err := h.usecase.Checkout(c.Request().Context(), req)

	switch {
	case err != nil && err.Error() == "customer account not found":
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	case err != nil && err.Error() == "product not found":
		return c.JSON(http.StatusNotFound, map[string]string{"message": err.Error()})
	case err != nil && err.Error() == "paid is less than total":
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	case err != nil && err.Error() == "change is not right":
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	case err != nil:
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusOK, "success", map[string]string{"id": id, "createdAt": createdAt})
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

	isImage := utils.CheckImageType(req.ImageUrl)

	if !isImage {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "image type not supported"})
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

func (h *handlerProduct) SearchSku(c echo.Context) error {

	var req request.SearchProductParam

	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &req); err != nil {
		return err
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	data, err := h.usecase.SearchSku(req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusOK, "Success get data", data)
}

func (h *handlerProduct) GetCheckoutHistory(c echo.Context) error {

	var req request.GetCheckoutHistoryParam

	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	data, err := h.usecase.GetCheckoutHistory(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusOK, "Success get data", data)
}
