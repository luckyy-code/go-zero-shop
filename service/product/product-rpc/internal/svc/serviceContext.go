package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-shop/service/product/model"
	"go-zero-shop/service/product/product-rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn, c.CacheRedis),
	}
}
