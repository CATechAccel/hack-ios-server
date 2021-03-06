package controller

import (
	"net/http"

	"github.com/ari1021/hack-ios-server/core"
	"github.com/ari1021/hack-ios-server/pkg/application"
	"github.com/labstack/echo/v4"
)

type User struct {
	UserApplication application.User
}

func NewUser(userApplication application.User) *User {
	return &User{
		UserApplication: userApplication,
	}
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// HandleCreateUser は，userを作成し，tokenを返します．
func (u *User) HandleCreateUser(c echo.Context) error {
	ctx := c.Request().Context()
	uuidGenerator := core.NewUUIDGenerator()
	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	token, err := u.UserApplication.CreateUser(ctx, uuidGenerator, req.Name, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "create user failed in UserApplication.CreateUser")
	}
	res := &CreateUserResponse{
		Token: token,
	}
	return c.JSON(http.StatusOK, res)
}

// HandleLogin は，ログイン処理を行い，tokenを返します．
func (u *User) HandleLogin(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	token, err := u.UserApplication.FindUser(ctx, req.Name, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid (userName, password)")
	}
	res := &LoginResponse{
		Token: token,
	}
	return c.JSON(http.StatusOK, res)
}
