package data

import (
	"context"
	"fmt"
	usV1 "social-network/api/user/service/v1"
	"social-network/app/network/interface/internal/biz"
	"social-network/app/network/interface/internal/models"

	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
	sg   *singleflight.Group
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
		sg:   &singleflight.Group{},
	}
}

func (rp *userRepo) VerifyPassword(ctx context.Context, u *models.User, password string) error {
	_, err := rp.data.uc.VerifyPassword(ctx, &usV1.VerifyPasswordReq{
		Username: u.Username,
		Password: password,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
Find retrieves a user by their unique identifier.

Parameters:
- ctx (context.Context): The context used to execute the retrieval.
- id (string): The unique identifier of the user to retrieve.

Returns:
- (*models.User): The user retrieved by the given ID.
- (error): An error if the user could not be retrieved or does not exist.
*/
func (rp *userRepo) Find(ctx context.Context, id string) (*models.User, error) {
	user, err, _ := rp.sg.Do(fmt.Sprintf("get_user_by_id_%s", id), func() (interface{}, error) {
		user, err := rp.data.uc.GetUser(ctx, &usV1.GetUserReq{
			Id: id,
		})
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		friends := make([]*models.Friendship, 0)
		for _, friend := range user.User.Friends {
			friends = append(friends, &models.Friendship{
				RequesterID: friend.RequesterId,
				RequesteeID: friend.RequesteeId,
				Status:      friend.Status,
			})
		}
		res := &models.User{
			ID:         user.User.Id,
			Username:   user.User.Username,
			Avatar:     user.User.Avatar,
			Bio:        user.User.Bio,
			Friends:    friends,
		}
		rp.log.Debug(res)
		return res, nil
	})
	if err != nil {
		return nil, err
	}
	return user.(*models.User), nil
}

func (rp *userRepo) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	result, err, _ := rp.sg.Do(fmt.Sprintf("find_user_by_name_%s", username), func() (interface{}, error) {
		user, err := rp.data.uc.GetUserByUsername(ctx, &usV1.GetUserByUsernameReq{
			Username: username,
		})
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		return &models.User{
			ID:       user.Id,
			Username: user.Username,
		}, nil
	})
	if err != nil {
		return nil, err
	}

	return result.(*models.User), nil
}

func (rp *userRepo) Save(ctx context.Context, u *models.User) (string, error) {
	user, err := rp.data.uc.CreateUser(ctx, &usV1.CreateUserReq{
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		return "", err
	}
	rp.log.Debug(user)
	return user.Id, nil
}

func (rp *userRepo) AddFollow(ctx context.Context, u *models.User, followerID string) (string, error) {
	result, err, _ := rp.sg.Do(fmt.Sprintf("add_follower_%s", followerID), func() (interface{}, error) {
		msg, err := rp.data.uc.AddFollower(ctx, &usV1.AddFollowerReq{
			FollowerId: followerID,
			UserId:     u.ID,
		})
		if err != nil {
			return nil, err
		}
		rp.log.Debug(msg)
		return msg.Message, nil
	})
	if err != nil {
		return "", err
	}
	return result.(string), nil
}
