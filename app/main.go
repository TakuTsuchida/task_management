package main

import (
	"app/handler"

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

	authRouter(e)
	circleRouter(e)

	return e
}

func authRouter(e *echo.Echo) {
	auth := e.Group("/auth")
	auth.Use(middleware.JWTWithConfig(handler.Config))
	auth.GET("/header", handler.Header())
}

func circleRouter(e *echo.Echo) {
	circle := e.Group("/circle")
	circle.GET("/select", handler.SelectCircleByUser)
	// circle.POST("/save/:circle_id", handler.SaveCircle)
	// circle.PUT("/update/:circle_id", handler.UpdateCircle)
	// circle.DELETE("delete/:circle_id", handler.DeleteCircle)
}
