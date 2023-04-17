package data

import (
	"context"
	"fmt"
	usV1 "social-network/api/user/service/v1"
	"social-network/app/network/interface/internal/biz"

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

func (rp *userRepo) VerifyPassword(ctx context.Context, u *biz.User, password string) error {
	_, err := rp.data.uc.VerifyPassword(ctx, &usV1.VerifyPasswordReq{
		Username: u.Username,
		Password: password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (rp *userRepo) Find(ctx context.Context, id string) (*biz.User, error) {
	result, err, _ := rp.sg.Do(fmt.Sprintf("find_user_by_id_%s", id), func() (interface{}, error) {
		user, err := rp.data.uc.GetUser(ctx, &usV1.GetUserReq{
			Id: id,
		})
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		return &biz.User{
			ID:       user.Id,
			Username: user.Username,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	return result.(*biz.User), nil
}

func (rp *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	result, err, _ := rp.sg.Do(fmt.Sprintf("find_user_by_name_%s", username), func() (interface{}, error) {
		user, err := rp.data.uc.GetUserByUsername(ctx, &usV1.GetUserByUsernameReq{
			Username: username,
		})
		if err != nil {
			return nil, biz.ErrUserNotFound
		}
		rp.log.Debug(user)
		return &biz.User{
			ID:       user.Id,
			Username: user.Username,
		}, nil
	})
	if err != nil {
		return nil, err
	}
	
	return result.(*biz.User), nil
}

func (rp *userRepo) Save(ctx context.Context, u *biz.User) (string, error) {
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


