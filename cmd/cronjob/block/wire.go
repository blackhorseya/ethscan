//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/portto/internal/app/domain/biz"
	"github.com/blackhorseya/portto/internal/pkg/config"
	"github.com/blackhorseya/portto/internal/pkg/log"
	"github.com/blackhorseya/portto/internal/pkg/storage/mariadb"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infrastructure
	config.ProviderSet,
	log.ProviderSet,

	// storage
	mariadb.ProviderSet,

	// transports

	// implementation
	biz.ProviderSet,

	// main
	NewService,
	NewCronjobOptions,
	NewCronjob,
)

func CreateService(path string, initHeight uint64) (*Service, error) {
	panic(wire.Build(providerSet))
}
