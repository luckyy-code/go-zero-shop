package logic

import (
	"context"
	"go-zero-shop/service/order/model"
	"go-zero-shop/service/order/order-rpc/internal/svc"
	"go-zero-shop/service/order/order-rpc/order"
	"go-zero-shop/service/user/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *order.ListRequest) (*order.ListResponse, error) {
	// 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: in.Uid,
	})
	if err != nil {
		return nil, err
	}

	// 查询订单是否存在
	list, err := l.svcCtx.OrderModel.FindAllByUid(l.ctx, in.Uid)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "订单不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	orderList := make([]*order.DetailResponseOrder, 0)
	for _, item := range list {
		orderList = append(orderList, &order.DetailResponseOrder{
			Id:     item.Id,
			Uid:    item.Uid,
			Pid:    item.Pid,
			Amount: item.Amount,
			Status: item.Status,
		})
	}

	return &order.ListResponse{
		Data: orderList,
	}, nil
}
