package data

import (
	"context"
	"errors"
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

/*
CreateUser creates a new user in the userRepo with the given context and user data.
It hashes the user's password and stores it in the database, and returns the created
user with only its ID and username fields filled out.

Args:
- ctx (context.Context): The context for the database transaction.
- u (*models.User): The user data to create.

Returns:
  - (*models.User, error): The created user object with only ID and username fields filled out,
    or an error if there was a problem creating the user.
*/
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}

/*
CreateUser creates a new user and stores it in the database. It returns the
created user on success, or an error on failure.

Parameters:
- ctx (context.Context): The context of the operation.
- u (*models.User): The user to be created.

Returns:
- (*models.User): The created user with its ID and username filled.
- (error): An error if the operation fails.
*/
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

/*
FindByUsername retrieves a user by their username from the repository.

Parameters:
- ctx (context.Context): The context of the operation.
- username (string): The username of the user to retrieve.

Returns:
- *models.User: The user with the given username.
- error: An error if the operation fails.
*/
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

/*
GetUser retrieves a user with the given id from the database and returns a user
object with all of their friends populated.

Parameters:
- ctx (context.Context): The context of the request.
- id (string): The id of the user to retrieve.

Returns:
- (*models.User): A user object with all of their friends populated.
- (error): An error if the user cannot be retrieved.
*/
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

/*
VerifyPassword verifies if the given user's password matches the password stored
in the database.

Parameters:
- ctx (context.Context): The context in which the function is being executed.
- u (*models.User): The user whose password needs to be verified.

Returns:
- bool: A boolean indicating whether the passwords match or not.
- error: An error indicating any issues encountered while running the function.
*/

func (r *userRepo) VerifyPassword(ctx context.Context, u *models.User) (bool, error) {
	var user *models.User
	result := r.data.DB(ctx).Where("username = ?", u.Username).First(&user)
	return util.CheckPasswordHash(u.Password, user.Password), result.Error
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
func (r *userRepo) AddFriend(ctx context.Context, requesterID, requesteeID string) error {
	var requester models.User
	if err := r.data.DB(ctx).Where("id = ?", requesterID).First(&requester).Error; err != nil {
		return err
	}

	var requestee models.User
	if err := r.data.DB(ctx).Where("id = ?", requesteeID).First(&requestee).Error; err != nil {
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

/*
ConfirmFriendship updates the friendship status between two users by accepting the request.
It takes in a context, requesterID, and requesteeID as strings.
Returns an error if any operation fails.
*/
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

	// friendship.Status = models.FriendshipStatusAccepted
	// friendship.ResolvedAt = time.Now()

	if err := r.data.DB(ctx).
		Where("requester_id = ? AND requestee_id = ?", friendship.RequesterID, friendship.RequesteeID).
		Updates(&models.Friendship{Status: models.FriendshipStatusAccepted, ResolvedAt: time.Now()}).Error; err != nil {
		return err
	}

	return nil
}

/*
RejectFriendship updates the status of a friendship to rejected and sets the time of resolution. It takes in a context.Context object, the IDs of the requester and the requestee, and returns an error object.

@param ctx: a context.Context object used to carry deadlines, cancellations, and other request-scoped values across API boundaries and between processes.
@param requesterID: a string representing the ID of the user who requested the friendship.
@param requesteeID: a string representing the ID of the user who received the friendship request.

@return: an error object, which is nil if the function succeeds and returns no errors, or an error message if the function fails.
*/
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

	// friendship.Status = models.FriendshipStatusRejected
	// friendship.ResolvedAt = time.Now()

	if err := r.data.DB(ctx).
		Where("requester_id = ? AND requestee_id = ?", friendship.RequesterID, friendship.RequesteeID).
		Updates(&models.Friendship{Status: models.FriendshipStatusRejected, ResolvedAt: time.Now()}).Error; err != nil {
		return err
	}

	return nil
}

/*
GetFriendsByID retrieves the list of friends for a given user ID.

Parameters:
- ctx (context.Context): The context of the request.
- userID (string): The ID of the user to retrieve friends for.

Returns:
- ([]*models.User): A slice of User objects representing the user's friends.
- (error): An error if any occurred during the retrieval process.
*/
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

/*
RemoveFriend removes the friendship between two users specified by their IDs.
The function returns an error if there is a problem while removing the friendship.

Parameters:
- ctx (context.Context): the context of the request
- userID (string): the ID of the user initiating the remove operation
- friendID (string): the ID of the user who is being removed as a friend

Returns:
- error: an error if there is a problem while removing the friendship, or nil if successful
*/

func (r *userRepo) RemoveFriend(ctx context.Context, userID, friendID string) error {
	var friendship models.Friendship
	if err := r.data.DB(ctx).
		Where("(requester_id = ? AND requestee_id = ?) OR (requester_id = ? AND requestee_id = ?)", userID, friendID, friendID, userID).
		First(&friendship).Error; err != nil {
		return err
	}

	if friendship.Status != models.FriendshipStatusAccepted {
		return errors.New("cannot remove friendship request")
	}

	if err := r.data.DB(ctx).Delete(&friendship).Error; err != nil {
		return err
	}

	return nil
}

/*
UpdateUser updates a user with the given id in the database and returns a user
object with the updated fields.

Parameters:
- ctx (context.Context): The context of the request.
- id (string): The id of the user to update.
- updates (map[string]interface{}): The fields to update and their new values.

Returns:
- (*models.User): A user object with the updated fields.
- (error): An error if the user cannot be updated.
*/
func (r *userRepo) UpdateUser(ctx context.Context, id string, updates map[string]interface{}) (*models.User, error) {
	var user models.User
	if err := r.data.DB(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	if err := r.data.DB(ctx).Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := r.data.DB(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

/*
DeleteUser deletes a user with the given id from the database.

Parameters:
- ctx (context.Context): The context of the request.
- id (string): The id of the user to delete.

Returns:
- (error): An error if the user cannot be deleted.
*/
func (r *userRepo) DeleteUser(ctx context.Context, id string) error {
	var user models.User
	if err := r.data.DB(ctx).Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	if err := r.data.DB(ctx).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
