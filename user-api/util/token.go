package util

import (
	"crypto/rand"
	"github.com/OVillas/user-api/config"
	"github.com/OVillas/user-api/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"io"
	"strings"
	"time"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func CreateToken(user model.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.Id,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 6).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.SecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, model.ErrUnexpectedSigningMethod
	}

	return config.SecretKey, nil
}

func extractToken(c echo.Context) string {
	token := c.Request().Header.Get("Authorization")

	length := len(strings.Split(token, " "))
	if length == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func ExtractUserIdFromToken(c echo.Context) (string, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return "", err
	}

	permissions, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return "", model.ErrInvalidToken
	}

	idInterface, exists := permissions["id"]
	if !exists {
		return "", model.ErrIdNotFoundInPermissions
	}

	id, ok := idInterface.(string)
	if !ok {
		return "", model.ErrIdIsNotAString
	}

	if err := IsValidUUID(id); err != nil {
		return "", model.ErrInvalidId
	}

	return id, nil
}

func GenerateOTP(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
