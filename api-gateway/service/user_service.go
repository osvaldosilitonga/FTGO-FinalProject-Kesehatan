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
