package application

import (
	"context"
	"time"

	"github.com/ari1021/hack-ios-server/core"
	"github.com/ari1021/hack-ios-server/pkg/domain/entity"
	"github.com/ari1021/hack-ios-server/pkg/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	CreateUser(ctx context.Context, u core.UUIDGenerator, userName string, password string) (token string, err error)
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

// CreateUser は，userID(uuid)を生成し，user情報をDBにinsertし，tokenを返します．
func (ua *UserApplication) CreateUser(ctx context.Context, u core.UUIDGenerator, userName string, password string) (token string, err error) {
	userID, err := u.Generate()
	if err != nil {
		return "", err
	}
	// tokenを作成する
	t := entity.NewToken()
	t.SetClaim("id", userID)
	t.SetClaim("exp", time.Now().Add(time.Hour*72).Unix())
	// TODO: secret(秘密鍵)をどのようにして保持するのかを考える
	token, err = t.Sign()
	if err != nil {
		return "", err
	}
	// passwordをhash化する
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	// dbにinsertする
	if err := ua.userRepository.CreateUser(ctx, userID, userName, hashed); err != nil {
		return "", err
	}
	return token, nil
}

// FindUser は，userNameとpasswordをもとにして，DBからuser情報を取得し，tokenを返します．
func (ua *UserApplication) FindUser(ctx context.Context, userName string, password string) (token string, err error) {
	hashed, err := ua.userRepository.FindPasswordByName(ctx, userName)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(password)); err != nil {
		return "", err
	}
	userID, err := ua.userRepository.FindUserIDByName(ctx, userName)
	if err != nil {
		return "", err
	}
	// tokenを作成する
	t := entity.NewToken()
	t.SetClaim("id", userID)
	t.SetClaim("exp", time.Now().Add(time.Hour*72).Unix())
	// TODO: secret(秘密鍵)をどのようにして保持するのかを考える
	token, err = t.Sign()
	if err != nil {
		return "", err
	}
	return token, nil
}
