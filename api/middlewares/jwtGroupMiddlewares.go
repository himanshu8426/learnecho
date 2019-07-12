package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetJwtMiddlewares(jwtGroup *echo.Group)  {

	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("SecretString"),
		SigningMethod:"HS512",
	}))
}
