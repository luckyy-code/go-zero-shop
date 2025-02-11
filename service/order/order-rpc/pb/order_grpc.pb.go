// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.1
// source: order.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Order_Create_FullMethodName       = "/pb.order/Create"
	Order_Update_FullMethodName       = "/pb.order/Update"
	Order_Remove_FullMethodName       = "/pb.order/Remove"
	Order_Detail_FullMethodName       = "/pb.order/Detail"
	Order_List_FullMethodName         = "/pb.order/List"
	Order_Paid_FullMethodName         = "/pb.order/Paid"
	Order_CreateRevert_FullMethodName = "/pb.order/CreateRevert"
)

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	Create(ctx context.Context, in *CreateRequestOrder, opts ...grpc.CallOption) (*CreateResponseOrder, error)
	Update(ctx context.Context, in *UpdateRequestOrder, opts ...grpc.CallOption) (*UpdateResponseOrder, error)
	Remove(ctx context.Context, in *RemoveRequestOrder, opts ...grpc.CallOption) (*RemoveResponseOrder, error)
	Detail(ctx context.Context, in *DetailRequestOrder, opts ...grpc.CallOption) (*DetailResponseOrder, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Paid(ctx context.Context, in *PaidRequest, opts ...grpc.CallOption) (*PaidResponse, error)
	CreateRevert(ctx context.Context, in *CreateRequestOrder, opts ...grpc.CallOption) (*CreateResponseOrder, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) Create(ctx context.Context, in *CreateRequestOrder, opts ...grpc.CallOption) (*CreateResponseOrder, error) {
	out := new(CreateResponseOrder)
	err := c.cc.Invoke(ctx, Order_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) Update(ctx context.Context, in *UpdateRequestOrder, opts ...grpc.CallOption) (*UpdateResponseOrder, error) {
	out := new(UpdateResponseOrder)
	err := c.cc.Invoke(ctx, Order_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) Remove(ctx context.Context, in *RemoveRequestOrder, opts ...grpc.CallOption) (*RemoveResponseOrder, error) {
	out := new(RemoveResponseOrder)
	err := c.cc.Invoke(ctx, Order_Remove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) Detail(ctx context.Context, in *DetailRequestOrder, opts ...grpc.CallOption) (*DetailResponseOrder, error) {
	out := new(DetailResponseOrder)
	err := c.cc.Invoke(ctx, Order_Detail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, Order_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) Paid(ctx context.Context, in *PaidRequest, opts ...grpc.CallOption) (*PaidResponse, error) {
	out := new(PaidResponse)
	err := c.cc.Invoke(ctx, Order_Paid_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) CreateRevert(ctx context.Context, in *CreateRequestOrder, opts ...grpc.CallOption) (*CreateResponseOrder, error) {
	out := new(CreateResponseOrder)
	err := c.cc.Invoke(ctx, Order_CreateRevert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	Create(context.Context, *CreateRequestOrder) (*CreateResponseOrder, error)
	Update(context.Context, *UpdateRequestOrder) (*UpdateResponseOrder, error)
	Remove(context.Context, *RemoveRequestOrder) (*RemoveResponseOrder, error)
	Detail(context.Context, *DetailRequestOrder) (*DetailResponseOrder, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Paid(context.Context, *PaidRequest) (*PaidResponse, error)
	CreateRevert(context.Context, *CreateRequestOrder) (*CreateResponseOrder, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) Create(context.Context, *CreateRequestOrder) (*CreateResponseOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderServer) Update(context.Context, *UpdateRequestOrder) (*UpdateResponseOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOrderServer) Remove(context.Context, *RemoveRequestOrder) (*RemoveResponseOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedOrderServer) Detail(context.Context, *DetailRequestOrder) (*DetailResponseOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Detail not implemented")
}
func (UnimplementedOrderServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOrderServer) Paid(context.Context, *PaidRequest) (*PaidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Paid not implemented")
}
func (UnimplementedOrderServer) CreateRevert(context.Context, *CreateRequestOrder) (*CreateResponseOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRevert not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequestOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Create(ctx, req.(*CreateRequestOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequestOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Update(ctx, req.(*UpdateRequestOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRequestOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_Remove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Remove(ctx, req.(*RemoveRequestOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_Detail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailRequestOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Detail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_Detail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Detail(ctx, req.(*DetailRequestOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_Paid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PaidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Paid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_Paid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Paid(ctx, req.(*PaidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_CreateRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequestOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).CreateRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Order_CreateRevert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).CreateRevert(ctx, req.(*CreateRequestOrder))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Order_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Order_Update_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Order_Remove_Handler,
		},
		{
			MethodName: "Detail",
			Handler:    _Order_Detail_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Order_List_Handler,
		},
		{
			MethodName: "Paid",
			Handler:    _Order_Paid_Handler,
		},
		{
			MethodName: "CreateRevert",
			Handler:    _Order_CreateRevert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
