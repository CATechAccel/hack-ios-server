package repository

import "context"

type User interface {
	CreateUser(ctx context.Context, userID string, userName string, password string) (err error)
	FindUser(ctx context.Context, userName string, password string) (userID string, err error)
}
