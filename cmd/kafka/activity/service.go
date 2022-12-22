package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/portto/pkg/app"
	"go.uber.org/zap"
)

type Service struct {
	logger *zap.Logger
}

func NewService(logger *zap.Logger) (app.Service, error) {
	svc := &Service{
		logger: logger.With(zap.String("type", "service")),
	}

	return svc, nil
}

func (s *Service) Start() error {
	// todo: 2022/12/23|sean|start the service

	return nil
}

func (s *Service) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		s.logger.Info("receive a signal", zap.String("signal", sig.String()))

		// todo: 2022/12/23|sean|stop the service

		os.Exit(0)
	}

	return nil
}
