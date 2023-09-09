// Code generated by goctl. DO NOT EDIT.
// Source: order.proto

package server

import (
	"context"

	"go-zero-shop/service/order/order-rpc/internal/logic"
	"go-zero-shop/service/order/order-rpc/internal/svc"
	"go-zero-shop/service/order/order-rpc/pb"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) Create(ctx context.Context, in *pb.CreateRequestOrder) (*pb.CreateResponseOrder, error) {
	l := logic.NewCreateLogic(ctx, s.svcCtx)
	return l.Create(in)
}

func (s *OrderServer) Update(ctx context.Context, in *pb.UpdateRequestOrder) (*pb.UpdateResponseOrder, error) {
	l := logic.NewUpdateLogic(ctx, s.svcCtx)
	return l.Update(in)
}

func (s *OrderServer) Remove(ctx context.Context, in *pb.RemoveRequestOrder) (*pb.RemoveResponseOrder, error) {
	l := logic.NewRemoveLogic(ctx, s.svcCtx)
	return l.Remove(in)
}

func (s *OrderServer) Detail(ctx context.Context, in *pb.DetailRequestOrder) (*pb.DetailResponseOrder, error) {
	l := logic.NewDetailLogic(ctx, s.svcCtx)
	return l.Detail(in)
}

func (s *OrderServer) List(ctx context.Context, in *pb.ListRequest) (*pb.ListResponse, error) {
	l := logic.NewListLogic(ctx, s.svcCtx)
	return l.List(in)
}

func (s *OrderServer) Paid(ctx context.Context, in *pb.PaidRequest) (*pb.PaidResponse, error) {
	l := logic.NewPaidLogic(ctx, s.svcCtx)
	return l.Paid(in)
}

func (s *OrderServer) CreateRevert(ctx context.Context, in *pb.CreateRequestOrder) (*pb.CreateResponseOrder, error) {
	l := logic.NewCreateRevertLogic(ctx, s.svcCtx)
	return l.CreateRevert(in)
}
