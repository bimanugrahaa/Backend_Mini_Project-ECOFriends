package middleware

import (
	"Backend_Mini_Project-ECOFriends/config"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CreateToken(id int) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    id,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JwtSecret))
}

func ExtractTokenUserId(e echo.Context) (claims map[string]interface{}) {

	if user := e.Get("user"); user != nil {
		u := user.(*jwt.Token)
		claims = u.Claims.(jwt.MapClaims)
	}

	return
}
