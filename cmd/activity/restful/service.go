package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/ethscan/pkg/adapters"
	"github.com/blackhorseya/ethscan/pkg/app"
	"github.com/blackhorseya/ethscan/pkg/httpx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type service struct {
	logger     *zap.Logger
	httpserver httpx.Server
}

// NewService serve caller to create service instance
func NewService(logger *zap.Logger, hs httpx.Server, restful adapters.Restful) (app.Service, error) {
	err := restful.InitRouting()
	if err != nil {
		return nil, errors.Wrap(err, "init routing error")
	}

	svc := &service{
		logger:     logger.With(zap.String("type", "service")),
		httpserver: hs,
	}

	return svc, nil
}

func (s *service) Start() error {
	if s.httpserver != nil {
		err := s.httpserver.Start()
		if err != nil {
			return errors.Wrap(err, "http server start error")
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

		if s.httpserver != nil {
			err := s.httpserver.Stop()
			if err != nil {
				s.logger.Warn("stop http server error", zap.Error(err))
			}
		}

		os.Exit(0)
	}

	return nil
}
