package application

import (
	"context"
	"github.com/ari1021/hack-ios-server/pkg/domain/model"
	"github.com/ari1021/hack-ios-server/pkg/domain/repository"
)

type Task interface {
	CreateTask(ctx context.Context, taskID string, taskName string, taskDesc string) (err error)
	SelectTask(ctx context.Context, userID string) (list *[]model.Task, err error)
}

type TaskApplication struct {
	taskRepository repository.Task
}

// NewTaskApplication はapplication.Taskを返します
func NewTaskApplication(taskRepository repository.Task) Task {
	return &TaskApplication{
		taskRepository: taskRepository,
	}
}

// CreateTaskは，taskを生成しDBにinsertします
func (ta *TaskApplication) CreateTask(ctx context.Context, taskID string, taskName string, taskDesc string) (err error) {
	return err
}

// SelectTaskは，userIDに紐づくtaskをリストで返します
func (ta *TaskApplication) SelectTask(ctx context.Context, userID string) (list *[]model.Task, err error) {
	return nil, err
}
