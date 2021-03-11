package presentation

import (
	"log"

	"github.com/ari1021/hack-ios-server/pkg/application"
	"github.com/ari1021/hack-ios-server/pkg/infrastructure"
	"github.com/ari1021/hack-ios-server/pkg/presentation/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// NewEcho は，routing処理などを行い，*echo.Echoを返します．
func NewEcho() *echo.Echo {
	conn, err := infrastructure.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	ur := infrastructure.NewUserRepository(conn)
	ua := application.NewUserApplication(ur)
	uc := controller.NewUser(ua)
	e.POST("/users", uc.HandleCreateUser)
	e.POST("/login", uc.HandleLogin)
	return e
}
