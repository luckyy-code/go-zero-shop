// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"go-zero-shop/service/user/user-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginRequest     = pb.LoginRequest
	LoginResponse    = pb.LoginResponse
	RegisterRequest  = pb.RegisterRequest
	RegisterResponse = pb.RegisterResponse
	UserInfoRequest  = pb.UserInfoRequest
	UserInfoResponse = pb.UserInfoResponse

	User interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}
