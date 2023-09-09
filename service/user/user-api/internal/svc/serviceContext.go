package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-shop/service/user/user-api/internal/config"
	"go-zero-shop/service/user/user-rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
