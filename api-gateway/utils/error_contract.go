package utils

import (
	"gateway/models/web"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ApiBadRequest = web.ErrWebResponse{
		Code:   http.StatusBadRequest,
		Status: "Bad Request",
	}

	ApiNotFound = web.ErrWebResponse{
		Code:   http.StatusNotFound,
		Status: "not found",
	}

	ApiInternalServer = web.ErrWebResponse{
		Code:   http.StatusInternalServerError,
		Status: "internal server error",
	}
)

func ErrorMessage(c echo.Context, apiErr *web.ErrWebResponse, detail any) error {
	apiErr.Detail = detail

	return c.JSON(apiErr.Code, apiErr)
}
