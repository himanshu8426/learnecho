package main

import "echo/router"
func main()  {

	e := router.EchoStart()

	e.Start(":8000")
 }