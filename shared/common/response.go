package common

import "github.com/labstack/echo/v4"

type JSON interface {
	Ok(ctx echo.Context, message string, data interface{}) error
}
