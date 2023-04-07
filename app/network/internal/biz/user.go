package biz

import (
	"context"
	"errors"
	"math/rand"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrUserNotFound    = errors.New("user not found")
)

type User struct {
	ID       int64
	Username string
	Password string
}

func NewUser(
	username string,
	password string,
) (User, error) {

	// check username
	if len(username) <= 0 {
		return User{}, ErrUsernameInvalid
	}
	if len(password) <= 4 {
		return User{}, ErrPasswordInvalid
	}
	return User{
		ID:       rand.Int63(),
		Username: username,
		Password: password,
	}, nil
}

type UserRepo interface {
	Find(ctx context.Context, id int64) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Save(ctx context.Context, u *User) error

	VerifyPassword(ctx context.Context, u *User, password string) error
}

type UserUseCase struct {
	repo   UserRepo
	log    *log.Helper
	authUc *AuthUseCase
}

func NewUserUseCase(repo UserRepo, logger log.Logger, authUc *AuthUseCase) *UserUseCase {
	log := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &UserUseCase{
		repo:   repo,
		log:    log,
		authUc: authUc,
	}
}

func (uc *UserUseCase) Logout(ctx context.Context, u *User) error {
	return nil
}
