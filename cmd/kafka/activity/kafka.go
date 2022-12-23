package main

import (
	"github.com/blackhorseya/ethscan/pkg/adapters"
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

type impl struct {
	logger   *zap.Logger
	biz      ab.IBiz
	consumer *kafka.Consumer
}

func NewKafka(logger *zap.Logger, consumer *kafka.Consumer, biz ab.IBiz) adapters.Kafka {
	return &impl{
		logger:   logger,
		biz:      biz,
		consumer: consumer,
	}
}

func (i *impl) Start() error {
	// todo: 2022/12/23|sean|impl me
	panic("implement me")
}

func (i *impl) Stop() error {
	// todo: 2022/12/23|sean|impl me
	panic("implement me")
}
