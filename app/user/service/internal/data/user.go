package data

import (
	"context"
	"social-network/app/user/service/internal/biz"
	"social-network/app/user/service/internal/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	gorm.Model
	ID           string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username     string `gorm:"not null,unique_index"`
	HashPassword string `gorm:"not null"`
	Avatar       string
	Bio          string
	Followers    []Follow `gorm:"foreignkey:FollowingID"`
	Followings   []Follow `gorm:"foreignkey:FollowerID"`
}

type Follow struct {
	gorm.Model
	Follower    User
	FollowerID  string `gorm:"primaryKey" sql:"type:uuid not null"`
	Following   User
	FollowingID string `gorm:"primary_key" sql:"type:uuid not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewString()
	return
}


func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	user := User{Username: u.Username, HashPassword: ph}
	if err := r.data.DB(ctx).Omit(clause.Associations).Create(&user).Error; err != nil {
		return nil, err
	}
	return &biz.User{ID: user.ID, Username: user.Username}, nil
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	var user User
	if err := r.data.DB(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &biz.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.HashPassword,
	}, nil
}
func (r *userRepo) GetUser(ctx context.Context, id string) (*biz.User, error) {
	var user User
	if err := r.data.DB(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &biz.User{
		ID:       user.ID,
		Username: user.Username,
		Avatar: user.Avatar,
		Bio: user.Bio,
	}, nil
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	var user User
	result := r.data.DB(ctx).Where("username = ?", u.Username).First(&user)
	return util.CheckPasswordHash(u.Password, user.HashPassword), result.Error
}

func (r *userRepo) UpdateUser(ctx context.Context, u *biz.User) error {
	return r.data.DB(ctx).Updates(&u).Error
}

func (r *userRepo) AddFollower(ctx context.Context, u *biz.User, followerID string) error {
	if err := r.data.DB(ctx).Model(u).Association("Followers").Append(Follow{FollowerID: followerID, FollowingID: u.ID}); err != nil {
		return err
	}
	return nil
}

func (r *userRepo) RemoveFollower(ctx context.Context, u *biz.User, followerID string) error {
	f := Follow{
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
	var f Follow
	if err := r.data.DB(ctx).Where("following_id = ? and follower_id = ?", userID, followerID).Find(&f).Error; err != nil {
		return false, err
	}
	return true, nil
}
