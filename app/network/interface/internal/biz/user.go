package biz

import (
	"context"
	"errors"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrUserNotFound    = errors.New("user not found")
)

type User struct {
	ID         string
	Username   string
	Password   string
	Avatar     string
	Bio        string
	Followers  []Follow
	Followings []Follow
}

type Follow struct {
	Follower    User
	FollowerID  string
	Following   User
	FollowingID string
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
		Username: username,
		Password: password,
	}, nil
}

type UserRepo interface {
	Find(ctx context.Context, id string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	Save(ctx context.Context, u *User) (string, error)

	VerifyPassword(ctx context.Context, u *User, password string) error
	AddFollow(ctx context.Context, u *User, followerID string) (string, error)
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

func (uc *UserUseCase) GetUser(ctx context.Context, id string) (*User, error) {
	return uc.repo.Find(ctx, id)
}

func (uc *UserUseCase) AddFollow(ctx context.Context, u *User, followerID string) (string, error) {
	return uc.repo.AddFollow(ctx, u, followerID)
}