package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/models/web"
	"io"
	"net/http"
	"os"
)

type User interface {
	Login(req *web.UsersLoginRequest) (*web.HttpUserLogin, int, error)
	Register(req *web.UsersRegisterRequest) (*web.HttpUserRegister, int, error)
	RegisterAdmin(req *web.UsersRegisterRequest) (*web.HttpUserRegister, int, error)
}

type UserImpl struct{}

func NewUserService() *UserImpl {
	return &UserImpl{}
}

func (u *UserImpl) Login(data *web.UsersLoginRequest) (*web.HttpUserLogin, int, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	baseUrl := os.Getenv("USER_SERVICE_BASE_URL")
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/login", baseUrl), bytes.NewBuffer(d))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	stringBody := string(body)

	user := web.HttpUserLogin{}

	err = json.Unmarshal([]byte(stringBody), &user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &user, resp.StatusCode, nil
}

func (u *UserImpl) Register(data *web.UsersRegisterRequest) (*web.HttpUserRegister, int, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	baseUrl := os.Getenv("USER_SERVICE_BASE_URL")
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/register", baseUrl), bytes.NewBuffer(d))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	stringBody := string(body)

	user := web.HttpUserRegister{}

	err = json.Unmarshal([]byte(stringBody), &user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &user, resp.StatusCode, nil
}

func (u *UserImpl) RegisterAdmin(data *web.UsersRegisterRequest) (*web.HttpUserRegister, int, error) {
	d, err := json.Marshal(data)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	baseUrl := os.Getenv("USER_SERVICE_BASE_URL")
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/register/admin", baseUrl), bytes.NewBuffer(d))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	stringBody := string(body)

	user := web.HttpUserRegister{}

	err = json.Unmarshal([]byte(stringBody), &user)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &user, resp.StatusCode, nil
}
