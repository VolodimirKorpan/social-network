package service

import (
	"context"
	"fmt"
	v1 "social-network/api/network/interface/v1"
	"social-network/app/network/interface/internal/models"
)

func (s *NetworkInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	return s.ac.Register(ctx, req)
}

func (s *NetworkInterface) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return s.ac.Login(ctx, req)
}

func (s *NetworkInterface) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	err := s.uc.Logout(ctx, &models.User{})
	return &v1.LogoutReply{}, err
}

func (s *NetworkInterface) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	user, err := s.uc.GetUser(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	friends := make([]*v1.Friendship, 0, len(user.Friends))
	for _, f := range user.Friends {
		friends = append(friends, &v1.Friendship{
			RequesterId: f.RequesterID,
			RequesteeId: f.RequesteeID,
			Status:      f.Status,
		})
	}
	return &v1.GetUserReply{
		User: &v1.User{
			Id:       user.ID,
			Username: user.Username,
			Avatar:   user.Avatar,
			Bio:      user.Bio,
			Friends:  friends,
		},
	}, nil
}

func (s *NetworkInterface) AddFollower(ctx context.Context, req *v1.AddFollowerReq) (*v1.AddFollowerReply, error) {
	msg, err := s.uc.AddFollow(ctx, &models.User{
		ID: req.Id,
	}, req.FollowerId)
	if err != nil {
		return nil, err
	}

	return &v1.AddFollowerReply{
		Message: msg,
	}, nil
}
