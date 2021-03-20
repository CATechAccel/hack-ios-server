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

type taskID struct {
	ID string `json:"id"`
}

type task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
}

type createTaskRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type mockCreateTaskResponse struct {
	task
}

type mockGetTaskResponse struct {
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
	mockRes := &mockCreateTaskResponse{
		task{
			ID:          "id",
			Name:        req.Name,
			Description: req.Description,
			IsDone:      false,
		},
	}
	return c.JSON(http.StatusOK, mockRes)
}

func (t *Task) HandleGetTask(c echo.Context) error {
	var tasks []*task
	tasks = append(tasks, &task{
		ID:          "testID1",
		Name:        "testName1",
		Description: "testDescription1",
		IsDone:      true,
	})
	tasks = append(tasks, &task{
		ID:          "testID2",
		Name:        "testName2",
		Description: "testDescription2",
		IsDone:      false,
	})
	mockRes := &mockGetTaskResponse{
		Tasks: tasks,
	}

	return c.JSON(http.StatusOK, mockRes)
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
			Description: "doneDescription",
			IsDone:      true,
		}
		tasks = append(tasks, task)
	}
	mockRes := mockTaskDoneResponse{
		Tasks: tasks,
	}
	return c.JSON(http.StatusOK, mockRes)
}
