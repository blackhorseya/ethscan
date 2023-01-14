//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/ethscan/internal/adapter/consumer"
	"github.com/blackhorseya/ethscan/internal/entity/domain/activity/biz"
	"github.com/blackhorseya/ethscan/internal/pkg/config"
	"github.com/blackhorseya/ethscan/internal/pkg/log"
	"github.com/blackhorseya/ethscan/internal/pkg/storage/mariadb"
	"github.com/blackhorseya/ethscan/internal/pkg/transports/kafka"
	"github.com/blackhorseya/ethscan/pkg/app"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infrastructure
	config.ProviderSet,
	log.ProviderSet,

	// storage
	mariadb.ProviderSet,

	// transports
	kafka.ProviderConsumer,

	// implementation
	consumer.ActivitySet,
	biz.ActivitySet,

	// main
	NewService,
)

func CreateService(path string, id int64) (app.Service, error) {
	panic(wire.Build(providerSet))
}
