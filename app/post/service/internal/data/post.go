package data

import (
	"context"
	"social-network/app/post/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var _ biz.PostRepo = (*postRepo)(nil)

type postRepo struct {
	data *Data
	log  *log.Helper
}

type Post struct {
	gorm.Model
	ID       string `gorm:"type:uuid;primary_key;"`
	UserID   string `gorm:"not null"`
	Title    string
	Text     string
	Comments []Comment
	Likes    []Like
}

type Comment struct {
	gorm.Model
	Text   string
	UserID string `gorm:"not null"`
	PostID string `gorm:"not null"`
}

type Like struct {
	gorm.Model
	PostID string `gorm:"not null"`
	UserID string `gorm:"not null"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.NewString()
	return
}

func NewPostRepo(data *Data, logger log.Logger) biz.PostRepo {
	return &postRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/post")),
	}
}

func (p *postRepo) CreatePost(ctx context.Context, ps *biz.Post) (*biz.Post, error) {
	if err := p.data.DB(ctx).Create(&ps).Error; err != nil {
		return nil, err
	}
	return &biz.Post{ID: ps.ID}, nil
}

func (p *postRepo) GetPostByID(ctx context.Context, postId string) (*biz.Post, error) {
	var ps *Post
	if err := p.data.DB(ctx).Where("id = ?", postId).First(&ps).Error; err != nil {
		return nil, err
	}
	return &biz.Post{ID: ps.ID, Title: ps.Title, Text: ps.Text}, nil
}

func (p *postRepo) GetPostsByUserID(ctx context.Context, userId string) ([]*biz.Post, error) {
	var ps []*Post
	if err := p.data.DB(ctx).Where("user_id = ?", userId).Find(&ps).Error; err != nil {
		return nil, err
	}
	ls := make([]*biz.Post, 0)
	for _, p := range ps {
		ls = append(ls, &biz.Post{
			ID:     p.ID,
			UserID: p.UserID,
			Title:  p.Title,
			Text:   p.Text})
	}
	return ls, nil
}
