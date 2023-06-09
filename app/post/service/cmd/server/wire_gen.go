// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"social-network/app/post/service/internal/biz"
	"social-network/app/post/service/internal/conf"
	"social-network/app/post/service/internal/data"
	"social-network/app/post/service/internal/server"
	"social-network/app/post/service/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, registry *conf.Registry, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	db := data.NewGormClient(confData, logger)
	dataData, cleanup, err := data.NewData(db, logger)
	if err != nil {
		return nil, nil, err
	}
	postRepo := data.NewPostRepo(dataData, logger)
	postUseCase := biz.NewPostUseCase(postRepo, logger)
	postService := service.NewPostService(postUseCase, logger)
	grpcServer := server.NewGRPCServer(confServer, postService, logger)
	registrar := server.NewRegistar(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
