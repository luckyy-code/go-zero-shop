package logic

import (
	"context"
	"encoding/json"
	"go-zero-shop/service/user/user-api/internal/svc"
	"go-zero-shop/service/user/user-api/internal/types"
	"go-zero-shop/service/user/user-rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	uid, _ := l.ctx.Value("uid").(json.Number).Int64()
	res, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}

	return &types.UserInfoResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
