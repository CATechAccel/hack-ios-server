package application

import (
	"context"

	"github.com/ari1021/hack-ios-server/pkg/domain/repository"
	"github.com/google/uuid"
)

type User interface {
	CreateUser(ctx context.Context, u UUIDGenerator, userName string, password string) (userID string, err error)
	FindUser(ctx context.Context, userName string, password string) (userID string, err error)
}

type UserApplication struct {
	userRepository repository.User
}

// NewUserApplication は，application.Userを返します．
func NewUserApplication(userRepository repository.User) User {
	return &UserApplication{
		userRepository: userRepository,
	}
}

type UUIDGenerator interface {
	Generate() (string, error)
}

type UUID struct{}

// NewUUIDGenerator は，application.UUIDGeneratorを返します．
func NewUUIDGenerator() UUIDGenerator {
	return &UUID{}
}

// Generate は，uuidを返します．
func (u *UUID) Generate() (string, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uuid.String(), nil
}

// CreateUser は，userID(uuid)を生成し，user情報をDBにinsertし，userIDを返します．
func (ua *UserApplication) CreateUser(ctx context.Context, u UUIDGenerator, userName string, password string) (userID string, err error) {
	userID, err = u.Generate()
	if err != nil {
		return "", err
	}
	if err := ua.userRepository.CreateUser(ctx, userID, userName, password); err != nil {
		return "", err
	}
	return userID, nil
}

// FindUser は，userNameとpasswordをもとにして，DBからuser情報を取得し，userIDを返します．
func (ua *UserApplication) FindUser(ctx context.Context, userName string, password string) (userID string, err error) {
	userID, err = ua.userRepository.FindUser(ctx, userName, password)
	if err != nil {
		return "", err
	}
	return userID, nil
}
