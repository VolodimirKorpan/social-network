package biz

import (
	"context"
	"errors"
	"social-network/app/network/interface/internal/models"

	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrPasswordInvalid = errors.New("password invalid")
	ErrUsernameInvalid = errors.New("username invalid")
	ErrUserNotFound    = errors.New("user not found")
)

func NewUser(
	username string,
	password string,
) (models.User, error) {

	// check username
	if len(username) <= 0 {
		return models.User{}, ErrUsernameInvalid
	}
	if len(password) <= 4 {
		return models.User{}, ErrPasswordInvalid
	}
	return models.User{
		Username: username,
		Password: password,
	}, nil
}

type UserRepo interface {
	Find(ctx context.Context, id string) (*models.User, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	Save(ctx context.Context, u *models.User) (string, error)
	VerifyPassword(ctx context.Context, u *models.User, password string) error
	AddFollow(ctx context.Context, u *models.User, followerID string) (string, error)
	ConfirmFriendship(ctx context.Context, id, requesterID string) (string, error)
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

func (uc *UserUseCase) Logout(ctx context.Context, u *models.User) error {
	return nil
}

func (uc *UserUseCase) GetUser(ctx context.Context, id string) (*models.User, error) {
	return uc.repo.Find(ctx, id)
}

func (uc *UserUseCase) AddFollow(ctx context.Context, u *models.User, followerID string) (string, error) {
	return uc.repo.AddFollow(ctx, u, followerID)
}

func (uc *UserUseCase) ConfirmFriendship(ctx context.Context, id, requesterID string) (string, error) {
	return uc.repo.ConfirmFriendship(ctx, id, requesterID)
}
