package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	router := newRouter()
	router.Logger.Fatal(router.Start(":8082"))
}

func newRouter() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// auth := e.Group("/auth")
	//auth.POST("/signup", authHandler.Signup)
	//auth.POST("/signin", authHandler.Signin)
	return e
}
