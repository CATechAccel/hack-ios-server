package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) HandleCreateUser(c echo.Context) error {
	mockRes := &mockCreateUserResponse{
		Token: "token",
	}
	return c.JSON(http.StatusOK, mockRes)
}

type mockCreateUserResponse struct {
	Token string `json:"token"`
}
