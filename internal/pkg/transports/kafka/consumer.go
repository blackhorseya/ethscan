package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type ConsumerOptions struct {
	Connection map[string]interface{} `json:"connection" yaml:"connection"`
	Consumer   map[string]interface{} `json:"consumer" yaml:"consumer"`
}

func NewConsumerOptions(v *viper.Viper) (*ConsumerOptions, error) {
	opts := &ConsumerOptions{}

	err := v.UnmarshalKey("kafka", &opts)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal kafka key error")
	}

	return opts, nil
}

func NewConsumer(opts *ConsumerOptions) (*kafka.Consumer, error) {
	cfg := &kafka.ConfigMap{}
	for key, val := range opts.Connection {
		err := cfg.SetKey(key, val)
		if err != nil {
			return nil, err
		}
	}
	for key, val := range opts.Consumer {
		err := cfg.SetKey(key, val)
		if err != nil {
			return nil, err
		}
	}

	ret, err := kafka.NewConsumer(cfg)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

var ProviderConsumer = wire.NewSet(NewConsumerOptions, NewConsumer)
