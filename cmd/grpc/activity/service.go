package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/ethscan/pkg/adapters"
	"github.com/blackhorseya/ethscan/pkg/app"
	"github.com/blackhorseya/ethscan/pkg/grpcx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type service struct {
	logger     *zap.Logger
	grpcserver grpcx.Server
}

// NewService serve caller to create service instance
func NewService(logger *zap.Logger, gs grpcx.Server, grpc adapters.Grpc) (app.Service, error) {
	err := grpc.RegisterServer()
	if err != nil {
		return nil, errors.Wrap(err, "init service error")
	}

	svc := &service{
		logger:     logger.With(zap.String("type", "service")),
		grpcserver: gs,
	}

	return svc, nil
}

func (s *service) Start() error {
	if s.grpcserver != nil {
		err := s.grpcserver.Start()
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		s.logger.Info("receive a signal", zap.String("signal", sig.String()))

		if s.grpcserver != nil {
			err := s.grpcserver.Stop()
			if err != nil {
				s.logger.Warn("stop grpc server error", zap.Error(err))
			}
		}

		os.Exit(0)
	}

	return nil
}
