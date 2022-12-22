//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/ethscan/internal/app/domain/block/biz"
	"github.com/blackhorseya/ethscan/internal/pkg/config"
	"github.com/blackhorseya/ethscan/internal/pkg/log"
	"github.com/blackhorseya/ethscan/internal/pkg/storage/mariadb"
	"github.com/blackhorseya/ethscan/internal/pkg/transports/httpx"
	"github.com/blackhorseya/ethscan/internal/pkg/transports/kafka"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infrastructure
	config.ProviderSet,
	log.ProviderSet,

	// storage
	mariadb.ProviderSet,

	// transports
	httpx.ProviderServerSet,
	kafka.ProviderProducer,

	// implementation
	biz.ProviderSet,

	// main
	NewService,
	NewRestful,
)

func CreateService(path string, id int64) (*Service, error) {
	panic(wire.Build(providerSet))
}
