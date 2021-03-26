package presentation

import (
	"log"
	"net/http"

	"github.com/ari1021/hack-ios-server/pkg/application"
	"github.com/ari1021/hack-ios-server/pkg/infrastructure"
	"github.com/ari1021/hack-ios-server/pkg/presentation/controller"
	"github.com/dgrijalva/jwt-go"
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
	tr := infrastructure.NewTaskRepository(conn)
	ta := application.NewTaskApplication(tr)
	tc := controller.NewTask(ta)
	t := &controller.Task{}
	e.POST("/users", uc.HandleCreateUser)
	e.POST("/login", uc.HandleLogin)
	// tokenが正常に作動しているかの確認API
	r := e.Group("/tasks")
	// TODO: 秘密鍵の保持
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)
	r.POST("", tc.HandleCreateTask)
	r.GET("", tc.HandleGetTask)
	e.POST("/tasks/done", t.HandleTaskDone)
	return e
}

// restricted は，jwtの認証を行い，userIDをjwtから取得します．
func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["id"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
