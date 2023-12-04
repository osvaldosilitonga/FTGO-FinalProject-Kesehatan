package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"gateway/configs"
	"gateway/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID    int    `json:"id"`
	Role  string `json:"role"`
	Email string `json:"email"`
}

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return utils.ErrorMessage(c, &utils.ApiUnauthorized, fmt.Errorf("Token is missing"))
		}

		// check token in redis
		user, err := CheckRedisToken(tokenString)
		if err == nil {
			fmt.Println("Token found in redis")
			c.Set("id", user.ID)
			c.Set("role", user.Role)
			c.Set("email", user.Email)
			return next(c)
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return utils.ErrorMessage(c, &utils.ApiUnauthorized, fmt.Errorf("Invalid token"))
		}

		claims := token.Claims.(jwt.MapClaims)

		// check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return utils.ErrorMessage(c, &utils.ApiUnauthorized, fmt.Errorf("Token expired"))
		}

		userID := int(claims["id"].(float64))

		c.Set("id", userID)
		c.Set("role", claims["role"])
		c.Set("email", claims["email"])

		return next(c)
	}
}

// Check Redis for token
func CheckRedisToken(token string) (*User, error) {
	redisClient := configs.InitRedis()

	rt := redisClient.Get(context.Background(), token)
	if err := rt.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	res, err := rt.Result()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var u User
	err = json.Unmarshal([]byte(res), &u)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &u, nil
}
