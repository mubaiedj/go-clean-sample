package authentication

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/config"
	"github.com/mubaiedj/go-clean-sample/app/shared/utils/log"
)

type JwtCustomClaims struct {
	Client string `json:"client"`
	jwt.StandardClaims
}

func GetMiddlewareConfig() echo.MiddlewareFunc {
	jwtKey := GetJwtKey()

	return middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &JwtCustomClaims{},
		SigningKey: []byte(jwtKey),
	})
}

func GetJwtKey() string {
	jwtKey := config.GetString("jwt.key")
	if len(jwtKey) == 0 {
		log.Fatal("error undefined jwtKey")
	}
	return jwtKey
}

func GetClientToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.Claims)
	newClaims := claims.(*JwtCustomClaims)
	appClient := newClaims.Client
	return appClient
}
