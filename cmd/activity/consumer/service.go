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

type Service struct {
	logger *zap.Logger
	kafka  adapters.Consumer
}

func NewService(logger *zap.Logger, kafka adapters.Consumer) (app.Service, error) {
	svc := &Service{
		logger: logger.With(zap.String("type", "service")),
		kafka:  kafka,
	}

	return svc, nil
}

func (s *Service) Start() error {
	if s.kafka != nil {
		err := s.kafka.Start()
		if err != nil {
			return errors.Wrap(err, "kafka start error")
		}
	}

	return nil
}

func (s *Service) AwaitSignal() error {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	if sig := <-c; true {
		s.logger.Info("receive a signal", zap.String("signal", sig.String()))

		if s.kafka != nil {
			err := s.kafka.Stop()
			if err != nil {
				s.logger.Warn("stop kafka error", zap.Error(err))
			}
		}

		os.Exit(0)
	}

	return nil
}
