package service

import (
	"bytes"
	"encoding/json"
	"gateway/models/web"
	"io"
	"net/http"
	"os"
)

type User interface {
	Login(req *web.UsersLoginRequest) error
}

type UserImpl struct{}

func NewUserService() *UserImpl {
	return &UserImpl{}
}

func (u *UserImpl) Login(data *web.UsersLoginRequest) error {
	d, err := json.Marshal(data)
	if err != nil {
		return err
	}

	baseUrl := os.Getenv("USER_SERVICE_BASE_URL")
	req, err := http.NewRequest("POST", baseUrl+"/login", bytes.NewBuffer(d))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	stringBody := string(body)

	user := web.UserLoginResponse{}
	err = json.Unmarshal([]byte(stringBody), &user)
	if err != nil {
		return err
	}

	return nil
}
