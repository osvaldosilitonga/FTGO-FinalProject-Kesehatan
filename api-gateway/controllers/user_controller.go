package controllers

import (
	"context"
	"encoding/json"
	"gateway/models/entity"
	"gateway/models/web"
	"gateway/service"
	"gateway/utils"
	"log"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type User interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
	RegisterAdmin(c echo.Context) error
	GetUserProfile(c echo.Context) error
	UpdateUserProfile(c echo.Context) error
}

type UserImpl struct {
	UserService service.User
	RedisClient *redis.Client
}

func NewUserController(us service.User, rc *redis.Client) User {
	return &UserImpl{
		UserService: us,
		RedisClient: rc,
	}
}

// @Summary 	Login
// @Description Login
// @Tags 			User
// @Accept 		json
// @Produce 	json
// @Param 		data body web.UsersLoginRequest true "User Credentials"
// @Success 	200 {object} web.SwUserLogin
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/user/login [post]
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
		ID:    resp.ID,
		Role:  resp.Role,
		Email: resp.Email,
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

// @Summary 	Register
// @Description Register new user
// @Tags 			User
// @Accept 		json
// @Produce 	json
// @Param 		data body web.UsersRegisterRequest true "User Data"
// @Success 	201 {object} web.SwUserRegister
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/user/register [post]
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
	return utils.SuccessMessage(c, &utils.ApiCreate, response)
}

// @Summary 	Register Admin
// @Description Register for admin
// @Tags 			User
// @Accept 		json
// @Produce 	json
// @Param 		data body web.UsersRegisterRequest true "Admin Data"
// @Success 	201 {object} web.SwUserRegister
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/user/register/admin [post]
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

// @Summary 	Profile (Owner Only)
// @Description Get user profile
// @Tags 			User
// @Accept 		json
// @Produce 	json
// @Param        Authorization header string true "JWT Token"
// @Param 			id path integer true "User ID"
// @Success 	200 {object} web.SwUserProfile
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/user/profile/{id} [get]
func (u *UserImpl) GetUserProfile(c echo.Context) error {
	param := c.Param("id")
	paramID, err := strconv.Atoi(param)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid user id")
	}

	userID := c.Get("id").(int)

	if paramID != userID {
		return utils.ErrorMessage(c, &utils.ApiForbidden, "Forbidden, Cannot access profile")
	}

	// make request to user service
	resp, code, err := u.UserService.GetUserProfile(userID)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, nil)
	}
	if code != 200 {
		return utils.HttpCodeError(c, code, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiOk, resp)
}

// @Summary 	Update Profile (Owner Only)
// @Description Update user profile
// @Tags 			User
// @Accept 		json
// @Produce 	json
// @Param        Authorization header string true "JWT Token"
// @Param 			id path integer true "User ID"
// @Param 		data body web.UsersUpdateProfileRequest true "User Data"
// @Success 	200 {object} web.SwUserProfileUpdate
// @Failure 	400 {object} web.ErrWebResponse
// @Failure 	401 {object} web.ErrWebResponse
// @Failure 	500 {object} web.ErrWebResponse
// @Router 		/user/profile/{id} [put]
func (u *UserImpl) UpdateUserProfile(c echo.Context) error {
	param := c.Param("id")
	paramID, err := strconv.Atoi(param)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, "invalid user id")
	}

	userID := c.Get("id").(int)

	if paramID != userID {
		return utils.ErrorMessage(c, &utils.ApiForbidden, "Forbidden, Cannot access profile")
	}

	req := web.UsersUpdateProfileRequest{}
	if err := c.Bind(&req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return utils.ErrorMessage(c, &utils.ApiBadRequest, err.Error())
	}

	// make request to user service
	resp, code, err := u.UserService.UpdateUserProfile(userID, &req)
	if err != nil {
		return utils.ErrorMessage(c, &utils.ApiInternalServer, nil)
	}
	if code != 200 {
		return utils.HttpCodeError(c, code, err.Error())
	}

	return utils.SuccessMessage(c, &utils.ApiOk, resp)
}
