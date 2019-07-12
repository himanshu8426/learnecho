package api

import (
	"echo/api/handlers"
	"github.com/labstack/echo"
)

func CookieGroup(cookieGroup *echo.Group)  {
	cookieGroup.GET("/main", handlers.MainCookie)

}
