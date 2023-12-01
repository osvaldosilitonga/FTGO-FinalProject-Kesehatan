package utils

import (
	"github.com/labstack/echo/v4"
)

func HttpCodeError(c echo.Context, code int, message string) error {
	switch code {
	case 400:
		return ErrorMessage(c, &ApiBadRequest, message)
	case 401:
		return ErrorMessage(c, &ApiUnauthorized, message)
	case 404:
		return ErrorMessage(c, &ApiNotFound, message)
	default:
		return ErrorMessage(c, &ApiInternalServer, message)
	}
}
