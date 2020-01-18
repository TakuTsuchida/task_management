package main

import (
	"app/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	// ミドルウェアのログ、リカバーを全てに挟む
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	e.GET("/hello", handler.MyPage())

	e.Start(":8082")
}
