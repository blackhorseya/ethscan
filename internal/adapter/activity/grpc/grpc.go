package grpc

import (
	"context"

	"github.com/blackhorseya/ethscan/internal/pkg/errorx"
	"github.com/blackhorseya/ethscan/pkg/adapters"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
	"github.com/google/wire"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type impl struct {
	gs  *grpc.Server
	biz ab.IBiz
}

func NewGrpc(logger *zap.Logger, gs *grpc.Server, biz ab.IBiz) adapters.Grpc {
	return &impl{
		gs:  gs,
		biz: biz,
	}
}

func (i *impl) RegisterServer() error {
	am.RegisterServiceServer(i.gs, i)

	return nil
}

func (i *impl) ListTxnsByBlockHash(c context.Context, req *am.ListTxnsByBlockHashRequest) (*am.ListTxnsByBlockHashResponse, error) {
	ctx, ok := c.Value(contextx.KeyCtx).(contextx.Contextx)
	if !ok {
		return nil, errorx.ErrContextx
	}

	condition := ab.ListTxnsCondition{
		BlockHash: req.Hash,
	}
	ret, err := i.biz.ListTxns(ctx, condition)
	if err != nil {
		ctx.Error(errorx.ErrGetTx.LogMessage, zap.Error(err), zap.Any("condition", condition))
		return nil, err
	}

	return &am.ListTxnsByBlockHashResponse{
		Transactions: ret,
	}, nil
}

var ActivitySet = wire.NewSet(NewGrpc)
