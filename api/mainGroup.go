package api

import (
	"github.com/labstack/echo"
	"echo/api/handlers"
)
func MainGroup(e *echo.Echo)  {

	e.GET("/login", handlers.Login)
	e.GET("/", handlers.HomePage)
	e.GET("/users/:id", handlers.GetUser)
	e.POST("/users", handlers.AddUser)
}
