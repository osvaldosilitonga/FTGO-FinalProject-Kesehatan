package utils

import (
	"net/http"
	"payment/models/web"

	"github.com/labstack/echo/v4"
)

var (
	ApiOk = web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	ApiDelete = web.WebResponse{
		Code:   http.StatusOK,
		Status: "Delete Success",
	}

	ApiCreate = web.WebResponse{
		Code:   http.StatusCreated,
		Status: "Create Success",
	}

	ApiUpdate = web.WebResponse{
		Code:   http.StatusOK,
		Status: "Update Success",
	}
)

func SuccessMessage(c echo.Context, apiSuccess *web.WebResponse, data any) error {
	apiSuccess.Data = data

	return c.JSON(apiSuccess.Code, apiSuccess)
}
