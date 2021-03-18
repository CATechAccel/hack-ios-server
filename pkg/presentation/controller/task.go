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

type GetTaskResult struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type mockGetTaskResponse struct {
	Response []GetTaskResult
}

type TaskDoneRequest struct {
	ID string `json:"id"`
}

type mockTaskDoneResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
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

func (t *Task) HandleGetTask(c echo.Context) error {
	var TaskResultList []GetTaskResult

	TaskResultList = append(TaskResultList, GetTaskResult{
		ID:          "testID1",
		Name:        "testName1",
		Description: "testDescription1",
	})
	TaskResultList = append(TaskResultList, GetTaskResult{
		ID:          "testID2",
		Name:        "testName2",
		Description: "testDescription2",
	})
	mockRes := &mockGetTaskResponse{
		Response: TaskResultList,
	}

	return c.JSON(http.StatusOK, mockRes)
}

func (t *Task) HandleTaskDone(c echo.Context) error {
	req := new(TaskDoneRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	mockRes := &mockTaskDoneResponse{
		ID:          "doneID",
		Name:        "doneName",
		Description: "doneDescription",
		Done:        true,
	}
	return c.JSON(http.StatusOK, mockRes)
}
