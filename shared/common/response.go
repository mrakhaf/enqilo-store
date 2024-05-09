package common

import "github.com/labstack/echo/v4"

type JSON interface {
	FormatJson(ctx echo.Context, status int, message string, data interface{}) error
}
