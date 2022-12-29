package main

import (
	"github.com/blackhorseya/ethscan/pkg/adapters"
	"go.uber.org/zap"
)

type grpc struct {
}

func NewGrpc(logger *zap.Logger) adapters.Grpc {
	return &grpc{}
}

func (g *grpc) InitService() error {
	// todo: 2022/12/30|sean|impl me
	return nil
}
