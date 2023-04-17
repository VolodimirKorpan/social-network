package service

import (
	"context"
	v1 "social-network/api/post/service/v1"
	"social-network/app/post/service/internal/biz"
)

func (s *PostService) CreatePost(ctx context.Context, req *v1.CreatePostRequest) (*v1.CreatePostReply, error) {
	p, err := s.pc.CreatePost(ctx, &biz.Post{})
	if err != nil {
		return nil, err
	}
	return &v1.CreatePostReply{
		Id: p.ID,
	}, nil
}

func (s *PostService) GetPostByID(ctx context.Context, req *v1.GetPostByIDRequest) (*v1.GetPostByIDReply, error) {
	p, err := s.pc.GetPostByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.GetPostByIDReply{
		Id: p.ID,
		UserId: p.UserID,
		Title: p.Title,
		Text: p.Text,
	}, nil
}

func (s *PostService) GetPostsByUserID(ctx context.Context, req *v1.GetPostsByUserIDRequest) (*v1.GetPostsByUserIDReply, error) {
	ps, err := s.pc.GetPostsByUserID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	ls := make([]*v1.GetPostsByUserIDReply_Post, 0)
	for _, p := range ps {
		ls = append(ls, &v1.GetPostsByUserIDReply_Post{
			Id: p.ID,
			UserId: p.UserID,
			Title: p.Title,
			Text: p.Text,
		})
	}
	return &v1.GetPostsByUserIDReply{
		Posts: ls,
	}, nil
}