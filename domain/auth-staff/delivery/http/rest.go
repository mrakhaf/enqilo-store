package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/enqilo-store/domain/auth-staff/interfaces"
	"github.com/mrakhaf/enqilo-store/models/request"
	"github.com/mrakhaf/enqilo-store/shared/common"
)

type handlerAuth struct {
	usecase    interfaces.Usecase
	repository interfaces.Repository
	Json       common.JSON
}

func AuthHandler(authRoute *echo.Group, usecase interfaces.Usecase, repository interfaces.Repository, Json common.JSON) {
	handler := handlerAuth{
		usecase:    usecase,
		repository: repository,
		Json:       Json,
	}

	authRoute.POST("/staff/login", handler.Login)
	authRoute.POST("/staff/register", handler.Register)
}

func (h *handlerAuth) Login(c echo.Context) error {
	var req request.Login

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	//check is email exist
	isEmailExist, dataUser, err := h.usecase.CheckIsUserExist(c.Request().Context(), req.PhoneNumber)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if !isEmailExist {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Email not found"})
	}

	//check password
	isPasswordCorrect := h.usecase.CheckUserPassword(c.Request().Context(), req.Password, dataUser)

	if !isPasswordCorrect {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Wrong password"})
	}

	data, err := h.usecase.Login(c.Request().Context(), req.PhoneNumber, req.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	data.Id = dataUser.Id
	data.Name = dataUser.Name

	return h.Json.FormatJson(c, http.StatusOK, "Login success", data)

}

func (h *handlerAuth) Register(c echo.Context) error {

	var req request.Register

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	//validate email
	isEmailExist, _, err := h.usecase.CheckIsUserExist(c.Request().Context(), req.PhoneNumber)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if isEmailExist {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Email already exist"})
	}

	data, err := h.usecase.Register(c.Request().Context(), req)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return h.Json.FormatJson(c, http.StatusCreated, "Register success", data)
}
