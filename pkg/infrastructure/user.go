package infrastructure

import (
	"context"
	"errors"
	"time"

	"github.com/ari1021/hack-ios-server/pkg/domain/repository"
	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"index"`
	Name      string    `gorm:"UNIQUE_INDEX:compositeindex;not null"`
	Password  string    `gorm:"UNIQUE_INDEX:compositeindex;not null"`
}

// NewUser は，infrastructure.User(gormのマッピングオブジェクト)を返します．
func NewUser(userID string, userName string, password string) *User {
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
func (ur *UserRepository) CreateUser(ctx context.Context, userID string, userName string, password string) (err error) {
	user := NewUser(userID, userName, password)
	if err := ur.Conn.WithContext(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

// FindUser は，userNameとpasswordをもとにして，DBからuserを取得し，userIDを返します．
func (ur *UserRepository) FindUser(ctx context.Context, userName string, password string) (userID string, err error) {
	user := &User{}
	res := ur.Conn.WithContext(ctx).Find(user, "name = ? AND password = ?", userName, password)
	if err := res.Error; err != nil {
		return "", err
	}
	if res.RowsAffected == 0 {
		return "", errors.New("user not found")
	}
	return user.ID, nil
}
