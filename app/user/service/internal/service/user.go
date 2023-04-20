package service

import (
	"context"
	v1 "social-network/api/user/service/v1"
	"social-network/app/user/service/internal/models"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	rv, err := s.uc.Create(ctx, &models.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{
		Id:       rv.ID,
		Username: rv.Username,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserReply, error) {
	rv, err := s.uc.Get(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	friends := make([]*v1.Friendship, 0, len(rv.Friends))
	for _, f := range rv.Friends {
		friends = append(friends, &v1.Friendship{
			RequesterId: f.RequesterID,
			RequesteeId: f.RequesteeID,
			Status:      f.Status,
		})
	}
	res := &v1.GetUserReply{
		User: &v1.User{
			Id:       rv.ID,
			Username: rv.Username,
			Avatar:   rv.Avatar,
			Bio:      rv.Bio,
			Friends:  friends,
		}}
	s.log.Debug(res.String())
	return res, nil
}

func (s *UserService) VerifyPassword(ctx context.Context, req *v1.VerifyPasswordReq) (*v1.VerifyPasswordReply, error) {
	rv, err := s.uc.VerifyPassword(ctx, &models.User{Username: req.Username, Password: req.Password})
	if err != nil {
		return nil, err
	}

	return &v1.VerifyPasswordReply{
		Ok: rv,
	}, nil
}

func (s *UserService) GetUserByUsername(ctx context.Context, in *v1.GetUserByUsernameReq) (*v1.GetUserByUsernameReply, error) {
	return s.uc.GetUserByUsername(ctx, in)
}

func (s *UserService) AddFollower(ctx context.Context, in *v1.AddFollowerReq) (*v1.AddFollowerReply, error) {
	msg, err := s.uc.AddFollower(ctx, &models.User{
		ID: in.UserId,
	}, in.FollowerId)
	if err != nil {
		return nil, err
	}

	return &v1.AddFollowerReply{
		Message: msg,
	}, nil
}


func (s *UserService) ConfirmFriendship(ctx context.Context, in *v1.ConfirmFriendshipReq) (*v1.ConfirmFriendshipReply, error) {
	msg, err := s.uc.ConfirmFriendship(ctx, in.RequesterId, in.RequesteeId)
	if err != nil {
		return nil, err
	}
	return &v1.ConfirmFriendshipReply{
		Message: msg,
	}, nil
}