package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Post struct {
	ID       string `gorm:"type:uuid;primary_key;"`
	UserID   string `gorm:"not null"`
	Title    string
	Text     string
	Comments []Comment
	Likes    []Like
}

type Comment struct {
	Text   string
	UserID string `gorm:"not null"`
	PostID string `gorm:"not null"`
}

type Like struct {
	PostID string `gorm:"not null"`
	UserID string `gorm:"not null"`
}

type PostRepo interface {
	CreatePost(ctx context.Context, ps *Post) (*Post, error)
	GetPostByID(ctx context.Context, postId string) (*Post, error)
	GetPostsByUserID(ctx context.Context, userId string) ([]*Post, error)
}

type PostUseCase struct {
	repo PostRepo
	log  *log.Helper
}

func NewPostUseCase(repo PostRepo, logger log.Logger) *PostUseCase {
	return &PostUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "usecase/post"))}
}

func (uc *PostUseCase) CreatePost(ctx context.Context, in *Post) (*Post, error) {
	ps := &Post{
		UserID: in.UserID,
		Title:  in.Title,
		Text:   in.Text,
	}
	p, err := uc.repo.CreatePost(ctx, ps)
	if err != nil {
		return nil, err
	}
	return &Post{
		ID: p.ID,
	}, nil
}

func (uc *PostUseCase) GetPostByID(ctx context.Context, id string) (*Post, error) {
	p, err := uc.repo.GetPostByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &Post{
		ID:     p.ID,
		UserID: p.UserID,
		Title:  p.Title,
		Text:   p.Text,
	}, nil
}

func (uc *PostUseCase) GetPostsByUserID(ctx context.Context, id string) ([]*Post, error) {
	p, err := uc.repo.GetPostsByUserID(ctx, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}
