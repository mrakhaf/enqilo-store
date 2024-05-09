package json

import (
	"github.com/labstack/echo/v4"
	"github.com/mrakhaf/enqilo-store/shared/common"
)

type response struct {
}

func NewResponse() common.JSON {
	return &response{}
}

func (c *response) FormatJson(ctx echo.Context, status int, message string, data interface{}) error {
	return ctx.JSON(status, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}
