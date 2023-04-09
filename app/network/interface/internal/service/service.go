package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	v1 "social-network/api/network/interface/v1"
	"social-network/app/network/interface/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewNetworkInterface)

type NetworkInterface struct {
	v1.UnimplementedNetworkInterfaceServer

	uc *biz.UserUseCase
	ac *biz.AuthUseCase

	log *log.Helper
}

func NewNetworkInterface(uc *biz.UserUseCase, ac *biz.AuthUseCase, logger log.Logger) *NetworkInterface {
	return &NetworkInterface{
		log: log.NewHelper(log.With(logger, "module", "service/inteface")),
		uc:  uc,
		ac:  ac,
	}
}
