package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type ProducerOptions struct {
	Connection map[string]interface{} `json:"connection" yaml:"connection"`
}

func NewProducerOptions(v *viper.Viper) (*ProducerOptions, error) {
	opts := &ProducerOptions{}

	err := v.UnmarshalKey("kafka", &opts)
	if err != nil {
		return nil, errors.Wrap(err, "unmarshal kafka key error")
	}

	return opts, nil
}

func NewProducer(opts *ProducerOptions) (*kafka.Producer, error) {
	cfg := &kafka.ConfigMap{}
	for key, val := range opts.Connection {
		err := cfg.SetKey(key, val)
		if err != nil {
			return nil, err
		}
	}

	ret, err := kafka.NewProducer(cfg)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

var ProviderProducer = wire.NewSet(NewProducerOptions, NewProducer)
