package logic

import (
	"context"
	"go-zero-shop/service/product/model"
	"go-zero-shop/service/product/product-rpc/product"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-shop/service/product/product-rpc/internal/svc"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *product.UpdateRequest) (*product.UpdateResponse, error) {
	// 查询产品是否存在
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	if in.Name != "" {
		res.Name = in.Name
	}
	if in.Desc != "" {
		res.Desc = in.Desc
	}
	if in.Stock != 0 {
		res.Stock = in.Stock
	}
	if in.Amount != 0 {
		res.Amount = in.Amount
	}
	if in.Status != 0 {
		res.Status = in.Status
	}

	err = l.svcCtx.ProductModel.Update(l.ctx, res)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.UpdateResponse{}, nil
}
