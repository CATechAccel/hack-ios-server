package infrastructure

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"index"`
	Name      string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
}

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

func (ur *UserRepository) GetDBConn() *gorm.DB {
	return ur.Conn
}

func (ur *UserRepository) CreateUser(userID string, userName string, password string) error {
	conn := ur.GetDBConn()
	user := NewUser(userID, userName, password)
	if err := conn.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) FindUser(userName string, password string) error {
	conn := ur.GetDBConn()
	user := &User{}
	if err := conn.Find(user, "name = ? AND password = ?", userName, password).Error; err != nil {
		return err
	}
	return nil
}
