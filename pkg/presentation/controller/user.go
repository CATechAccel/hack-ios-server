package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type mockCreateUserResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type mockLoginResponse struct {
	Token string `json:"token"`
}

func (u *User) HandleCreateUser(c echo.Context) error {
	req := new(CreateUserRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	mockRes := &mockCreateUserResponse{
		Token: req.Name + req.Password,
	}
	return c.JSON(http.StatusOK, mockRes)
}

func (u *User) HandleLogin(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	mockRes := &mockLoginResponse{
		Token: req.Name + req.Password,
	}
	return c.JSON(http.StatusOK, mockRes)
}
