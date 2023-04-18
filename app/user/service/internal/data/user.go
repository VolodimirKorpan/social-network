package data

import (
	"context"
	"social-network/app/user/service/internal/biz"
	"social-network/app/user/service/internal/models"
	"social-network/app/user/service/internal/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm/clause"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *models.User) (*models.User, error) {
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	user := &models.User{Username: u.Username, Password: ph}
	if err := r.data.DB(ctx).Omit(clause.Associations).Create(&user).Error; err != nil {
		return nil, err
	}
	return &models.User{ID: user.ID, Username: user.Username}, nil
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := r.data.DB(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &models.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}, nil
}
func (r *userRepo) GetUser(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	if err := r.data.DB(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	if err := r.data.DB(ctx).Model(&user).Association("Followers").Find(&user.Followers); err != nil {
		return nil, err
	}
	if err := r.data.DB(ctx).Model(&user).Association("Followings").Find(&user.Followings); err != nil {
		return nil, err
	}
	return &models.User{
		ID:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
		Bio:      user.Bio,
		Followers: user.Followers,
		Followings: user.Followings,
	}, nil
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *models.User) (bool, error) {
	var user models.User
	result := r.data.DB(ctx).Where("username = ?", u.Username).First(&user)
	return util.CheckPasswordHash(u.Password, user.Password), result.Error
}

func (r *userRepo) UpdateUser(ctx context.Context, u *models.User) error {
	return r.data.DB(ctx).Updates(&u).Error
}

func (r *userRepo) AddFollower(ctx context.Context, u *models.User, followerID string) error {
	if err := r.data.DB(ctx).Model(u).Association("Followers").Append(models.Follow{FollowerID: followerID, FollowingID: u.ID}); err != nil {
		return err
	}
	return nil
}

func (r *userRepo) RemoveFollower(ctx context.Context, u *models.User, followerID string) error {
	f := models.Follow{
		FollowerID:  followerID,
		FollowingID: u.ID,
	}
	if err := r.data.DB(ctx).Model(u).Association("Followers").Find(&f); err != nil {
		return err
	}
	if err := r.data.DB(ctx).Delete(f).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepo) IsFollower(ctx context.Context, userID, followerID string) (bool, error) {
	var f models.Follow
	if err := r.data.DB(ctx).Where("following_id = ? and follower_id = ?", userID, followerID).Find(&f).Error; err != nil {
		return false, err
	}
	return true, nil
}
