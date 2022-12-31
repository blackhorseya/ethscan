package main

import (
	"encoding/json"
	"time"

	"github.com/blackhorseya/ethscan/pkg/adapters"
	"github.com/blackhorseya/ethscan/pkg/contextx"
	ab "github.com/blackhorseya/ethscan/pkg/entity/domain/activity/biz"
	bm "github.com/blackhorseya/ethscan/pkg/entity/domain/block/model"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

type impl struct {
	logger   *zap.Logger
	biz      ab.IBiz
	consumer *kafka.Consumer

	done chan struct{}
}

func NewKafka(logger *zap.Logger, consumer *kafka.Consumer, biz ab.IBiz) adapters.Kafka {
	return &impl{
		logger:   logger,
		biz:      biz,
		consumer: consumer,
		done:     make(chan struct{}),
	}
}

func (i *impl) Start() error {
	i.logger.Info("starting kafka consumer...")

	go i.subscribe()

	i.logger.Info("started kafka consumer")

	return nil
}

func (i *impl) Stop() error {
	i.logger.Info("stopping kafka consumer...")

	i.done <- struct{}{}
	_ = i.consumer.Close()

	i.logger.Info("stop kafka consumer...")

	return nil
}

func (i *impl) subscribe() {
	topics := []string{"new_block"}
	_ = i.consumer.SubscribeTopics(topics, nil)

	for {
		select {
		case <-i.done:
			return
		default:
			message, err := i.consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				continue
			}

			topic := *(message.TopicPartition.Topic)
			if topic == "new_block" {
				var newBlock *bm.BlockRecord
				err = json.Unmarshal(message.Value, &newBlock)
				if err != nil {
					i.logger.Error("cannot parse message value to block record", zap.Error(err), zap.String("value", string(message.Value)))
					continue
				}

				i.logger.Info("received new block", zap.Uint64("height", newBlock.Height), zap.String("hash", newBlock.Hash))

				ctx := contextx.BackgroundWithLogger(i.logger)

				_, err = i.biz.HandleNewBlock(ctx, newBlock)
				if err != nil {
					i.logger.Error("handle new block error", zap.Error(err), zap.Any("new_block", newBlock))
					continue
				}
			}
		}
	}
}
