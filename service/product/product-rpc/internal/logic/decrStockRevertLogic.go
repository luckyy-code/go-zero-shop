package logic

import (
	"context"
	"database/sql"
	"go-zero-shop/service/product/product-rpc/internal/svc"
	"go-zero-shop/service/product/product-rpc/product"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
)

type DecrStockRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockRevertLogic {
	return &DecrStockRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DecrStockRevertLogic) DecrStockRevert(in *product.DecrStockRequest) (*product.DecrStockResponse, error) {
	// 获取 RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	// 获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}
	// 开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		// 更新产品库存
		_, err := l.svcCtx.ProductModel.TxAdjustStock(l.ctx, tx, in.Id, 1)
		return err
	})

	if err != nil {
		return nil, err
	}

	return &product.DecrStockResponse{}, nil
}
