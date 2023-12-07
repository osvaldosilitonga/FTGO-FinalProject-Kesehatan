package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignJwtForGrpc() (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(0, 0, 1)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("GRPC_JWT_SECRET")))
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %v", err)
	}
	return tokenString, nil
}
