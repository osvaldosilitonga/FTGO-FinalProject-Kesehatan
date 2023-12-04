package controllers

import (
	"context"
	"encoding/json"
	"gateway/models/entity"
	"gateway/models/web"
	"gateway/service"
	"gateway/utils"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type User interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type UserImpl struct {
	UserService service.User
	RedisClient *redis.Client
}

func NewUserController(us service.User, rc *redis.Client) *UserImpl {
	return &UserImpl{
		UserService: us,
		RedisClient: rc,
	}
}

func (u *UserImpl) Login(c echo.Context) error {
	req := web.UsersLoginRequest{}
	if err := c.Bind(&req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "bind error")
	}
	if err := c.Validate(&req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	// make request to user service
	resp, code, err := u.UserService.Login(&req)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, err.Error())
	}
	if code != 200 {
		return utils.HttpCodeError(c, code, resp.Message)
	}

	authData := entity.AuthUser{
		ID:   resp.ID,
		Role: resp.Role,
	}

	data, err := json.Marshal(authData)
	if err != nil {
		log.Println("error marshal data: ", err.Error())
	} else {
		// set data to redis
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = u.RedisClient.Set(ctx, resp.Token, data, time.Duration(24)*time.Hour).Err()
		if err != nil {
			log.Println("error set data to redis: ", err.Error())
		}
	}

	response := web.LoginResponse{
		Token: resp.Token,
	}
	return utils.SuccessMessage(c, &utils.ApiOk, response)
}

func (u *UserImpl) Register(c echo.Context) error {
	req := web.UsersRegisterRequest{}
	if err := c.Bind(&req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	// make request to user service
	resp, code, err := u.UserService.Register(&req)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, nil)
	}
	if code != 201 {
		return utils.HttpCodeError(c, code, resp.Message)
	}

	response := web.RegisterResponse{
		ID:        resp.User.ID,
		Name:      resp.User.Name,
		Email:     resp.User.Email,
		CreatedAt: resp.User.CreatedAt,
	}
	return utils.SuccessMessage(c, &utils.ApiOk, response)
}

func (u *UserImpl) RegisterAdmin(c echo.Context) error {
	req := web.UsersRegisterRequest{}
	if err := c.Bind(&req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	// make request to user service
	resp, code, err := u.UserService.RegisterAdmin(&req)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, nil)
	}
	if code != 201 {
		return utils.HttpCodeError(c, code, resp.Message)
	}

	response := web.RegisterResponse{
		ID:        resp.User.ID,
		Name:      resp.User.Name,
		Email:     resp.User.Email,
		CreatedAt: resp.User.CreatedAt,
	}
	return utils.SuccessMessage(c, &utils.ApiOk, response)
}
