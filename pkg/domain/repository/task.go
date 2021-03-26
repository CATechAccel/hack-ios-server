package repository

import (
	"context"

	"github.com/ari1021/hack-ios-server/pkg/domain/entity"
)

type Task interface {
	CreateTask(ctx context.Context, task *entity.Task) (err error)
	FindTasksByUserID(ctx context.Context, userID string) (tasks []*entity.Task, err error)
}
