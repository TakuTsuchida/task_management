package controller

import (
	"github.com/TakuTsuchida/task_management/src/domain"
	"github.com/labstack/echo"
)

func Signup(c echo.Context) error {
	auth := new(domain.Auth)
	err := c.Bind(auth)
	if err != nil {
		return err
	}

}
