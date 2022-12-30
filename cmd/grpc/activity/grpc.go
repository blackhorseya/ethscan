package main

import (
	"context"

	"github.com/blackhorseya/ethscan/pkg/adapters"
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	am "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/model"
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

func (i *impl) ListTxnsByBlockHash(ctx context.Context, req *am.ListTxnsByBlockHashRequest) (*am.ListTxnsByBlockHashResponse, error) {
	// todo: 2022/12/30|sean|impl me
	panic("implement me")
}
