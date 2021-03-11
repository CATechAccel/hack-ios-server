package application

import (
	"context"
	"time"

	"github.com/ari1021/hack-ios-server/pkg/domain/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type User interface {
	CreateUser(ctx context.Context, u UUIDGenerator, userName string, password string) (token string, err error)
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

// CreateUser は，userID(uuid)を生成し，user情報をDBにinsertし，tokenを返します．
func (ua *UserApplication) CreateUser(ctx context.Context, u UUIDGenerator, userName string, password string) (token string, err error) {
	userID, err := u.Generate()
	if err != nil {
		return "", err
	}
	// tokenを作成する
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// TODO: secret(秘密鍵)をどのようにして保持するのかを考える
	token, err = t.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	// dbにinsertする
	if err := ua.userRepository.CreateUser(ctx, userID, userName, password); err != nil {
		return "", err
	}
	return token, nil
}

// FindUser は，userNameとpasswordをもとにして，DBからuser情報を取得し，tokenを返します．
func (ua *UserApplication) FindUser(ctx context.Context, userName string, password string) (token string, err error) {
	userID, err := ua.userRepository.FindUser(ctx, userName, password)
	if err != nil {
		return "", err
	}
	// tokenを作成する
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// TODO: secret(秘密鍵)をどのようにして保持するのかを考える
	token, err = t.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return token, nil
}
