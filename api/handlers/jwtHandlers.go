package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func MainJwt(c echo.Context) error  {

	user := c.Get("user")

	token := user.(*jwt.Token)


	claim := token.Claims.(jwt.MapClaims)


	//log.Println(claim)
	log.Println("username", claim["name"].(string), "User id: ", claim["jti"])

	return c.String(http.StatusOK, "you are on jwt page.")
}

