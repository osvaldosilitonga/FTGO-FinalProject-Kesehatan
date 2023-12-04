package middlewares

import (
	"fmt"
	"gateway/utils"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("authorization")

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("failed to verify token signature")
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return utils.ErrorMessage(c, &utils.ApiUnauthorized, err.Error())
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// c.Set("user", claims)
			c.Set("userId", claims["id"])
			c.Set("role", claims["role"])
			c.Set("email", claims["email"])

			return next(c)
		}
		return utils.ErrorMessage(c, &utils.ApiUnauthorized, "Please Login First")
	}
}
