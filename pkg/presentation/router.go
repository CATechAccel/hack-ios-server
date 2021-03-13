package presentation

import (
	"github.com/ari1021/hack-ios-server/pkg/presentation/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewEcho() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	u := controller.NewUser()
	e.POST("/users", u.HandleCreateUser)
	e.POST("/login", u.HandleLogin)

	t := controller.NewTask()
	e.POST("/tasks", t.HandleCreateTask)
	return e
}
