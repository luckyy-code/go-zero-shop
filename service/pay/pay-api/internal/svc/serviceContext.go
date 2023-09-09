package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-shop/service/pay/pay-api/internal/config"
	"go-zero-shop/service/pay/pay-rpc/pay"
)

type ServiceContext struct {
	Config config.Config
	PayRpc pay.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: pay.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
