package data

import (
	"context"
	"social-network/app/user/service/internal/biz"
	"social-network/app/user/service/internal/models"
	"social-network/app/user/service/internal/pkg/util"
	"time"

	"github.com/go-kratos/kratos/v2/log"
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
	if err := r.data.DB(ctx).Create(&user).Error; err != nil {
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
	if err := r.data.DB(ctx).Model(&user).Association("Friends").Find(&user.Friends); err != nil {
		return nil, err
	}
	r.log.Debug("GetUser", "user", user.Friends)
	return &models.User{
		ID:       user.ID,
		Username: user.Username,
		Avatar:   user.Avatar,
		Bio:      user.Bio,
		Friends:  user.Friends,
	}, nil
}

func (r *userRepo) VerifyPassword(ctx context.Context, u *models.User) (bool, error) {
	var user *models.User
	result := r.data.DB(ctx).Where("username = ?", u.Username).First(&user)
	return util.CheckPasswordHash(u.Password, user.Password), result.Error
}

func (r *userRepo) UpdateUser(ctx context.Context, u *models.User) error {
	return r.data.DB(ctx).Updates(&u).Error
}

/**
 * AddFollower adds a new follower for a given user to the database.
 *
 * @param ctx: A context.Context object.
 * @param u: A pointer to a models.User object.
 * @param followerID: A string representing the ID of the follower to be added.
 *
 * @return error: An error indicating success or failure of the operation.
 */
func (r *userRepo) AddFriend(ctx context.Context, requesterUsername, requesteeUsername string) error {
	var requester models.User
	if err := r.data.DB(ctx).Where("id = ?", requesterUsername).First(&requester).Error; err != nil {
		return err
	}

	var requestee models.User
	if err := r.data.DB(ctx).Where("id = ?", requesteeUsername).First(&requestee).Error; err != nil {
		return err
	}

	friendship := models.Friendship{
		RequesterID: requester.ID,
		RequesteeID: requestee.ID,
		Status:      models.FriendshipStatusPending,
		RequestedAt: time.Now(),
	}

	if err := r.data.DB(ctx).Create(&friendship).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) ConfirmFriendship(ctx context.Context, requesterID, requesteeID string) error {
	var requester models.User
	if err := r.data.DB(ctx).Where("id = ?", requesterID).First(&requester).Error; err != nil {
		return err
	}

	var requestee models.User
	if err := r.data.DB(ctx).Where("id = ?", requesteeID).First(&requestee).Error; err != nil {
		return err
	}

	friendship := models.Friendship{}
	if err := r.data.DB(ctx).Where("requester_id = ? AND requestee_id = ?", requester.ID, requestee.ID).First(&friendship).Error; err != nil {
		return err
	}

	friendship.Status = models.FriendshipStatusAccepted
	friendship.ResolvedAt = time.Now()

	if err := r.data.DB(ctx).Save(&friendship).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) RejectFriendship(ctx context.Context, requesterID, requesteeID string) error {
	var requester models.User
	if err := r.data.DB(ctx).Where("id = ?", requesterID).First(&requester).Error; err != nil {
		return err
	}

	var requestee models.User
	if err := r.data.DB(ctx).Where("id = ?", requesteeID).First(&requestee).Error; err != nil {
		return err
	}

	friendship := models.Friendship{}
	if err := r.data.DB(ctx).Where("requester_id = ? AND requestee_id = ?", requester.ID, requestee.ID).First(&friendship).Error; err != nil {
		return err
	}

	friendship.Status = models.FriendshipStatusRejected
	friendship.ResolvedAt = time.Now()

	if err := r.data.DB(ctx).Save(&friendship).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepo) GetFriendsByID(ctx context.Context, userID string) ([]*models.User, error) {
	var friendships []*models.Friendship
	if err := r.data.DB(ctx).
		Where("requester_id = ? OR requestee_id = ?", userID, userID).
		Where("status = ? or status = ?", models.FriendshipStatusAccepted, models.FriendshipStatusPending).
		Find(&friendships).Error; err != nil {
		return nil, err
	}

	friendIDs := make([]string, 0, len(friendships))
	for _, f := range friendships {
		if f.RequesterID == userID {
			friendIDs = append(friendIDs, f.RequesteeID)
		} else {
			friendIDs = append(friendIDs, f.RequesterID)
		}
	}

	var friends []*models.User
	if err := r.data.DB(ctx).Where("id IN ?", friendIDs).Find(&friends).Error; err != nil {
		return nil, err
	}

	return friends, nil
}

// func (r *userRepo) RemoveFollower(ctx context.Context, u *models.User, followerID string) error {
// 	f := models.Follow{
// 		FollowerID:  followerID,
// 		FollowingID: u.ID,
// 	}
// 	if err := r.data.DB(ctx).Model(u).Association("Followers").Find(&f); err != nil {
// 		return err
// 	}
// 	if err := r.data.DB(ctx).Delete(f).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (r *userRepo) IsFollower(ctx context.Context, userID, friendID string) (bool, error) {
// 	var f models.Friendship
// 	if err := r.data.DB(ctx).Where("user_id = ? and friend_id = ?", friendID, userID).Find(&f).Error; err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
