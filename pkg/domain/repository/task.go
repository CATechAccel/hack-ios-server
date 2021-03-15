package repository

import (
	"context"
	"github.com/ari1021/hack-ios-server/pkg/domain/model"
)

type Task interface {
	CreateTask(ctx context.Context, taskID string, taskName string, taskDesc string) (err error)
	SelectTask(ctx context.Context, userID string) (list *[]model.Task, err error)
}
