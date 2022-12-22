package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/blackhorseya/ethscan/pkg/adapters"
	"github.com/blackhorseya/ethscan/pkg/app"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type service struct {
	logger  *zap.Logger
	cronjob adapters.Cronjob
}

// NewService serve caller to create service instance
func NewService(logger *zap.Logger, cronjob adapters.Cronjob) (app.Service, error) {
	svc := &service{
		logger:  logger.With(zap.String("type", "service")),
		cronjob: cronjob,
	}

	return svc, nil
}

func (s *service) Start() error {
	if s.cronjob != nil {
		err := s.cronjob.Start()
		if err != nil {
			return errors.Wrap(err, "cronjob server start error")
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

		if s.cronjob != nil {
			err := s.cronjob.Stop()
			if err != nil {
				s.logger.Warn("stop cronjob server error", zap.Error(err))
			}
		}

		os.Exit(0)
	}

	return nil
}
