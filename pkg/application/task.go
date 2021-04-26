package application

import (
	"context"

	"github.com/ari1021/hack-ios-server/core"
	"github.com/ari1021/hack-ios-server/pkg/domain/entity"
	"github.com/ari1021/hack-ios-server/pkg/domain/repository"
)

type Task interface {
	CreateTask(ctx context.Context, u core.UUIDGenerator, taskName string, taskDescription *string, userID string) (task *entity.Task, err error)
	FindTasks(ctx context.Context, userID string) (tasks []*entity.Task, err error)
	UpdateTasksDone(ctx context.Context, taskIDs []string) (tasks []*entity.Task, err error)
}

type TaskApplication struct {
	taskRepository repository.Task
}

func NewTaskApplication(taskRepository repository.Task) Task {
	return &TaskApplication{
		taskRepository: taskRepository,
	}
}

func (ta *TaskApplication) CreateTask(ctx context.Context, u core.UUIDGenerator, taskName string, taskDescription *string, userID string) (task *entity.Task, err error) {
	taskID, err := u.Generate()
	if err != nil {
		return nil, err
	}
	task = &entity.Task{
		ID:          taskID,
		Name:        taskName,
		Description: taskDescription,
		IsDone:      false,
		UserID:      userID,
	}
	if err := ta.taskRepository.CreateTask(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (ta *TaskApplication) FindTasks(ctx context.Context, userID string) (tasks []*entity.Task, err error) {
	tasks, err = ta.taskRepository.FindTasksByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (ta *TaskApplication) UpdateTasksDone(ctx context.Context, taskIDs []string) (tasks []*entity.Task, err error) {
	if err := ta.taskRepository.UpdateTaskDoneByTaskIDs(ctx, taskIDs); err != nil {
		return nil, err
	}
	tasks, err = ta.taskRepository.FindTasksByTaskIDs(ctx, taskIDs)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
