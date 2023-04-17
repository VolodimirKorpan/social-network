package service

import (
	v1 "social-network/api/post/service/v1"
	"social-network/app/post/service/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPostService)

type PostService struct {
	v1.UnimplementedPostServer

	pc  *biz.PostUseCase
	log *log.Helper
}

func NewPostService(pc *biz.PostUseCase, logger log.Logger) *PostService {
	return &PostService{
		pc:  pc,
		log: log.NewHelper(log.With(logger, "module", "service/post")),
	}
}
