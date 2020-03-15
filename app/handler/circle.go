package handler

import (
	"app/model"
	"net/http"

	"github.com/labstack/echo"
	"github.com/thoas/go-funk"
)

func SelectCircleByUser(c echo.Context) error {
	uid := UserFromToken(c)
	user := model.FindUser(&model.User{ID: uid})
	if user.ID == 0 {
		return echo.ErrNotFound
	}
	all_circles := model.AllCircles()

	filter_circles := func(circle *model.Circle) bool {
		users := circle.Users
		for _, u := range users {
			if user == u {
				return true
			}
		}
		return false
	}
	circles := funk.Filter(all_circles, filter_circles)
	return c.JSON(http.StatusOK, circles)
}
