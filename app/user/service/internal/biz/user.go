package biz

import (
	"context"
	v1 "social-network/api/user/service/v1"

	"github.com/go-kratos/kratos/v2/log"
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

type UserRepo interface {
	CreateUser(ctx context.Context, u *User) (*User, error)
	GetUser(ctx context.Context, id string) (*User, error)
	VerifyPassword(ctx context.Context, u *User) (bool, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	UpdateUser(ctx context.Context, u *User) error
	AddFollower(ctx context.Context, u *User, followerID string) error
	RemoveFollower(ctx context.Context, u *User, followerID string) error
	IsFollower(ctx context.Context, userID, followerID string) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) Save(ctx context.Context, in *User) (*User, error) {
	user := &User{
		// ID:       rand.Uint32(),
		Username: in.Username,
		Password: in.Password,
	}
	_, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		// todo: handle error
		return nil, err
	}
	return &User{
		ID: user.ID,
	}, nil
}

func (uc *UserUseCase) GetUserByUsername(ctx context.Context, in *v1.GetUserByUsernameReq) (*v1.GetUserByUsernameReply, error) {
	user, err := uc.repo.FindByUsername(ctx, in.Username)
	if err != nil {
		//todo: handle error
		return nil, err
	}
	return &v1.GetUserByUsernameReply{
		Id:       user.ID,
		Username: user.Username,
	}, nil
}

func (uc *UserUseCase) Create(ctx context.Context, u *User) (*User, error) {
	out, err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *UserUseCase) Get(ctx context.Context, id string) (*User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, u *User) (bool, error) {
	return uc.repo.VerifyPassword(ctx, u)
}
