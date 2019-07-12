package middlewares

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetAdminGroupMiddlewars(adminGroup *echo.Group)  {

	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method}  ${host}${path}  ${latency_human}` + "\n",
	}))

	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (b bool, e error) {
		if username == "himanshu" && password == "himanshu12345" {
			return true, nil
		}
		return false, nil
	}))

}
