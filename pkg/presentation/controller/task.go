package controller

import (
	"log"
	"net/http"

	"github.com/ari1021/hack-ios-server/core"
	"github.com/ari1021/hack-ios-server/pkg/application"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type Task struct {
	TaskApplication application.Task
}

func NewTask(taskApplication application.Task) *Task {
	return &Task{
		TaskApplication: taskApplication,
	}
}

type taskID struct {
	ID string `json:"id"`
}

type task struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	IsDone      bool    `json:"isDone"`
}

type createTaskRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CreateTaskResponse struct {
	task `json:"task"`
}

type GetTaskResponse struct {
	Tasks []*task `json:"tasks"`
}

type taskDoneRequest struct {
	TaskIDs []*taskID `json:"taskIDs"`
}

type mockTaskDoneResponse struct {
	Tasks []*task `json:"tasks"`
}

func (t *Task) HandleCreateTask(c echo.Context) error {
	req := new(createTaskRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	ctx := c.Request().Context()
	uuidGenerator := core.NewUUIDGenerator()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	te, err := t.TaskApplication.CreateTask(ctx, uuidGenerator, req.Name, req.Description, userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	res := &CreateTaskResponse{
		task{
			ID:          te.ID,
			Name:        te.Name,
			Description: te.Description,
			IsDone:      te.IsDone,
		},
	}
	return c.JSON(http.StatusOK, res)
}

func (t *Task) HandleGetTask(c echo.Context) error {
	ctx := c.Request().Context()
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["id"].(string)
	ts, err := t.TaskApplication.FindTasks(ctx, userID)
	if err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	var tasks []*task
	for _, t := range ts {
		tasks = append(tasks, &task{
			ID:          t.ID,
			Name:        t.Name,
			Description: t.Description,
			IsDone:      t.IsDone,
		})
	}
	res := &GetTaskResponse{
		Tasks: tasks,
	}

	return c.JSON(http.StatusOK, res)
}

func (t *Task) HandleTaskDone(c echo.Context) error {
	req := new(taskDoneRequest)
	if err := c.Bind(req); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}
	taskIDs := req.TaskIDs
	tasks := make([]*task, 0, len(taskIDs))
	for _, taskID := range taskIDs {
		task := &task{
			ID:          taskID.ID,
			Name:        "doneName",
			Description: nil,
			IsDone:      true,
		}
		tasks = append(tasks, task)
	}
	mockRes := mockTaskDoneResponse{
		Tasks: tasks,
	}
	return c.JSON(http.StatusOK, mockRes)
}
