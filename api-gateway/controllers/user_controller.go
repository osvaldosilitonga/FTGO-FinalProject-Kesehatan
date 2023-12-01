package controllers

import (
	"gateway/models/web"
	"gateway/service"
	"gateway/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User interface {
	Login(c echo.Context) error
}

type UserImpl struct {
	UserService service.User
}

func NewUserController(us service.User) *UserImpl {
	return &UserImpl{
		UserService: us,
	}
}

func (u *UserImpl) Login(c echo.Context) error {
	req := &web.UsersLoginRequest{}
	if err := c.Bind(&req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid request body")
	}
	if err := c.Validate(&req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid request body")
	}

	err := u.UserService.Login(req)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, echo.Map{
			"msg": "error from user login service",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"msg": "OK",
	})
}
