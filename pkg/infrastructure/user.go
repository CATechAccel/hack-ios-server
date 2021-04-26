package infrastructure

import (
	"context"
	"errors"
	"time"

	"github.com/ari1021/hack-ios-server/pkg/domain/repository"
	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"size:255;unique;not null"`
	Password  []byte         `gorm:"size:255;not null"`
}

// NewUser は，infrastructure.User(gormのマッピングオブジェクト)を返します．
func NewUser(userID string, userName string, password []byte) *User {
	return &User{
		ID:       userID,
		Name:     userName,
		Password: password,
	}
}

type UserRepository struct {
	Conn *gorm.DB
}

// NewUserRepository は，infrastructure.UserRepositoryを返します．
func NewUserRepository(conn *gorm.DB) repository.User {
	return &UserRepository{
		Conn: conn,
	}
}

// CreateUser は，userをDBにinsertします．
func (ur *UserRepository) CreateUser(ctx context.Context, userID string, userName string, password []byte) (err error) {
	user := NewUser(userID, userName, password)
	if err := ur.Conn.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FindPasswordByName は，userNameをもとにして，DBからuserを取得し，passwordを返します．
func (ur *UserRepository) FindPasswordByName(ctx context.Context, userName string) (password []byte, err error) {
	user := &User{}
	res := ur.Conn.WithContext(ctx).Find(user, "name = ?", userName)
	if err := res.Error; err != nil {
		return nil, err
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}
	return user.Password, nil
}

// FindPasswordByName は，userNameをもとにして，DBからuserを取得し，userIDを返します．
func (ur *UserRepository) FindUserIDByName(ctx context.Context, userName string) (userID string, err error) {
	user := &User{}
	res := ur.Conn.WithContext(ctx).Find(user, "name = ?", userName)
	if err := res.Error; err != nil {
		return "", err
	}
	if res.RowsAffected == 0 {
		return "", errors.New("user not found")
	}
	return user.ID, nil
}
