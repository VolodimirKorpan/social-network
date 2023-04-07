package service

import (
	"context"
	v1 "social-network/api/network/interface/v1"
	"social-network/app/network/internal/biz"
)


func (s *NetworkInterface) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterReply, error) {
	return s.ac.Register(ctx, req)
}

func (s *NetworkInterface) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginReply, error) {
	return s.ac.Login(ctx, req)
}
func (s *NetworkInterface) Logout(ctx context.Context, req *v1.LogoutReq) (*v1.LogoutReply, error) {
	err := s.uc.Logout(ctx, &biz.User{})
	return &v1.LogoutReply{}, err
}
