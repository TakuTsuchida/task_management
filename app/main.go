package main

import (
	"app/handler"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func newRouter() *echo.Echo {
	e := echo.New()

	// ミドルウェアのログ、リカバーを全てに挟む
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.POST("/signup", handler.Signup)
	e.POST("/login", handler.Login)

	log.Print("before Auth")
	api := e.Group("/auth")
	api.Use(middleware.JWTWithConfig(handler.Config))
	log.Print("after Auth")
	api.GET("/header", handler.Header())
	// api.GET("/header", handler.Header)

	return e
}
