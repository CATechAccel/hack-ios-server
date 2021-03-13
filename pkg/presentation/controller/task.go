package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Task struct {
}

func NewTask() *Task {
	return &Task{}
}

type CreateTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type mockCreateTaskResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (t *Task) HandleCreateTask(c echo.Context) error {
	req := new(CreateTaskRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	mockRes := &mockCreateTaskResponse{
		Name:        req.Name,
		Description: req.Description,
	}
	return c.JSON(http.StatusOK, mockRes)
}
