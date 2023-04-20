package biz

import (
	"context"
	v1 "social-network/api/user/service/v1"
	"social-network/app/user/service/internal/models"

	"github.com/go-kratos/kratos/v2/log"
)

// type models.User struct {
// 	ID         string
// 	Username   string
// 	Password   string
// 	Avatar     string
// 	Bio        string
// 	Followers  []*Follow
// 	Followings []*Follow
// }

// type Follow struct {
// 	Follower    models.User
// 	FollowerID  string
// 	Following   models.User
// 	FollowingID string
// }

type UserRepo interface {
	CreateUser(ctx context.Context, u *models.User) (*models.User, error)
	GetUser(ctx context.Context, id string) (*models.User, error)
	VerifyPassword(ctx context.Context, u *models.User) (bool, error)
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, u *models.User) error
	AddFriend(ctx context.Context, userID, friendID string) error
	// RemoveFollower(ctx context.Context, u *models.User, followerID string) error
	//IsFollower(ctx context.Context, userID, followerID string) (bool, error)
}

type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/user"))}
}

func (uc *UserUseCase) Save(ctx context.Context, in *models.User) (*models.User, error) {
	user := &models.User{
		// ID:       rand.Uint32(),
		Username: in.Username,
		Password: in.Password,
	}
	_, err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		// todo: handle error
		return nil, err
	}
	return &models.User{
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

func (uc *UserUseCase) Create(ctx context.Context, u *models.User) (*models.User, error) {
	out, err := uc.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (uc *UserUseCase) Get(ctx context.Context, id string) (*models.User, error) {
	return uc.repo.GetUser(ctx, id)
}

func (uc *UserUseCase) VerifyPassword(ctx context.Context, u *models.User) (bool, error) {
	return uc.repo.VerifyPassword(ctx, u)
}

// Adds a follower to a user's list of followers
//
// Returns "success" on success, "false" if the follower already exists,
// and an error otherwise.
func (uc *UserUseCase) AddFollower(ctx context.Context, user *models.User, followerID string) (string, error) {
	err := uc.repo.AddFriend(ctx, user.ID, followerID)
	if err != nil {
		return "", err
	}
	return "success", nil
}
