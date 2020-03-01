package handler

import (
	"app/model"
	"net/http"

	"github.com/labstack/echo"
)

func MyPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	}
}

func Header() echo.HandlerFunc {
	return func(c echo.Context) error {
		uid := UserFromToken(c)
		user_model := model.FindUser(&model.User{ID: uid})
		return c.JSON(http.StatusOK, map[string]string{
			"username": user_model.Name,
		})
	}
}
