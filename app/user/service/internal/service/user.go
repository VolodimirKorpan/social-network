package service

import (
	"context"
	v1 "social-network/api/user/service/v1"
	"social-network/app/user/service/internal/biz"
)

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserReply, error) {
	rv, err := s.uc.Create(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{
		Id: rv.ID,
		Username: rv.Username,
	}, nil
}