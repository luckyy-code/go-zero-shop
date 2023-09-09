package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-shop/service/order/order-api/internal/config"
	"go-zero-shop/service/order/order-rpc/order"
	"go-zero-shop/service/product/product-rpc/product"
)

type ServiceContext struct {
	Config     config.Config
	OrderRpc   order.Order
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
