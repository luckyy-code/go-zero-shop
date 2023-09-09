package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-shop/service/order/model"
	"go-zero-shop/service/order/order-rpc/internal/config"
	"go-zero-shop/service/product/product-rpc/product"
	"go-zero-shop/service/user/user-rpc/user"
)

type ServiceContext struct {
	Config     config.Config
	OrderModel model.OrderModel
	UserRpc    user.User
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
