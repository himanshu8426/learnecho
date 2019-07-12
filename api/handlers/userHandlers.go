package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"github.com/labstack/echo"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func GetUser(c echo.Context) error {

	userName := c.QueryParam("name")
	userType := c.QueryParam("type")

	id := c.Param("id")

	if id == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name": userName,
			"type": userType,
		})
	}

	if id == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Your name is %s \nand your type is %s", userName, userType))
	}
	return c.String(http.StatusOK, id)
}

func AddUser(c echo.Context) error {
	user := User{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Printf("failed to load data. %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your user %#v", user)
	return c.String(http.StatusOK, "we got your user")
}
