package data

import (
	"context"
	"social-network/app/user/service/internal/biz"
	"social-network/app/user/service/internal/pkg/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

type User struct {
	ID           int64
	Email        string
	Username     string
	HashPassword string
	CreateAt     time.Time
	UpdateAt     time.Time
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	ph, err := util.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	user := User{Username: u.Username, Email: u.Email, HashPassword: ph}
	result := r.data.DB(ctx).Create(&user)
	return &biz.User{ID: user.ID, Username: user.Username}, result.Error
}

func (r *userRepo) FindByUsername(ctx context.Context, username string) (*biz.User, error) {
	var user User
	result := r.data.db.WithContext(ctx).Where("username = ?", username).First(&user)
	return &biz.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.HashPassword,
	}, result.Error
}
func (r *userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	var user User
	result := r.data.db.WithContext(ctx).Where("id = ?", id).First(&user)
	return &biz.User{ID: user.ID, Username: user.Username}, result.Error
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	var user User
	result := r.data.db.WithContext(ctx).Where("username = ?", u.Username).First(&user)
	return util.CheckPasswordHash(u.Password, user.HashPassword), result.Error
}
