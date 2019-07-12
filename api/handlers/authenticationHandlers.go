package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

func Login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if username == "himanshu" && password == "1234" {
		cookie := new(http.Cookie)

		cookie.Name = "SessionID"
		cookie.Value = "some_string"

		cookie.Expires = time.Now().Add(48 * time.Hour)

		// TODO: create jwt token
		token, err := createJWTtoken()
		if err != nil {
			log.Println("error creating JWT token",err)
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		c.SetCookie(cookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message": "You were logged in!!!",
			"token": token,
		})
	}

	return c.String(http.StatusUnauthorized, "Unauthorized user")
}


type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func createJWTtoken() (string, error) {
	claims := JwtClaims{
		"Himanshu",
		jwt.StandardClaims{
			Id: "main_user_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte("SecretString"))
	if err != nil{
		return "", err
	}

	return token, nil
}