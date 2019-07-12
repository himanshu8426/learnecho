package api

import "github.com/labstack/echo"
import "echo/api/handlers"

func JwtGroup(jwtGroup *echo.Group)  {
	jwtGroup.GET("/main", handlers.MainJwt)
}
