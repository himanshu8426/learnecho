package middlewares

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
)

func SetCookieGroupMiddlewares(cookieGroup *echo.Group)  {

	cookieGroup.Use(checkCookie)

}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("SessionID")

		if err != nil {

			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "you't dont have any cookie.")
			}

			log.Println(err)
			return err
		}

		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "invalid session")
	}
}
