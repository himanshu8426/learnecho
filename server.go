package main

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"log"
)

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {

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

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "You hit home page")
}


func addUser(c echo.Context) error {
	user :=  User{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&user)
	if err != nil {
		log.Printf("failed to load data. %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is your user %#v", user)
	return c.String(http.StatusOK, "we got your user")
}

func mainAdmin(c echo.Context) error  {
	return c.String(http.StatusOK, "Welcome Admin")
}

////////////// Middlewares /////////////////////////

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Server/1.0")
		return next(c)
	}
}


func main() {

	e := echo.New()

	e.Use(ServerHeader)

	e.GET("/", homePage)
	e.GET("/users/:id", getUser)
	e.POST("/users", addUser)

	g := e.Group("/admin")

	//this logs to server interaction
	//g.Use(middleware.Logger())
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:`[${time_rfc3339}] ${status} ${method}  ${host}${path}  ${latency_human}` + "\n",
	}))

	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (b bool, e error) {
		if username == "himanshu" && password == "himanshu12345"{
			return true, nil
		}
		return false, nil
	}))

	g.GET("/main", mainAdmin)


	//server start
	e.Logger.Fatal(e.Start(":1323"))
}
