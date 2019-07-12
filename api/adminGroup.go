package api

import "github.com/labstack/echo"
import "echo/api/handlers"

func AdminGroup(adminGroup *echo.Group )  {
	adminGroup.GET("/main", handlers.MainAdmin)
}
