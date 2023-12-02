package utils

import (
	"net/http"
	"payment/models/web"

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

	ApiForbidden = web.ErrWebResponse{
		Code:   http.StatusForbidden,
		Status: "forbidden",
	}

	ApiInternalServer = web.ErrWebResponse{
		Code:   http.StatusInternalServerError,
		Status: "internal server error",
	}

	ApiUnauthorized = web.ErrWebResponse{
		Code:   http.StatusUnauthorized,
		Status: "unauthorized",
	}
)

func ErrorMessage(c echo.Context, apiErr *web.ErrWebResponse, detail any) error {
	apiErr.Detail = detail

	return echo.NewHTTPError(apiErr.Code, apiErr)
}
