package middleware

import (
	"github.com/OVillas/user-api/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func CheckLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		authorizationHeader := c.Request().Header.Get("Authorization")

		if authorizationHeader == "" {
			return c.NoContent(http.StatusUnauthorized)
		}

		parts := strings.Split(authorizationHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.NoContent(http.StatusUnauthorized)

		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.SecretKey), nil
		})

		if err != nil {
			return c.NoContent(http.StatusUnauthorized)
		}

		if !token.Valid {
			return c.NoContent(http.StatusUnauthorized)
		}

		return next(c)
	}
}
