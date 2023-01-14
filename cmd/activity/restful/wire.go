//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/ethscan/internal/adapter/activity/restful"
	"github.com/blackhorseya/ethscan/internal/entity/domain/activity/biz"
	"github.com/blackhorseya/ethscan/internal/pkg/config"
	"github.com/blackhorseya/ethscan/internal/pkg/log"
	"github.com/blackhorseya/ethscan/internal/pkg/storage/mariadb"
	"github.com/blackhorseya/ethscan/internal/pkg/transports/httpx"
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
	httpx.ProviderServerSet,

	// implementation
	restful.ActivitySet,
	biz.ActivitySet,

	// main
	NewService,
)

func CreateService(path string, id int64) (app.Service, error) {
	panic(wire.Build(providerSet))
}
