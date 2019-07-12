package router

import (
	"echo/api"
	"echo/api/middlewares"
	"github.com/labstack/echo"
	)
func EchoStart() *echo.Echo{

	e := echo.New()

	//create groups
	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")


	//set all middlewares
	middlewares.SetMainMiddleware(e)
	middlewares.SetAdminGroupMiddlewars(adminGroup)
	middlewares.SetJwtMiddlewares(jwtGroup)
	middlewares.SetCookieGroupMiddlewares(cookieGroup)

	// set main routes
	api.MainGroup(e)

	//set group routes
	api.AdminGroup(adminGroup)
	api.JwtGroup(jwtGroup)
	api.CookieGroup(cookieGroup)

	return e
}
