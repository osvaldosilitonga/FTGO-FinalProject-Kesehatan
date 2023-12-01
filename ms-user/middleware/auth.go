package middleware

import (
	"log"
	"miniproject/config"
	"miniproject/entity"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Token is missing"})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Invalid token"})
		}

		claims := token.Claims.(jwt.MapClaims)

		// Check if users id exist
		var user entity.User
		if err := config.DB.First(&user, claims["id"]).Error; err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Unauthorized"})
		}

		// check exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Token expired"})
		}

		path := c.Request().URL.Path
		p := strings.Split(path, "/")
		// fmt.Println(p[1], "<----- path")
		if claims["role"] != p[1] {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{"message": "Unauthorized"})
		}

		userID := int(claims["id"].(float64))
		log.Printf("User ID from token: %d", userID)

		c.Set("user", userID)

		return next(c)
	}
}
