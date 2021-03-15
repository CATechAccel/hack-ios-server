package repository

import "context"

type User interface {
	CreateUser(ctx context.Context, userID string, userName string, password []byte) (err error)
	FindPasswordByName(ctx context.Context, userName string) (password []byte, err error)
	FindUserIDByName(ctx context.Context, userName string) (userID string, err error)
}
