//go:generate wire
//go:build wireinject

package main

import (
	"github.com/blackhorseya/portto/internal/pkg/config"
	"github.com/blackhorseya/portto/internal/pkg/log"
	"github.com/blackhorseya/portto/pkg/app"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	// infrastructure
	config.ProviderSet,
	log.ProviderSet,

	// storage

	// transports

	// implementation

	// main
	NewService,
)

func CreateService(path string, id int64) (app.Service, error) {
	panic(wire.Build(providerSet))
}
